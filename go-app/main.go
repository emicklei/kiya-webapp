package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kramphub/kiya/backend"
)

var isDebug = false
var prof *backend.Profile

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/fetch", handleFetch)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	prof = &backend.Profile{
		ProjectID: os.Getenv("PROJECT_ID"),
		Location:  os.Getenv("LOCATION"),
		Keyring:   os.Getenv("KEY_RING"),
		CryptoKey: os.Getenv("CRYPTO_KEY"),
		Bucket:    os.Getenv("BUCKET"),
	}
	log.Printf("%#v\n", prof)

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("%v", err)
	}
}
