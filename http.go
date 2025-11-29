package main

import (
	"io"
	"net/http"
	"net/url"
)

// request with context?

const GRANT_TYPE = "client_credentials"

func postReq(data []string, urlString string, envVars map[string]string) ([]byte, error) {
	client_secret := envVars["IGDB_CLIENT_SECRET"]
	client_id := envVars["IGDB_CLIENT_ID"]

	urlData := url.Values{
		"client_id":     {client_id},
		"client_secret": {client_secret},
		"grant_type":    {GRANT_TYPE},
	}
	resp, err := http.PostForm(urlString, urlData)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
