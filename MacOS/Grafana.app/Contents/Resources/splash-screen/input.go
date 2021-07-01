package main

import (
	"log"

	"github.com/go-gl/glfw/v3.2/glfw"
)

var oldXpos float64

func handleMouseMove(w *glfw.Window, xpos float64, ypos float64) {
	//log.Printf("Mouse moved: %v,%v", xpos, ypos)
	diff := xpos - oldXpos
	rotX = rotX + diff/180.0
	oldXpos = xpos
}
func handleMouseButton(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
	log.Printf("Got mouse button %v,%v,%v", button, mod, action)
	//handleKey(w, key, scancode, action, mods)
}

func handleKey(win *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {

	if mods == 2 && action == 1 && key != 341 {
		mask := ^byte(64 + 128)
		log.Printf("key mask: %#b", mask)
		val := byte(key)
		log.Printf("key val: %#b", val)
		b := mask & val
		log.Printf("key byte: %#b", b)

	}

	if action == 0 && mods == 0 {
		log.Printf("Acting on  key %c,%v", key, key)
		switch key {

		case 256:
			log.Println("Quitting")
			win.SetShouldClose(true)

		}
	}

}
