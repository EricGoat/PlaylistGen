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

func (s *Server) handleGenres() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		accessToken, err := GetAccessToken(code, redirectURI)
		if err != nil {
			sendError(err, w)
		}

		genres, err := GetGenres(accessToken)
		if err != nil {
			sendError(err, w)
		}

		sendResponse(genres, w)
	}
}

func (s *Server) handleClientID() http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		sendResponse(ClientID, w)
	}
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
