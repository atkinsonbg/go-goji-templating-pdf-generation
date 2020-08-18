package fileio

import (
	"io/ioutil"
	"log"
	"os/exec"
)

// ConvertHTMLtoPDF converts the supplied HTML file to a PDF file
func ConvertHTMLtoPDF(htmlFilePath string, pdfFilePath string) error {
	args := []string{htmlFilePath, pdfFilePath}
	cmd := exec.Command("wkhtmltopdf", args...)
	err := cmd.Run()
	if err != nil {
		log.Print(err.Error())
		return err
	}
	return nil
}

// GetPdfBytes returns a []byte of the requested file to return in the http.ResponseWriter
func GetPdfBytes(pdfPath string) ([]byte, error) {
	content, err := ioutil.ReadFile(pdfPath)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return content, nil
}
