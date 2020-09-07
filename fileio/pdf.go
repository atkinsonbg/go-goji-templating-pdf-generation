package fileio

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

// ConvertHTMLtoPDF converts the supplied HTML file to a PDF file
func ConvertHTMLtoPDF(htmlFilePath string, pdfFilePath string, pageSize string) error {
	args := []string{"--enable-local-file-access", "--page-size", pageSize, htmlFilePath, pdfFilePath}
	cmd := exec.Command("wkhtmltopdf", args...)
	err := cmd.Run()
	if err != nil {
		log.Printf(`Error in ConvertHTMLtoPDF, ERROR: %s :: %s`, err, err.Error())
		return err
	}
	return nil
}

// OptimizePDF optimizes the PDF, reusing images, and reducing the overall size and applies metadata & bookmarks
func OptimizePDF(pdfPath string, metadata interface{}) (string, error) {
	m := metadata.(map[string]interface{})

	optPdfPath := strings.ReplaceAll(pdfPath, ".pdf", "-opt.pdf")
	argNoPause := "-dNOPAUSE"
	argBatch := "-dBATCH"
	argDevice := "-sDEVICE=pdfwrite"
	argPdfSettings := "-dPDFSETTINGS=/ebook"
	argDuplicateImages := "-dDetectDuplicateImages=true"
	argOutput := fmt.Sprintf(`-sOutputFile="%s"`, optPdfPath)
	argMarks1 := "-c"
	argMarks2 := fmt.Sprintf(`[ /Title (%s) 
						/Author (%s) 
						/Subject (%s) 
						/Keywords (%s) 
						/DOCINFO pdfmark
						
						[ /Title (%s) /Page 1 /OUT pdfmark
						
						[ {Catalog} <</Lang (%s)>> /PUT pdfmark`, m["title"], m["author"], m["subject"], m["keywords"], m["title"], m["language"])
	args := []string{argNoPause, argBatch, argDevice, argPdfSettings, argDuplicateImages, argOutput, pdfPath, argMarks1, argMarks2}

	cmd := exec.Command("gs", args...)
	err := cmd.Run()
	if err != nil {
		log.Printf(`Error in OptimizePDF, ERROR: %s :: %s`, err, err.Error())
		return "ERROR", err
	}
	return optPdfPath, nil
}

// GetPdfBytes returns a []byte of the requested file to return in the http.ResponseWriter
func GetPdfBytes(pdfPath string) ([]byte, error) {
	content, err := ioutil.ReadFile(pdfPath)
	if err != nil {
		log.Printf(`Error in GetPdfBytes, ERROR: %s :: %s`, err, err.Error())
		return nil, err
	}

	if len(content) == 0 {
		return nil, errors.New("content is zero")
	}

	return content, nil
}
