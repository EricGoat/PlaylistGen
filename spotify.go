package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const TOKEN_BASE_URL = "https://accounts.spotify.com/api/token"
const CLIENT_ID = "725c53094c0449379beb24431dda70cf"

func GetAccessToken(code string, redirectURI string) (string, error) {
	url := buildTokenURL(code, redirectURI)
	token := fmt.Sprintf("%s%s", CLIENT_ID, getClientSecret())
	resp, err := makeRequest(url, token)
	if err != nil {
		return "", nil
	}

	// TODO: parse resp into json and get access_token from it
}

func buildTokenURL(code string, redirectURI string) string {
	return fmt.Sprintf("%s?grant_type=authorization_code&code=%s&redirect_uri=%s",
		TOKEN_BASE_URL,
		code,
		redirectURI)
}

func makeRequest(url string, token string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func getClientSecret() string {
	return os.Getenv("SPOTIFY_CLIENT_SECRET")
}
