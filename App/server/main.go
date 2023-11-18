package main

import (
	"log"
	"net/http"
)

const port = ":9092"

func main() {
	//http.Handle("/App/", http.StripPrefix("/App/", http.FileServer(http.Dir("assets/"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../assets/index.html")
	})

	http.HandleFunc("/profile.png", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../assets/profile.png")
	})

	// Servir les fichiers statiques (CSS, JavaScript, WebAssembly)
	http.HandleFunc("/index.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../assets/index.css")
	})

	http.HandleFunc("/script.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../assets/script.js")
	})

	http.HandleFunc("/main.wasm", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../assets/main.wasm")
	})

	http.HandleFunc("/wasm_exec.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../assets/wasm_exec.js")
	})
	log.Printf("Server started on http://localhost%s ...", port)
	http.ListenAndServe(port, nil)
}
