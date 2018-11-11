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
	sendResponse("Hello, World!", w)
}

func GenresController(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	redirectURL := r.URL.Query().Get("redirect_url")
	accessToken, err := GetAccessToken(code, redirectURL)
	if err != nil {
		sendError(err, w)
	}
}

func sendResponse(data interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(okResponse{Ok: true, Data: data})
}

func sendError(err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(errResponse{Ok: false, Error: err.Error()})
}
