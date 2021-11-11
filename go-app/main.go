package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/emicklei/renderbee"
	"github.com/kramphub/kiya/backend"
)

func main() {
	http.HandleFunc("/", handleIndex)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("%v", err)
	}
}

var isDebug = false

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	isDebug = r.Form.Get("debug") != ""

	if isDebug {
		for _, k := range os.Environ() {
			fmt.Fprintf(w, "%s:%s\n", k, os.Getenv(k))
		}
		for k, v := range r.Header {
			fmt.Fprintf(w, "%s:%v\n", k, v)
		}
	}

	query := r.Form.Get("q")

	back, err := newKMSBackend()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}
	prof := &backend.Profile{
		ProjectID: os.Getenv("PROJECT_ID"),
		Location:  os.Getenv("LOCATION"),
		Keyring:   os.Getenv("KEY_RING"),
		CryptoKey: os.Getenv("CRYPTO_KEY"),
		Bucket:    os.Getenv("BUCKET"),
	}
	keys, err := back.List(r.Context(), prof)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}

	w.Header().Set("content-type", "text/html")
	selection := []backend.Key{}
	for _, each := range keys {
		if strings.Contains(each.Name, query) {
			selection = append(selection, each)
		}
	}
	sort.Slice(selection, func(i, j int) bool {
		return selection[i].Name < selection[j].Name
	})

	canvas := renderbee.NewHtmlCanvas(w)
	page := renderbee.NewFragmentMap(PageLayout_Template)
	page.Put("TableOfKeys", TableOfKeys{Keys: selection})
	canvas.Render(page)
}
