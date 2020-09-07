package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/atkinsonbg/go-goji-templating-pdf-generation/fileio"
)

// TemplatingHandler is responsible for taking input from the HTTP call, executing the template, and returning a PDF
func TemplatingHandler(w http.ResponseWriter, r *http.Request) {

	m, err := fileio.DecodeRequestBody(r.Body)
	if err != nil {
		log.Printf(`Error in TemplatingHandler, ERROR: %s :: %s`, err, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	filename := m["filename"].(string)
	dir, htmlPath, pdfPath, err := fileio.GetTempDirAndPaths(filename)
	if err != nil {
		log.Printf(`Error in TemplatingHandler, ERROR: %s :: %s`, err, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer os.RemoveAll(dir)

	templateName := m["template"].(string)
	err = fileio.GenerateHTMLFromData(m["data"], dir, templateName, htmlPath)
	if err != nil {
		log.Printf(`Error in TemplatingHandler, ERROR: %s :: %s`, err, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = fileio.ConvertHTMLtoPDF(htmlPath, pdfPath)
	if err != nil {
		log.Printf(`Error in TemplatingHandler, ERROR: %s :: %s`, err, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	pdfPath, err = fileio.OptimizePDF(pdfPath, m["metadata"])
	if err != nil {
		log.Printf(`Error in TemplatingHandler, ERROR: %s :: %s`, err, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	pdfContent, err := fileio.GetPdfBytes(pdfPath)
	if err != nil {
		log.Printf(`Error in TemplatingHandler, ERROR: %s :: %s`, err, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.pdf", filename))
	w.Header().Set("Content-Type", "application/pdf")
	w.Write(pdfContent)
	w.WriteHeader(http.StatusOK)
}
