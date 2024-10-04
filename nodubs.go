package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/caddyserver/certmagic"
)

func main() {
	// Set up certmagic to use a filesystem cache
	certmagic.Default.Storage = &certmagic.FileStorage{Path: "./certs"}

	// Enable on-demand TLS
	certmagic.Default.OnDemand = new(certmagic.OnDemandConfig)

	// Use Let's Encrypt's staging environment for testing
	certmagic.DefaultACME.CA = certmagic.LetsEncryptStagingCA

	// Ze Handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			fmt.Fprintf(w, "Hello, World")
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Start HTTPS server with on-demand TLS
	err := certmagic.HTTPS([]string{}, nil) // Empty slice means no predefined domains
	if err != nil {
		log.Fatal(err)
	}
}
