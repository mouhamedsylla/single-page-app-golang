package main

import (
	"net/http"
	"log"
)

const port = ":9092"

func main() {
	http.Handle("/", http.FileServer(http.Dir("../assets")))
	log.Printf("Server started on http://localhost%s ...", port)
	http.ListenAndServe(port, nil)
}