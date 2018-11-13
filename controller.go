package main

import (
	"encoding/json"
	"errors"
	"fmt"
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

func (s *server) handleSongs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accessToken, err := getAccessTokenFromRequest(r)
		if err != nil {
			sendError(err, w)
			return
		}

		genre := r.URL.Query().Get("genre")
		if genre == "" {
			sendError(errors.New("genre parameter is blank"), w)
			return
		}

		duration := r.URL.Query().Get("duration")
		if duration == "" {
			sendError(errors.New("duration parameter is blank"), w)
			return
		}

		songs, err := getSongIDs(genre, duration, accessToken)
		if err != nil {
			sendError(err, w)
			return
		}

		for _, song := range songs {
			fmt.Println(song)
		}

		sendResponse("NOT IMPLEMENTED", w)
	}
}

func (s *server) handleGenres() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accessToken, err := getAccessTokenFromRequest(r)
		if err != nil {
			sendError(err, w)
			return
		}

		genres, err := getGenres(accessToken)
		if err != nil {
			sendError(err, w)
		}

		sendResponse(genres, w)
	}
}

func (s *server) handleClientID() http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		sendResponse(ClientID, w)
	}
}

func getAccessTokenFromRequest(r *http.Request) (AccessToken, error) {
	code := r.URL.Query().Get("code")
	if code == "" {
		return AccessToken{}, errors.New("code parameter is blank")
	}

	redirectURI := r.URL.Query().Get("redirect_uri")
	if redirectURI == "" {
		return AccessToken{}, errors.New("redirect_uri parameter is blank")
	}

	accessToken, err := GetAccessToken(code, redirectURI)
	if err != nil {
		return AccessToken{}, err
	}

	return accessToken, nil
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
