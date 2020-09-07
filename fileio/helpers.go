package fileio

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
)

// GetTempDirAndPaths creates a temp directory and HTML and PDF filepaths in that order: TempDir, HTMLPath, PDFPath
func GetTempDirAndPaths(filename string) (string, string, string, error) {
	dir, err := ioutil.TempDir("", "temp")
	if err != nil {
		log.Printf(`Error in GetTempDirAndPaths, ERROR: %s :: %s`, err, err.Error())
		return "", "", "", err
	}

	htmlFilePath := filepath.Join(dir, fmt.Sprintf("%s.html", filename))
	pdfFilePath := filepath.Join(dir, fmt.Sprintf("%s.pdf", filename))

	return dir, htmlFilePath, pdfFilePath, nil
}

// CopyAllAssetsToTempDir stages HTML, CSS, and images to the temp directory for templating
func CopyAllAssetsToTempDir(tempDir string, template string) error {
	srcDir := filepath.Join("../templates", template)

	srcFiles, err := ioutil.ReadDir(srcDir)
	if err != nil {
		log.Printf(`Error in CopyAllAssetsToTempDir, ERROR: %s :: %s`, err, err.Error())
		return err
	}

	for _, file := range srcFiles {
		srcFile := path.Join(srcDir, file.Name())
		dstFile := path.Join(tempDir, file.Name())

		in, err := os.Open(srcFile)
		if err != nil {
			log.Printf(`Error in CopyAllAssetsToTempDir, ERROR: %s :: %s`, err, err.Error())
			return err
		}
		defer in.Close()

		out, err := os.Create(dstFile)
		if err != nil {
			log.Printf(`Error in CopyAllAssetsToTempDir, ERROR: %s :: %s`, err, err.Error())
			return err
		}
		defer out.Close()

		_, err = io.Copy(out, in)
		if err != nil {
			log.Printf(`Error in CopyAllAssetsToTempDir, ERROR: %s :: %s`, err, err.Error())
			return err
		}
	}

	return nil
}
