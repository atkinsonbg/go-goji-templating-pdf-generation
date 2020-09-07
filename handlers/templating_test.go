package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTemplateHandler(t *testing.T) {
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

	b := bytes.NewBuffer(data)

	req, err := http.NewRequest("POST", "/template", b)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(TemplatingHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}
}

func TestTemplateHandlerMissingKey(t *testing.T) {
	var data = []byte(`
	{
		"filename": "starfleet-academy-letter",
		"template": "academy",
		"pagesize": "Letter",
		"data": {
			"firstname": "Brandon",
			"lastname": "Atkinson",
			"address1": "123 South Till St",
			"address22222222": "San Francisco, CA 90034",
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

	b := bytes.NewBuffer(data)

	req, err := http.NewRequest("POST", "/template", b)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(TemplatingHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusInternalServerError)
	}
}

func TestTemplateHandlerNoBody(t *testing.T) {

	var data = []byte(``)

	b := bytes.NewBuffer(data)

	req, err := http.NewRequest("POST", "/template", b)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(TemplatingHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusInternalServerError)
	}
}

func TestTemplateHandlerWrongTemplate(t *testing.T) {
	var data = []byte(`
	{
		"filename": "starfleetnew",
		"template": "template1wrong",
		"optimize": true,
		"data": {
			"firstname": "Brandon",
			"lastname": "Atkinson",
			"examdate": "12/20/3045",
			"replydate": "11/10/3045"
		}
	}`)

	b := bytes.NewBuffer(data)

	req, err := http.NewRequest("POST", "/template", b)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(TemplatingHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusInternalServerError)
	}
}

func BenchmarkGenerateHTMLFromData(b *testing.B) {

	b.StartTimer()
	for i := 0; i < b.N; i++ {
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

		buff := bytes.NewBuffer(data)

		req, _ := http.NewRequest("POST", "/template", buff)
		req.Header.Add("Content-Type", "application/json")
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(TemplatingHandler)

		handler.ServeHTTP(rr, req)
	}
}
