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
		"filename": "starfleetnew",
		"template": "template1",
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

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}
}

func TestTemplateHandlerMissingKey(t *testing.T) {
	var data = []byte(`
	{
		"filename": "starfleetnew",
		"template": "template1",
		"optimize": true,
		"data": {
			"wrongkey": "Brandon"
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
