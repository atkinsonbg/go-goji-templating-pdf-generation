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
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
