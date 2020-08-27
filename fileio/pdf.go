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
	args := []string{htmlFilePath, pdfFilePath}
	cmd := exec.Command("wkhtmltopdf", args...)
	err := cmd.Run()
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

// AddPDFMetadata adds metadata to the PDF file, this is for accessibility
func AddPDFMetadata(title string, author string, keywords string, subject string, pdfPath string) error {
	argTitle := fmt.Sprintf(`-Title="%s"`, title)
	argAuthor := fmt.Sprintf(`-PDF:Author="%s"`, author)
	argKeywords := fmt.Sprintf(`-PDF:Keywords="%s"`, keywords)
	argSubject := fmt.Sprintf(`-PDF:Subject="%s"`, subject)
	argOverwrite := "-overwrite_original_in_place"
	argPath := pdfPath
	args := []string{argTitle, argAuthor, argKeywords, argSubject, argOverwrite, argPath}

	cmd := exec.Command("exiftool", args...)
	err := cmd.Run()
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

// OptimizePDF optimizes the PDF, reusing images, and reducing the overall size
func OptimizePDF(pdfPath string) (string, error) {
	optPdfPath := strings.ReplaceAll(pdfPath, ".pdf", "-opt.pdf")
	argNoPause := "-dNOPAUSE"
	argBatch := "-dBATCH"
	argDevice := "-sDEVICE=pdfwrite"
	argPdfSettings := "-dPDFSETTINGS=/printer"
	argDuplicateImages := "-dDetectDuplicateImages=true"
	argOutput := fmt.Sprintf(`-sOutputFile="%s"`, optPdfPath)
	args := []string{argNoPause, argBatch, argDevice, argPdfSettings, argDuplicateImages, argOutput, pdfPath}

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
