package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const IGDB_TOKEN_URL = "https://id.twitch.tv/oauth2/token"
const IGDB_GAMES_URL = "https://api.igdb.com/v4/games"

// sends POST req to IGDB server to get data for auth token to be parsed into JSON
func requestIgdbToken(envVars map[string]string) ([]byte, int, error) {
	client_secret := envVars["IGDB_CLIENT_SECRET"]
	client_id := envVars["IGDB_CLIENT_ID"]

	req_data := url.Values{
		"client_id":     {client_id},
		"client_secret": {client_secret},
		"grant_type":    {"client_credentials"},
	}

	data, statusCode, err := postFormReq(IGDB_TOKEN_URL, req_data)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if statusCode != http.StatusOK {
		return data, statusCode, fmt.Errorf("http error: %d", statusCode)
	}
	return data, statusCode, nil
}

// sends POST req to IGDB server to get games data based on user query
func fetchIgdbData(query, accessToken string) ([]byte, int, error) {
	clientId, err := getEnvValue("IGDB_CLIENT_ID")
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	headers := map[string]string{
		"Client-ID":     clientId,
		"Authorization": "Bearer " + accessToken,
		"Accept":        "application/json",
	}

	resData, statusCode, err := postReq(IGDB_GAMES_URL, query, headers)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return resData, statusCode, nil
}

// sends POST req with data URL encoded
func postFormReq(url string, data url.Values) ([]byte, int, error) {
	res, err := http.PostForm(url, data)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return body, res.StatusCode, nil
}

// sends POST req with given headers and body
func postReq(url, body string, headers map[string]string) ([]byte, int, error) {
	client := http.DefaultClient
	marshalled := []byte(body)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(marshalled))
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}
	// fmt.Println("request:") //debugging
	// fmt.Println(req)

	res, err := client.Do(req)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	defer res.Body.Close()
	bodyData, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return bodyData, res.StatusCode, nil
}
