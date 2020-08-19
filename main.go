package main

import (
	"log"
	"net/http"

	"github.com/atkinsonbg/go-goji-templating-pdf-generation/handlers"
	"goji.io"
	"goji.io/pat"
)

func main() {
	log.SetFlags(log.Llongfile)

	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/health"), handlers.HealthHandler)
	mux.HandleFunc(pat.Post("/template"), handlers.TemplatingHandler)
	log.Fatal(http.ListenAndServe(":8000", mux))
}
