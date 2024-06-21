package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
)

func hostHandler(w http.ResponseWriter, r *http.Request) {
	name, err := os.Hostname()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "<h1>Hostname:  %s!</h1><br>", name)
	fmt.Fprintf(w, "<h2>Environment Variables:  %s!</h2><br>", name)
	fmt.Fprintf(w, "<ul>")


	for _, evar := range os.Environ() {
		fmt.Fprintf(w, "<li>%s</li><br>", evar)
	}
	fmt.Fprintf(w, "</ul><br>")

	fmt.Fprintf(w, "<h3>Go Version:  %s!</h3><br>", runtime.Version())
	fmt.Fprintf(w, "<h4>Go OS/Arch:  %s!</h4><br>", runtime.GOOS + "/" + runtime.GOARCH)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello, World!</h1><br>")
	fmt.Fprintf(w, "<a href='/host'>Host Info</a><br>")
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/host/", hostHandler)
	http.ListenAndServe(":8080", nil)
}
