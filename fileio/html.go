package fileio

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"os"
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
func GenerateHTMLFromData(data interface{}, htmlPath string) error {

	t, err := template.New("foo").Option("missingkey=error").Parse(`Hello {{.world}}!`)
	if err != nil {
		return err
	}

	file, err := os.Create(htmlPath)
	if err != nil {
		return err
	}
	defer file.Close()

	err = t.Execute(file, data)
	if err != nil {
		return err
	}

	err = file.Sync()
	if err != nil {
		return err
	}

	return nil
}
