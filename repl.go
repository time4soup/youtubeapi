package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	// "github.com/gofor-little/env"
)

const IGDB_URL = "https://id.twitch.tv/oauth2/token"

func repl() error {
	// if err := scanner.Err(); err != nil {
	// 	fmt.Fprintln(os.Stderr, "error reading: ", err)
	// }

	fmt.Print("running...\n")

	scanner := bufio.NewScanner(os.Stdin)
	var input []string
	envVars, err := getEnvVars()
	if err != nil {
		return err
	}

	for scanner.Scan() {
		input = parseInput(scanner.Text())

		data, err := postReq(input, IGDB_URL, envVars)
		if err != nil {
			return err
		}
		var token IgdbToken
		unmarshallJson(data, &token)

		fmt.Println("yep!")
		fmt.Printf("token: %s\n", token.AccessToken)
		fmt.Printf("type: %s\n", token.TokenType)
	}
	return nil
}

func parseInput(text string) []string {
	return strings.Fields(text)
}
