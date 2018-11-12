package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
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
	body := fmt.Sprintf("grant_type=authorization_code&code=%s&redirect_uri=%s", code, url.PathEscape(redirectURI))
	token := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", ClientID, ClientSecret)))

	resp, err := Post(tokenBaseUrl, body, token)
	if err != nil {
		return AccessToken{}, nil
	}

	return parseTokenResponse(resp), nil
}

func parseTokenResponse(resp []byte) AccessToken {
	accessToken := AccessToken{}
	json.Unmarshal(resp, &accessToken)

	return accessToken
}
