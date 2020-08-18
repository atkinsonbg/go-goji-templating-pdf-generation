package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/atkinsonbg/go-goji-templating-pdf-generation/fileio"
)

func TemplatingHandler(w http.ResponseWriter, r *http.Request) {
	htmlPath := "temp/test.html"
	pdfPath := "temp/test.pdf"

	t, err := template.New("foo").Option("missingkey=error").Parse(`Hello {{.world}}!`)
	if err != nil {
		fmt.Println(err)
	}

	// Creating the JSON string and marshaling it to a map

	jsondata := `{ "world": "Earth" }`
	m := map[string]interface{}{}
	err = json.Unmarshal([]byte(jsondata), &m)
	if err != nil {
		fmt.Println(err)
	}

	err = t.Execute(os.Stdout, m)
	if err != nil {
		fmt.Println(err)
	}

	err = os.Mkdir("temp", 0755)

	var file, err2 = os.Create(htmlPath)
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	defer file.Close()
	_, err = file.WriteString("<html><body>Hello WORLD!! sdfsdfdsfsdfsdf<br />fsdfsd</body></html>")
	if err != nil {
		fmt.Println(err2.Error())
	}
	err = file.Sync()
	if err != nil {
		fmt.Println(err2.Error())
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
