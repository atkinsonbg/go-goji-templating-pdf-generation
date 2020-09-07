package fileio

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestDecodeRequestBodyFail(t *testing.T) {
	r := ioutil.NopCloser(bytes.NewReader([]byte("this will fail")))
	defer r.Close()
	_, err := DecodeRequestBody(r)
	if err != nil {
		t.Log("PASS: Did not convert Body.")
		return
	}

	t.Error("FAIL: Converted body somehow.")
	return
}

func TestDecodeRequestBodyPass(t *testing.T) {
	r := ioutil.NopCloser(bytes.NewReader([]byte(`{"this": "works"}`)))
	defer r.Close()
	m, err := DecodeRequestBody(r)
	if err != nil {
		t.Error("FAIL: Did not convert Body.")
		return
	}

	if m["this"] != "works" {
		t.Error("Pass: Converted body and we can query it.")
	}

	t.Log("PASS: Decoded body and we could read it.")

	return
}

func TestGenerateHTMLFromDataPass(t *testing.T) {
	m := map[string]interface{}{}
	m["firstname"] = "Brandon"
	m["lastname"] = "Atkinson"
	m["examdate"] = "2/3/4"
	m["replydate"] = "1/2/3"

	tempdir, htmlPath, _, _ := GetTempDirAndPaths("unittest")

	err := GenerateHTMLFromData(m, tempdir, "academy", htmlPath)
	if err != nil {
		t.Error("FAIL: generating HTML was not successful")
		return
	}

	t.Log("PASS: was able to generate HTML")
}

func TestGenerateHTMLFromDataFail(t *testing.T) {
	m := map[string]interface{}{}
	m["doesnotexist"] = "Brandon"

	tempdir, htmlPath, _, _ := GetTempDirAndPaths("unittest")

	err := GenerateHTMLFromData(m, tempdir, "academy", htmlPath)
	if err != nil {
		t.Log("PASS: missingkey option caught error")
		return
	}

	t.Error("FAIL: generating HTML should not have been successful, missingkey error")
}

func BenchmarkGenerateHTMLFromData(b *testing.B) {
	m := map[string]interface{}{}
	m["firstname"] = "Brandon"
	m["lastname"] = "Atkinson"
	m["examdate"] = "2/3/4"
	m["replydate"] = "1/2/3"

	tempdir, htmlPath, _, _ := GetTempDirAndPaths("unittest")

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = GenerateHTMLFromData(m, tempdir, "academy", htmlPath)
	}
}
