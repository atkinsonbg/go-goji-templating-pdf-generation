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
func ConvertHTMLtoPDF(htmlFilePath string, pdfFilePath string) error {
	args := []string{"--enable-local-file-access", "--outline", htmlFilePath, pdfFilePath}
	cmd := exec.Command("wkhtmltopdf", args...)
	err := cmd.Run()
	if err != nil {
		log.Print(err)
		log.Print(err.Error())
		return err
	}
	return nil
}

// OptimizePDF optimizes the PDF, reusing images, and reducing the overall size and applies metadata & bookmarks
func OptimizePDF(pdfPath string) (string, error) {
	optPdfPath := strings.ReplaceAll(pdfPath, ".pdf", "-opt.pdf")
	argNoPause := "-dNOPAUSE"
	argBatch := "-dBATCH"
	argDevice := "-sDEVICE=pdfwrite"
	argPdfSettings := "-dPDFSETTINGS=/ebook"
	argDuplicateImages := "-dDetectDuplicateImages=true"
	argOutput := fmt.Sprintf(`-sOutputFile="%s"`, optPdfPath)
	argMarks1 := "-c"
	argMarks2 := `[ /Title (Jaziels Important Document) 
						/Author (Jaziel Aguirre)
						/DOCINFO pdfmark
						
						[ /Title (Contents) /Page 1 /OUT pdfmark
						
						[ /Subtype /Catalog /Lang (en-US) /StPNE pdfmark`
	args := []string{argNoPause, argBatch, argDevice, argPdfSettings, argDuplicateImages, argOutput, pdfPath, argMarks1, argMarks2}

	cmd := exec.Command("gs", args...)
	err := cmd.Run()
	if err != nil {
		log.Print(err)
		return "ERROR", err
	}
	return optPdfPath, nil
}

// GetPdfBytes returns a []byte of the requested file to return in the http.ResponseWriter
func GetPdfBytes(pdfPath string) ([]byte, error) {
	content, err := ioutil.ReadFile(pdfPath)
	if err != nil {
		return nil, err
	}

	if len(content) == 0 {
		return nil, errors.New("content is zero")
	}

	return content, nil
}
