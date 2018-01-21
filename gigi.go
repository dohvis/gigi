package main
import (
	"fmt"
	// "log"
	"net/http"
	"os/exec"
)

func handler(response http.ResponseWriter, request *http.Request) {
	shellScript := "./pull-and-build.sh"
	cmd := exec.Command("bash", shellScript)
	cmd.Run()
	fmt.Fprintf(response, "Done")
}

func main() {
	port := 8001
	fmt.Printf("Running on localhost:%d..\n", port)
	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
