package fileio

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"os"
	"path"
)

// DecodeRequestBody takes the http.Request body and decodes it to a map
func DecodeRequestBody(rbody io.ReadCloser) (map[string]interface{}, error) {
	m := map[string]interface{}{}
	err := json.NewDecoder(rbody).Decode(&m)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	return m, nil
}

// GenerateHTMLFromData takes a map of data and performs the HTML templating
func GenerateHTMLFromData(data interface{}, tempDir string, templateName string, htmlPath string) error {

	err := CopyAllAssetsToTempDir(tempDir, templateName)
	if err != nil {
		log.Print(err)
		return err
	}

	// expecting template to be named "index.html"
	htmlTemplate := path.Join(tempDir, "index.html")

	t := template.New("index.html")
	t, err = t.ParseFiles(htmlTemplate)
	if err != nil {
		log.Print(err)
		return err
	}

	file, err := os.Create(htmlPath)
	if err != nil {
		log.Print(err)
		return err
	}
	defer file.Close()

	err = t.Execute(file, data)
	if err != nil {
		log.Print(err)
		return err
	}

	err = file.Sync()
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}
