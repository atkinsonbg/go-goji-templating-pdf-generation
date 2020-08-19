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