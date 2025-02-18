package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Serve static files (HTML, JS, CSS)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	port := 8080
	fmt.Printf("Server started on http://localhost:%d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

