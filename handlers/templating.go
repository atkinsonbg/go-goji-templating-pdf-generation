package handlers

import (
	"log"
	"net/http"

	"github.com/atkinsonbg/go-goji-templating-pdf-generation/fileio"
)

// TemplatingHandler is responsible for taking input from the HTTP call, executing the template, and returning a PDF
func TemplatingHandler(w http.ResponseWriter, r *http.Request) {
	htmlPath := "temp/test.html"
	pdfPath := "temp/test.pdf"

	m, err := fileio.DecodeRequestBody(r.Body)
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = fileio.GenerateHTMLFromData(m["data"], htmlPath)
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = fileio.ConvertHTMLtoPDF(htmlPath, pdfPath)
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	pdfContent, err := fileio.GetPdfBytes(pdfPath)
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+"test.pdf")
	w.Header().Set("Content-Type", "application/pdf")
	w.Write(pdfContent)
	w.WriteHeader(http.StatusOK)
}
