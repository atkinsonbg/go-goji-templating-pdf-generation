package fileio

import (
	"log"
	"path/filepath"
	"testing"
)

func TestGetTempDirAndPathsPass(t *testing.T) {
	dir, htmlPath, pdfPath, err := GetTempDirAndPaths("test")
	if err != nil {
		log.Print("something went wrong with Temp directory creation")
	}

	if len(dir) <= 0 {
		log.Print("something went wrong with Temp directory creation")
	}

	expectedHTMLFilePath := filepath.Join(dir, "test.html")
	expectedPDFFilePath := filepath.Join(dir, "test.pdf")

	if htmlPath != expectedHTMLFilePath {
		t.Error("FAIL: HTML path was not created successfully.")
	}

	if pdfPath != expectedPDFFilePath {
		t.Error("FAIL: PDF path was not created successfully.")
	}
}
