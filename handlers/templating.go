package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/exec"

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

	err = convert(htmlPath, pdfPath)
	if err != nil {
		fmt.Println(err.Error())
	}

	//Send the headers
	w.Header().Set("Content-Disposition", "attachment; filename="+"pdffromapi.pdf")
	w.Header().Set("Content-Type", "application/pdf")

	content, errc := fileio.GetPdfBytes(pdfPath)
	if errc != nil {
		fmt.Println("FJDKSLJKDSLJFJSDLFKJ")
	}

	//buf := bytes.NewBuffer(content)
	//buf.WriteTo(w)

	w.Write(content)

	w.WriteHeader(http.StatusOK)
}

func convert(htmlFilePath string, pdfFilePath string) error {
	args := []string{htmlFilePath, pdfFilePath}
	cmd := exec.Command("wkhtmltopdf", args...)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

// func getPDF(filename string) error {
// 	Openfile, err := os.Open(filename)
// 	defer Openfile.Close() //Close after function return
// 	if err != nil {
// 		//File not found, send 404
// 		http.Error(writer, "File not found.", 404)
// 		return
// 	}

// 	//File is found, create and send the correct headers

// 	//Get the Content-Type of the file
// 	//Create a buffer to store the header of the file in
// 	FileHeader := make([]byte, 512)
// 	//Copy the headers into the FileHeader buffer
// 	Openfile.Read(FileHeader)
// 	//Get content type of file
// 	FileContentType := http.DetectContentType(FileHeader)

// 	//Send the headers
// 	writer.Header().Set("Content-Disposition", "attachment; filename="+Filename)
// 	writer.Header().Set("Content-Type", FileContentType)

// 	//Send the file
// 	//We read 512 bytes from the file already, so we reset the offset back to 0
// 	Openfile.Seek(0, 0)
// 	io.Copy(writer, Openfile)
// 	return nil
// }
