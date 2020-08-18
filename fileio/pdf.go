package fileio

import (
	"io/ioutil"
	"log"
)

func GetPdfBytes(pdfPath string) ([]byte, error) {
	content, err := ioutil.ReadFile(pdfPath)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return content, nil
}
