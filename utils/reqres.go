package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"reflect"
)

type Response struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type ResponseMessage struct {
	Message string `json:"message,omitempty"`
}

func JSONResponse(w http.ResponseWriter, statusCode int, data *Response) {

	w.WriteHeader(statusCode)

	if data != nil {
		var err error
		vof := reflect.ValueOf(data.Data)
		if !vof.IsValid() || vof.IsNil() {
			err = json.NewEncoder(w).Encode(&ResponseMessage{Message: data.Message})
		} else {
			err = json.NewEncoder(w).Encode(data)
		}
		if err != nil {
			JSONResponse(w, http.StatusInternalServerError, &Response{"error encoding response data", nil})
			return
		}
	}

}

// DecodeJSONBody decodes a JSON request
func DecodeJSONBody(w http.ResponseWriter, r *http.Request, data interface{}, contentType string, maxBytes int64, disallowUnknownFields bool) error {

	if r.Header.Get("Content-Type") != contentType {
		return errors.New("InvalidContentTypeError")
	}

	r.Body = http.MaxBytesReader(w, r.Body, maxBytes)

	decodedData := json.NewDecoder(r.Body)

	if disallowUnknownFields {
		decodedData.DisallowUnknownFields()
	}

	err := decodedData.Decode(&data)

	if err != nil && err != io.EOF {
		return err
	}

	return nil

}
