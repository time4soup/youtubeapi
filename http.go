package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// request with context?

const GRANT_TYPE = "client_credentials"
const IGDB_TOKEN_URL = "https://id.twitch.tv/oauth2/token"
const IGDB_GAMES_URL = "https://api.igdb.com/v4/games"

// sends POST req to IGDB server to get data for auth token to be parsed into JSON
func requestIgdbToken(envVars map[string]string) ([]byte, error) {
	client_secret := envVars["IGDB_CLIENT_SECRET"]
	client_id := envVars["IGDB_CLIENT_ID"]

	req_data := url.Values{
		"client_id":     {client_id},
		"client_secret": {client_secret},
		"grant_type":    {GRANT_TYPE},
	}

	data, err := postFormReq(IGDB_TOKEN_URL, req_data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func fetchIgdbData(query, accessToken string) ([]byte, error) {
	clientId, err := getEnvValue("IGDB_CLIENT_ID")
	if err != nil {
		return nil, err
	}
	headers := map[string]string{
		"Client-ID":     clientId,
		"Authorization": "Bearer " + accessToken,
	}

	resData, err := postReq(IGDB_GAMES_URL, query, headers)
	if err != nil {
		return nil, err
	}
	return resData, nil
}

// sends POST req with data URL encoded
func postFormReq(url string, data url.Values) ([]byte, error) {
	resp, err := http.PostForm(url, data)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func postReq(url, body string, headers map[string]string) ([]byte, error) {
	client := http.DefaultClient
	marshalled, err := json.Marshal(&body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(marshalled))
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}
	req.Header.Add("Accept", "application/json")
	fmt.Println(req)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	bodyData, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return bodyData, nil
}
