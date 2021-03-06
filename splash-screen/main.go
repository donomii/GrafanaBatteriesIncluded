package main

//go:generate go mod init github.com/donomii/splash-screen
//go:generate go mod tidy

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"time"

	"github.com/donomii/goof"
	"github.com/mattn/go-shellwords"

	//"time"

	"github.com/go-gl/mathgl/mgl32"

	"github.com/go-gl/gl/v3.2-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"

	_ "embed"
)

//go:embed logo.png
var logo_bytes []byte

// Arrange that main.main runs on main thread.
func init() {
	runtime.LockOSThread()
	debug.SetGCPercent(-1)
}

type State struct {
	prop           int32
	Program        uint32
	Vao            uint32
	Vbo            uint32
	Texture        uint32
	TextureUniform int32
	VertAttrib     uint32
	Angle          float64
	PreviousTime   float64
	ModelUniform   int32
	TexCoordAttrib uint32
}

var winWidth = 180
var winHeight = 180
var lasttime float64

var rotX, roty float64
var launchShellList arrayFlags
var launchList arrayFlags
var runningProcs []context.CancelFunc

func drainChannel(ch chan []byte) {
	for {
		<-ch
		time.Sleep(10 * time.Millisecond)
	}
}
func main() {
	flag.Var(&launchShellList, "launchShell", "Run shell command at start.  May be repeated to launch multiple commands.")
	flag.Var(&launchList, "launch", "Command line to start an app.  May be repeated to launch multiple apps.")
	flag.Parse()
	for _, commandStr := range launchShellList {
		ctx, cancel := context.WithCancel(context.Background())
		command := exec.CommandContext(ctx, "/bin/sh", "-c", commandStr)
		_, out, err := goof.WrapCmd(command, 3)
		runningProcs = append(runningProcs, cancel)
		go drainChannel(out)
		go drainChannel(err)
	}
	currentDir, _ := os.Getwd()
	for _, commandStr := range launchList {
		log.Printf("Launching %v", commandStr)
		os.Chdir(currentDir)
		args, _ := shellwords.Parse(commandStr)
		directory := filepath.Dir(args[0])
		os.Chdir(directory)
		exe := "./" + filepath.Base(args[0])
		X, _ := os.Getwd()
		log.Printf("In dir %v", X)
		log.Printf("Exe %v", exe)
		ctx, cancel := context.WithCancel(context.Background())
		command := exec.CommandContext(ctx, exe, args[1:]...)
		_, out, err := goof.WrapCmd(command, 3)
		runningProcs = append(runningProcs, cancel)
		go drainChannel(out)
		go drainChannel(err)
	}
	os.Chdir(currentDir)
	log.Println("Starting windowing system")
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	win, err := glfw.CreateWindow(winWidth, winHeight, "Grafana", nil, nil)
	if err != nil {
		panic(err)
	}
	go func() {
		time.Sleep(5 * time.Second)
		win.Iconify()
	}()

	win.MakeContextCurrent()

	win.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		log.Printf("Got key %c,%v,%v,%v", key, key, mods, action)
		handleKey(w, key, scancode, action, mods)
	})

	win.SetMouseButtonCallback(handleMouseButton)

	win.SetCursorPosCallback(handleMouseMove)

	if err := gl.Init(); err != nil {
		panic(err)
	}

	state := &State{
		prop: 1,
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)

	// Configure the vertex and fragment shaders
	state.Program, err = newProgram(vertexShader, fragmentShader)
	if err != nil {
		panic(err)
	}

	//Activate the program we just created.  This means we will use the render and fragment shaders we compiled above
	gl.UseProgram(state.Program)

	//Set a default projection matrix
	projection := mgl32.Perspective(mgl32.DegToRad(45.0), float32(winWidth)/float32(winHeight), 0.1, 10.0)
	projectionUniform := gl.GetUniformLocation(state.Program, gl.Str("projection\x00"))
	gl.UniformMatrix4fv(projectionUniform, 1, false, &projection[0])

	//Setup the camera
	camera := mgl32.LookAtV(mgl32.Vec3{3, 3, 3}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
	cameraUniform := gl.GetUniformLocation(state.Program, gl.Str("camera\x00"))
	gl.UniformMatrix4fv(cameraUniform, 1, false, &camera[0])

	//Setup the cube
	model := mgl32.Ident4()
	state.ModelUniform = gl.GetUniformLocation(state.Program, gl.Str("model\x00"))
	gl.UniformMatrix4fv(state.ModelUniform, 1, false, &model[0])

	//Find the location of the texture, so we can upload a picture to it
	state.TextureUniform = gl.GetUniformLocation(state.Program, gl.Str("tex\x00"))
	gl.Uniform1i(state.TextureUniform, 0)

	//This is the variable in the fragment shader that will hold the colour for each pixel
	gl.BindFragDataLocation(state.Program, 0, gl.Str("outputColor\x00"))

	// Load the texture
	state.Texture, err = newTexture(logo_bytes)
	if err != nil {
		log.Fatalln(err)
	}

	// Configure the vertex data
	gl.GenVertexArrays(1, &state.Vao)
	gl.BindVertexArray(state.Vao)

	gl.GenBuffers(1, &state.Vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, state.Vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(cubeVertices)*4, gl.Ptr(cubeVertices), gl.STATIC_DRAW)
	checkGlError()
	state.VertAttrib = uint32(gl.GetAttribLocation(state.Program, gl.Str("vert\x00")))
	gl.EnableVertexAttribArray(state.VertAttrib)
	gl.VertexAttribPointer(state.VertAttrib, 3, gl.FLOAT, false, 5*4, gl.PtrOffset(0))
	checkGlError()
	state.TexCoordAttrib = uint32(gl.GetAttribLocation(state.Program, gl.Str("vertTexCoord\x00")))
	gl.EnableVertexAttribArray(state.TexCoordAttrib)
	gl.VertexAttribPointer(state.TexCoordAttrib, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))
	checkGlError()

	// Configure global settings
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)
	gl.UseProgram(state.Program)
	gl.ClearColor(1.0, 1.0, 1.0, 0.0)

	//Activate the cube data, which will be drawn
	gl.BindVertexArray(state.Vao)

	//Choose the texture we just created and uploaded
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, state.Texture)

	for !win.ShouldClose() {

		gfxMain(win, state)
		glfw.PollEvents()
		time.Sleep(10 * time.Millisecond)
	}
	shutdown()

}

