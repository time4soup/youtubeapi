package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

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

	igdbTokenData, igdbTokenStatusCode, err := requestIgdbToken(envVars)
	if err != nil {
		return err
	}
	if igdbTokenStatusCode != http.StatusOK {
		return fmt.Errorf("http error: %d", igdbTokenStatusCode)
	}
	var token IgdbToken
	err = unmarshallJson(igdbTokenData, &token)
	if err != nil {
		return err
	}

	var igdbGames IgdbGames
	var igdbError IgdbError
	for scanner.Scan() {
		input = parseInput(scanner.Text())

		igdbReqData, igdbGamesStatusCode, err := fetchIgdbData(fmt.Sprintf(`search "%s"; fields *;`, input[0]), token.AccessToken)
		fmt.Printf("status code: %d\n", igdbGamesStatusCode)
		if err != nil {
			return err
		}
		if igdbGamesStatusCode != http.StatusOK {
			err = unmarshallJson(igdbReqData, &igdbError)
			if err != nil {
				fmt.Println(string(igdbReqData[:]))
				return err
			}
			fmt.Printf("IGDB req error: %v\n", igdbError)
			// return fmt.Errorf("IGDB Status Code: %d", igdbGamesStatusCode)
			continue
		}

		err = unmarshallJson(igdbReqData, &igdbGames)
		if err != nil {
			return err
		}

		fmt.Println("games: ")
		for _, game := range igdbGames {
			fmt.Printf("%s: %d\n", game.Name, int(game.Rating))
		}
	}
	return nil
}

func parseInput(text string) []string {
	return strings.Fields(text)
}
