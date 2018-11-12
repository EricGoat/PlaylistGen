package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const tokenBaseUrl = "https://accounts.spotify.com/api/token"

var (
	ClientSecret = os.Getenv("SPOTIFY_CLIENT_SECRET")
	ClientID     = os.Getenv("SPOTIFY_CLIENT_ID")
)

type AccessToken struct {
	Value        string `json:"access_token"`
	Type         string `json:"token_type"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

func GetAccessToken(code string, redirectURI string) (AccessToken, error) {
	url := buildTokenURL(code, redirectURI)
	token := fmt.Sprintf("%s%s", ClientID, ClientSecret)
	resp, err := makeRequest(url, token)
	if err != nil {
		return AccessToken{}, nil
	}

	return parseTokenResponse(resp), nil
}

func buildTokenURL(code string, redirectURI string) string {
	return fmt.Sprintf("%s?grant_type=authorization_code&code=%s&redirect_uri=%s",
		tokenBaseUrl,
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

func parseTokenResponse(resp []byte) AccessToken {
	accessToken := AccessToken{}
	json.Unmarshal(resp, &accessToken)

	return accessToken
}
