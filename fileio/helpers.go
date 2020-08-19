package fileio

import (
	"io/ioutil"
	"path/filepath"
	"fmt"
)

// GetTempDirAndPaths creates a temp directory and HTML and PDF filepaths in that order: TempDir, HTMLPath, PDFPath
func GetTempDirAndPaths(filename string) (string, string, string, error) {
	dir, err := ioutil.TempDir("", "temp")
	if err != nil {
		return "", "", "", err
	}

	htmlFilePath := filepath.Join(dir, fmt.Sprintf("%s.html", filename))
	pdfFilePath := filepath.Join(dir, fmt.Sprintf("%s.pdf", filename))

	return dir, htmlFilePath, pdfFilePath, nil
}