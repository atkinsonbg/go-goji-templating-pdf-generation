package fileio

import (
	"io"
	"encoding/json"
	"log"
)

// DecodeRequestBody takes the http.Request body and decodes it to a map
func DecodeRequestBody(rbody io.ReadCloser) (map[string]interface{}, error) {
	m := map[string]interface{}{}
	err := json.NewDecoder(rbody).Decode(&m)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	return m, nil
}