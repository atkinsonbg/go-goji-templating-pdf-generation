package fileio

import (
	"testing"
)

func TestConvertHTMLtoPDFFail(t *testing.T) {
	err := ConvertHTMLtoPDF("test/test.html", "test/test.pdf")
	if err != nil {
		t.Log("PASS: Did not convert HTML.")
		return
	}

	t.Error("FAIL: Ran somehow.")
	return
}

func TestGetPDFBytesFail(t *testing.T) {
	_, err := GetPdfBytes("foo.pdf")
	if err != nil {
		t.Log("PASS: Did not find a PDF.")
		return
	}

	t.Error("FAIL: Found a PDF.")
	return
}

func TestGetPDFBytes(t *testing.T) {
	content, err := GetPdfBytes("../test/test.pdf")
	if err != nil {
		t.Error("Fail: Did not get PDF bytes.")
		return
	}

	if len(content) > 0 {
		t.Log("PASS: Succsefully got PDF bytes.")
	}
	return
}
