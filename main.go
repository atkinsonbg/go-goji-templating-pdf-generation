package main

import (
	"github.com/atkinsonbg/go-goji-templating-pdf-generation/handlers"
	"goji.io"
	"goji.io/pat"
	"log"
	"net/http"
)

func main() {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/health"), handlers.HealthHandler)
	mux.HandleFunc(pat.Get("/template"), handlers.TemplatingHandler)
	log.Fatal(http.ListenAndServe(":8000", mux))
}
