package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/emicklei/renderbee"
	"github.com/kramphub/kiya/backend"
)

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

	selection := []backend.Key{}
	query := r.Form.Get("q")
	if query != "" {

		back, err := newKMSBackend()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, err.Error())
			return
		}

		keys, err := back.List(r.Context(), prof)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, err.Error())
			return
		}

		for _, each := range keys {
			if strings.Contains(each.Name, query) {
				selection = append(selection, each)
			}
		}
		sort.Slice(selection, func(i, j int) bool {
			return selection[i].Name < selection[j].Name
		})
	}

	w.Header().Set("content-type", "text/html")
	canvas := renderbee.NewHtmlCanvas(w)
	page := renderbee.NewFragmentMap(PageLayout_Template)
	page.Put("TableOfKeys", TableOfKeys{Keys: selection})
	canvas.Render(page)
}
