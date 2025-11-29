package main

import (
	"encoding/json"
)

func unmarshallJson(data []byte, v any) error {
	err := json.Unmarshal(data, v)
	if err != nil {
		return err
	}
	return nil
}
