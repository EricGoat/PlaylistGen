package main

import (
	"encoding/json"
	"errors"
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
	sendResponse("Hello, World!", w)
}

func ClientIDController(w http.ResponseWriter, _ *http.Request) {
	sendResponse(ClientID, w)
}

func GenresController(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		sendError(errors.New("code parameter is blank"), w)
		return
	}

	redirectURI := r.URL.Query().Get("redirect_uri")
	if redirectURI == "" {
		sendError(errors.New("redirect_uri parameter is blank"), w)
		return
	}

	_, err := GetAccessToken(code, redirectURI)
	if err != nil {
		sendError(err, w)
	}

	sendResponse("Hello, World!", w)
}

func sendResponse(data interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(okResponse{Ok: true, Data: data})
}

func sendError(err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(errResponse{Ok: false, Error: err.Error()})
}
