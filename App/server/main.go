package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const port = ":9092"

func main() {
	spaHandler := SpaHandler("../assets", "index.html")
	http.Handle("/", spaHandler)
	log.Printf("Server started on http://localhost%s ...", port)
	http.ListenAndServe(port, nil)
}

type spaHandler struct {
	publicDir string // The directory from which to serve
	indexFile string // The fallback/default file to serve
}

func (h *spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p := filepath.Join(h.publicDir, filepath.Clean(r.URL.Path))

	if info, err := os.Stat(p); err != nil {
		http.ServeFile(w, r, filepath.Join(h.publicDir, h.indexFile))
		return
	} else if info.IsDir() {
		http.ServeFile(w, r, filepath.Join(h.publicDir, h.indexFile))
		return
	}

	http.ServeFile(w, r, p)
}

func SpaHandler(publicDir string, indexFile string) http.Handler {
	return &spaHandler{publicDir, indexFile}
}
