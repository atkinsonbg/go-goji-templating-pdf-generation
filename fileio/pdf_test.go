package fileio

import (
	"testing"
)

func TestConvertHTMLtoPDF(t *testing.T) {
	err := ConvertHTMLtoPDF("templates/test.html", "templates/test.pdf")
	if err != nil {
		t.Error("FAIL: Did not convert HTML.")
		return
	}

	t.Log("PASS: was able to convert the HTML to PDF.")
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
	content, err := GetPdfBytes("../templates/test.pdf")
	if err != nil {
		t.Error("Fail: Did not get PDF bytes.")
		return
	}

	if len(content) > 0 {
		t.Log("PASS: Succsefully got PDF bytes.")
	}
	return
}

func TestAddPDFMetadata(t *testing.T) {
	err := AddPDFMetadata("title", "author", "keywords, more", "subject", "../templates/test.pdf")
	if err != nil {
		t.Error("Fail: Did not add PDF metadata.")
		return
	}
	t.Log("PASS: Succsefully added PDF metadata.")
	return
}

func BenchmarkConvertHTMLtoPDF(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = ConvertHTMLtoPDF("templates/test.html", "templates/test.pdf")
	}
}

func BenchmarkAddPDFMetadata(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = AddPDFMetadata("title", "author", "keywords, more", "subject", "../templates/test.pdf")
	}
}
