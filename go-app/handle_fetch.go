package main

import (
	"io"
	"net/http"
)

func handleFetch(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	isDebug = r.Form.Get("debug") != ""

	name := r.Form.Get("name")

	back, err := newKMSBackend()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}

	data, err := back.Get(r.Context(), prof, name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}

	w.Header().Set("content-type", "text/plain")
	io.WriteString(w, string(data))
}
