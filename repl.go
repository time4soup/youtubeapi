package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// REPL for CLI user input to query IGDB API
func repl() error {
	fmt.Print("running...\n")

	//Initializes scanner, env vars, IGDB token, and structs for IGDB games and errors
	scanner := bufio.NewScanner(os.Stdin)
	//add: handle Scanner.Err()?
	var input []string

	envVars, err := getEnvVars()
	if err != nil {
		return err
	}

	//add: move requestIgdbToken() and unmarshallJson() to new single func getToken()
	igdbTokenData, _, err := requestIgdbToken(envVars)
	if err != nil {
		return err
	}
	var token IgdbToken
	err = unmarshallJson(igdbTokenData, &token)
	if err != nil {
		return err
	}

	var igdbGames IgdbGames
	var igdbError IgdbError

	//REPL loop for user input to query IGDB API
	for scanner.Scan() {
		//add: reset IgdbGames and IgdbError to blank structs at start of each loop
		//add: parse multiple and multi word args for input, update parseInput()
		input = parseInput(scanner.Text())

		//add: validation of input before use
		igdbReqData, igdbGamesStatusCode, err := fetchIgdbData(fmt.Sprintf(`search "%s"; fields *;`, input[0]), token.AccessToken)
		if err != nil {
			return err
		}
		// fmt.Printf("status code: %d\n", igdbGamesStatusCode) //testing
		if igdbGamesStatusCode != http.StatusOK {
			err = unmarshallJson(igdbReqData, &igdbError)
			if err != nil {
				return err
			}
			fmt.Printf("IGDB req error: %v\n", igdbError)
			continue
		}

		err = unmarshallJson(igdbReqData, &igdbGames)
		if err != nil {
			return err
		}

		fmt.Println("games: ")
		for _, game := range igdbGames {
			//add: format output responsively to fields requested
			fmt.Printf("%s: %d\n", game.Name, int(game.Rating))
		}
	}
	return nil
}

func parseInput(text string) []string {
	return strings.Fields(text)
}
