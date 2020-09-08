package fileio

import (
	"bytes"
	"encoding/json"
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
	var data = []byte(`
	{
		"filename": "starfleet-academy-letter",
		"template": "academy",
		"pagesize": "Letter",
		"data": {
			"firstname": "Brandon",
			"lastname": "Atkinson",
			"address1": "123 South Till St",
			"address2": "San Francisco, CA 90034",
			"address3": "United States, Earth",
			"examdate": "12/20/3045",
			"replydate": "11/10/3045"
		},
		"metadata": {
			"language": "en-US",
			"title": "New Metadata Strcuture",
			"author": "Brandon Atkinson",
			"subject": "Awesome Subject",
			"keywords": "Keyword 1, keyword 2"
		}
	}`)
	m := map[string]interface{}{}
	_ = json.Unmarshal(data, &m)

	tempdir, htmlPath, _, _ := GetTempDirAndPaths("unittest")

	err := GenerateHTMLFromData(m["data"], tempdir, "academy", htmlPath)
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
	var data = []byte(`
	{
		"filename": "starfleet-academy-letter",
		"template": "academy",
		"pagesize": "Letter",
		"data": {
			"firstname": "Brandon",
			"lastname": "Atkinson",
			"address1": "123 South Till St",
			"address2": "San Francisco, CA 90034",
			"address3": "United States, Earth",
			"examdate": "12/20/3045",
			"replydate": "11/10/3045"
		},
		"metadata": {
			"language": "en-US",
			"title": "New Metadata Strcuture",
			"author": "Brandon Atkinson",
			"subject": "Awesome Subject",
			"keywords": "Keyword 1, keyword 2"
		}
	}`)
	m := map[string]interface{}{}
	_ = json.Unmarshal(data, &m)

	tempdir, htmlPath, _, _ := GetTempDirAndPaths("unittest")

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = GenerateHTMLFromData(m["data"], tempdir, "academy", htmlPath)
	}
}
