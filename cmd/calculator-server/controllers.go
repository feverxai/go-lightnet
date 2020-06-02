package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type Request struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

type Response struct {
	Result *float64 `json:"result,omitempty"`
	Error  *string  `json:"error,omitempty"`
}

func Controller(w http.ResponseWriter, r *http.Request) {
	resp := new(Response)
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg := fmt.Sprintf("ioutil.ReadAll() error : %v", err)
		resp.Error = &msg
		response(w, http.StatusBadRequest, resp)
		return
	}

	request := new(Request)
	err = json.Unmarshal(bodyBytes, request)
	if err != nil {
		msg := fmt.Sprintf("json.Unmarshal error : %v", err)
		resp.Error = &msg
		response(w, http.StatusBadRequest, resp)
		return
	}

	operand := mux.Vars(r)["operand"]
	switch operand {
	case "sum":
		resp.Result = sum(request)
	case "mul":
		resp.Result = mul(request)
	case "sub":
		resp.Result = sub(request)
	case "div":
		if request.B != 0 {
			resp.Result = div(request)
		} else {
			msg := fmt.Sprintf("division by zero")
			resp.Error = &msg
			response(w, http.StatusBadRequest, resp)
			return
		}
	default:
		msg := fmt.Sprintf("division by zero")
		resp.Error = &msg
		response(w, http.StatusBadRequest, resp)
		return
	}
	response(w, http.StatusOK, resp)
	return
}

func sum(request *Request) *float64 {
	result := request.A + request.B
	return &result
}

func mul(request *Request) *float64 {
	result := request.A * request.B
	return &result
}

func sub(request *Request) *float64 {
	result := request.A - request.B
	return &result
}

func div(request *Request) *float64 {
	result := request.A / request.B
	return &result
}

func response(w http.ResponseWriter, header int, data interface{}) {
	jData, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(header)
	w.Write(jData)
}