func shutdown() {
	for x, cancelF := range runningProcs {
		log.Printf("Stopping sub-process %v", x)
		cancelF()
	}
}

func gfxMain(win *glfw.Window, state *State) {
	//fmt.Println("Draw")
	//width, height := win.GetSize()
	//gl.Viewport(0, 0, int32(width-1), int32(height-1))

	// Render

	// Update

	now := glfw.GetTime()
	elapsed := now - state.PreviousTime

	if elapsed > 0.050 && 1 != win.GetAttrib(glfw.Iconified) {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		//fmt.Printf("elapsed: %v\n", elapsed)
		state.PreviousTime = now
		angle := state.Angle
		angle += elapsed
		state.Angle = angle

		model := mgl32.HomogRotate3D(float32(angle+rotX), mgl32.Vec3{0, 1, 0})

		// Render

		gl.UniformMatrix4fv(state.ModelUniform, 1, false, &model[0])

		gl.DrawArrays(gl.TRIANGLES, 0, 6*2*3)

		win.SwapBuffers()

	}
	time.Sleep(10 * time.Millisecond)
}

func checkGlError() {

	err := gl.GetError()
	if err > 0 {
		errStr := fmt.Sprintf("GLerror: %v\n", err)
		fmt.Printf(errStr)
		panic(errStr)
	}

}
