package main

import (
	"encoding/json"
	"net/http"
)

type okResponse struct {
	Ok   bool        `json:"ok"`
	Data interface{} `json:"data"`
}

type errResponse struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
}

func DataController(w http.ResponseWriter, _ *http.Request) {
	sendResponse(okResponse{Ok: true, Data: "Hello, World"}, w)
}

func sendResponse(data interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(okResponse{Ok: true, Data: data})
}

func _(err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(errResponse{Ok: false, Error: err.Error()})
}
