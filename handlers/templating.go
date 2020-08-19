package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/atkinsonbg/go-goji-templating-pdf-generation/fileio"
)

// TemplatingHandler is responsible for taking input from the HTTP call, executing the template, and returning a PDF
func TemplatingHandler(w http.ResponseWriter, r *http.Request) {
	htmlPath := "temp/test2.html"
	pdfPath := "temp/test2.pdf"

	t, err := template.New("foo").Option("missingkey=error").Parse(`Hello {{.world}}!`)
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	m, err := fileio.DecodeRequestBody(r.Body)
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = os.Mkdir("temp", 0755)

	var file, err2 = os.Create(htmlPath)
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	defer file.Close()

	err = t.Execute(file, m["data"])
	if err != nil {
		fmt.Println(err)
	}

	err = file.Sync()
	if err != nil {
		fmt.Println(err.Error())
	}

	err = fileio.ConvertHTMLtoPDF(htmlPath, pdfPath)
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var content []byte
	content, err = fileio.GetPdfBytes(pdfPath)
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+"pdffromapi.pdf")
	w.Header().Set("Content-Type", "application/pdf")
	w.Write(content)
	w.WriteHeader(http.StatusOK)
}
