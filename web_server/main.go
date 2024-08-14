package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/home.html")
}

func main() {
	fmt.Println("web server in go lang!")
	// Serve files from the "static" directory
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	// Associate the handler function with a route
	http.HandleFunc("/hello", helloHandler)

	http.HandleFunc("/home", serveHome)

	// Start the web server on port 8080
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
