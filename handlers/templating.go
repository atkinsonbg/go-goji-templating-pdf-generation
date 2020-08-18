package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/exec"
)

func TemplatingHandler(w http.ResponseWriter, r *http.Request) {
	// Using the Option object here to throw an error if a key is missing in the data

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

	var file, err2 = os.Create("temp/test.html")
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	defer file.Close()
	_, err = file.WriteString("<html><body>Hello WORLD!!</body></html>")
	if err != nil {
		fmt.Println(err2.Error())
	}
	err = file.Sync()
	if err != nil {
		fmt.Println(err2.Error())
	}

	err = convert()
	if err != nil {
		fmt.Println(err.Error())
	}

	//var err = os.Remove(path)

	w.WriteHeader(http.StatusOK)
}

func convert() error {
	args := []string{"temp/test.html", "temp/test.pdf"}
	cmd := exec.Command("wkhtmltopdf", args...)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
