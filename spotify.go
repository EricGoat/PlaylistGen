package main

import (
	"encoding/json"
)

const genresURL = "https://api.spotify.com/v1/recommendations/available-genre-seeds"

type genresResponse struct {
	Genres []string `json:"genres"`
}

func GetGenres(token AccessToken) ([]string, error) {
	resp, err := Get(genresURL, token.Value)
	if err != nil {
		return nil, err
	}

	genres := genresResponse{}
	json.Unmarshal(resp, &genres)

	return genres.Genres, nil
}
