package main

import (
	"bufio"
	"fmt"
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

	igdbTokenData, err := requestIgdbToken(envVars)
	if err != nil {
		return err
	}
	var token IgdbToken
	err = unmarshallJson(igdbTokenData, &token)
	if err != nil {
		return err
	}

	for scanner.Scan() {
		input = parseInput(scanner.Text())

		fmt.Printf("token: %s\n\n", token.AccessToken)

		igdbReqData, err := fetchIgdbData(fmt.Sprintf(`search "%s"; fields *;`, input[0]), token.AccessToken)
		if err != nil {
			return err
		}
		var igdbGames IgdbGames
		err = unmarshallJson(igdbReqData, &igdbGames)
		if err != nil {
			return err
		}

		for _, game := range igdbGames {
			if game.Name != "" {
				fmt.Println(game.Name)
			}
		}
	}
	return nil
}

func parseInput(text string) []string {
	return strings.Fields(text)
}
