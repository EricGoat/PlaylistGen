package main

import (
	"encoding/json"
	"fmt"
)

const genresURL = "https://api.spotify.com/v1/recommendations/available-genre-seeds"
const recommendationsURL = "https://api.spotify.com/v1/recommendations"

type genresResponse struct {
	Genres []string `json:"genres"`
}

type recommendationsResponse struct {
	Tracks []song `json:"tracks"`
}

type song struct {
	ID string `json:"id"`
}

func getSongIDs(genre string, duration string, token AccessToken) ([]string, error) {
	// FIXME: implement logic for finding songs to fill duration
	url := fmt.Sprintf("%s?seed_genres=%s", recommendationsURL, genre)
	resp, err := Get(url, token.Value)
	if err != nil {
		return nil, err
	}

	jsonResp := recommendationsResponse{}
	json.Unmarshal(resp, &jsonResp)
	ids := parseSongIDs(jsonResp)

	return ids, nil
}

func getGenres(token AccessToken) ([]string, error) {
	resp, err := Get(genresURL, token.Value)
	if err != nil {
		return nil, err
	}

	genres := genresResponse{}
	json.Unmarshal(resp, &genres)

	return genres.Genres, nil
}

func parseSongIDs(resp recommendationsResponse) []string {
	tracks := resp.Tracks
	ids := make([]string, len(tracks))
	for _, song := range tracks {
		ids = append(ids, song.ID)
	}

	return ids
}
