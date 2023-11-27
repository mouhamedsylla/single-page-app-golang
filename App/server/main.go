package main

import (
	"log"
	"net/http"
	"os"
	"io/fs"
)

const port = ":9092"

func main() {
	var myFS fs.FS = os.DirFS("../assets")
	http.FileServer(http.FS(myFS))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../assets/index.html")
	})
	log.Printf("Server started on http://localhost%s ...", port)
	http.ListenAndServe(port, nil)
}
