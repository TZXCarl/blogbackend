package utils

import (
	"encoding/json"
	// "io/ioutil"
	// "log"
	"net/http"
	// "strings"
	// "strconv"
	"fmt"
	// "time"
	"runtime"
	"errors"
)

var errs []JsonapiError

type JsonapiDocument struct {
	Errors []JsonapiError `json:"errors,omitempty"`
}

type JsonapiError struct {
	Code     int               `json:"code,omitempty"`
	Detail   string            `json:"detail"`
	Meta     map[string]interface{} `json:"meta"`
}

func errs2doc(errs []JsonapiError) (string, error) {
	doc := JsonapiDocument{
		Errors: errs,
	}

	b, e := json.Marshal(doc)
	if e != nil {
		return "", errors.New("can not encode error object")
	}

	return string(b), nil;
}

func HandleServerError(w http.ResponseWriter, e error) {
	_, f, l, _ := runtime.Caller(1);
	errData := JsonapiError {
		Code: 500,
		Detail: e.Error(),
		Meta: map[string]interface{}{
			"file": f,
			"line": l,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)
	json, _ := json.Marshal(errData)

	fmt.Fprintln(w, string(json));
}

// 400
func HandleHTTPError(w http.ResponseWriter, err error) {
	_, f, l, _ := runtime.Caller(1)
	errs = append(errs, JsonapiError{
		Code: 400,
		Detail: err.Error(),
		Meta: map[string]interface{}{
			"file": f,
			"line": l,
		},
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)
	json, e := errs2doc(errs)

	if e != nil {
		w.WriteHeader(500)
		fmt.Fprintln(w, e.Error)
	}

	fmt.Fprintln(w, json)
}

func Handle404(w http.ResponseWriter, err error) {
	ResetHTTPErrors()
	w.Header().Set("Content-Type", "applciation/json")
	w.WriteHeader(404)
	// _, f, l, _ := runtime.Caller(1)
	errs = append(errs, JsonapiError{
		Detail: "Not Found",
	})
	json, e := errs2doc(errs)
	if e != nil {
		w.WriteHeader(500)
		fmt.Fprintln(w, e.Error)
	}

	fmt.Fprintln(w, json)
} 

func ResetHTTPErrors() {
	errs = nil;
}