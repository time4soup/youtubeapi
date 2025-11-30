package main

import (
	"github.com/gofor-little/env"
)

var igdbEnvKeys = []string{"IGDB_CLIENT_SECRET", "IGDB_CLIENT_ID"}

func getEnvVars() (map[string]string, error) {
	igdbEnvVals := make(map[string]string)

	err := env.Load(".env")
	if err != nil {
		return nil, err
	}

	var val string
	for _, key := range igdbEnvKeys {
		val, err = env.MustGet(key)
		if err != nil {
			return nil, err
		}
		igdbEnvVals[key] = val
	}
	return igdbEnvVals, nil
}

// fetches value of .env variable for given key
func getEnvValue(key string) (string, error) {
	err := env.Load(".env")
	if err != nil {
		return "", err
	}

	val, err := env.MustGet(key)
	if err != nil {
		return "", err
	}
	return val, nil
}
