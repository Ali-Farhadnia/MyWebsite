package main

import (
	"net/http"
)

func main() {
	// Specify the directory containing your static files
	fs := http.FileServer(http.Dir("./static"))

	// Serve index.html at the root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})

	// Serve other static files using the file server
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Start the HTTP server to work with Cloudflare Flexible SSL
	println("Server running on port 80 (HTTP)")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}
