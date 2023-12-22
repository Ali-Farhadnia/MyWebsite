package main

import (
	"crypto/tls"
	"net/http"

	"golang.org/x/crypto/acme/autocert"
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

	// Create an autocert manager to handle TLS certificates
	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("alifarhadnia.com"), // Replace with your domain
		Cache:      autocert.DirCache("/var/www/certs"),        // Directory outside the Go project for storing certificates
	}

	// Configure HTTPS server
	server := &http.Server{
		Addr:    ":443", // Default HTTPS port
		Handler: nil,    // Handler is set to nil for automatic configuration with autocert
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
	}

	// Start the HTTPS server
	println("Server running on port 443 (HTTPS)")
	err := server.ListenAndServeTLS("", "") // Empty strings because TLS config is managed by autocert
	if err != nil {
		panic(err)
	}
}
