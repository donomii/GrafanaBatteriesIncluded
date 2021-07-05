package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/donomii/goof"
)

func metrics(w http.ResponseWriter, req *http.Request) {
	out := `# HELP process Process details.
# TYPE process gauge
`
	output := goof.QC([]string{"ps", "auxc"})
	lines := strings.Split(output, "\n")
	lines = goof.ListGrepInv("USER", lines)
	for _, line := range lines {
		line = strings.ReplaceAll(line, "\t", " ")
		line = strings.ReplaceAll(line, "   ", " ")
		line = strings.ReplaceAll(line, "  ", " ")
		line = strings.ReplaceAll(line, "  ", " ")
		line = strings.ReplaceAll(line, "  ", " ")
		line = strings.ReplaceAll(line, "  ", " ")
		fields := strings.SplitN(line, " ", 11)
		if len(fields) > 10 {
			out = out + fmt.Sprintf("process{name=\"%v\",stat=\"cpu\"} %v\n", fields[10], fields[2])
			out = out + fmt.Sprintf("process{name=\"%v\",stat=\"memory\"} %v\n", fields[10], fields[3])
		}
	}
	log.Println(out)
	fmt.Fprintf(w, out)
}
func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/metrics", metrics)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}
