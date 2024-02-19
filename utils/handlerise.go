package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	validator "github.com/go-playground/validator/v10"
)

type Model[R any] interface {
	Controller(ctx context.Context) (int, string, R, error)
}

func Handlerize[M Model[R], R any](m M, r R) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqData M

		if r.Method == http.MethodGet || r.Method == http.MethodDelete {
			queryData := map[string]interface{}{}
			queryParams := getQueryKeys(m)
			for _, k := range queryParams {
				queryData[k] = r.URL.Query().Get(k)
			}

			jsonBody, err := json.Marshal(queryData)

			if err != nil {
				JSONResponse(w, http.StatusBadRequest, &Response{Message: "Encode query error: " + err.Error(), Data: nil})
				return
			}
			if err := json.Unmarshal(jsonBody, &reqData); err != nil {
				JSONResponse(w, http.StatusBadRequest, &Response{Message: "Decode query error: " + err.Error(), Data: nil})
				return
			}

		} else {
			if err := DecodeJSONBody(w, r, &reqData, "application/json", 64, true); err != nil {
				JSONResponse(w, http.StatusBadRequest, &Response{Message: "Decode req body error: " + err.Error(), Data: nil})
				return
			}
		}

		if r, err := ValidateData(reqData); err != nil {
			JSONResponse(w, http.StatusBadRequest, r)
			return
		}

		ctx := setcontextWR(r.Context(), r, w)

		status, msg, data, err := reqData.Controller(ctx)
		if status == 0 {
			return
		}

		if err != nil {
			JSONResponse(w, status, &Response{Message: msg, Data: nil})
			return
		}

		if status >= 400 {
			JSONResponse(w, status, &Response{Message: msg, Data: nil})
			return
		}

		if status == http.StatusNoContent {
			JSONResponse(w, status, nil)
			return
		}

		JSONResponse(w, status, &Response{Message: msg, Data: data})
	}
}

func ValidateData(m interface{}) (*Response, error) {
	response := &Response{Message: "validation error", Data: []string{}}
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			response.Data = append(response.Data.([]string), MakeValidationErr(err.Field(), err.Tag(), err.Param(), err.Value().(string)))
		}
		return response, err
	}
	return response, nil
}

func MakeValidationErr(field, expected, param, value string) string {
	return fmt.Sprintf("Field: [%s] Expected:[%s(%s)] Value: [%s]", field, expected, param, value)
}

func getQueryKeys[M any](d M) []string {
	val := reflect.Indirect(reflect.ValueOf(d))
	queryKeys := []string{}
	for i := 0; i < val.Type().NumField(); i++ {
		t := val.Type().Field(i)
		fieldName := t.Name

		switch jsonTag := t.Tag.Get("json"); jsonTag {
		case "-":
		case "":
			queryKeys = append(queryKeys, fieldName)
		default:
			parts := strings.Split(jsonTag, ",")
			name := parts[0]
			if name == "" {
				name = fieldName
			}
			queryKeys = append(queryKeys, name)
		}
	}
	return queryKeys
}
