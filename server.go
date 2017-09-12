package main

import (
	"fmt"
	"net/http"
	"runtime"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world, I'm running on %s with an %s CPU ", runtime.GOOS, runtime.GOARCH)
} 

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Listening on localhost:8080...")
	http.ListenAndServe(":8080", nil)
}