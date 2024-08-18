package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Pokedex!")

	for {
		fmt.Print("pokedex >")

		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}
		commandName := cleaned[0]

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("invalid command")
			continue
		}
		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints Help Menu",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits Pokedex",
			callback:    callbackExit,
		},
		"map": {
			name:        "map",
			description: "Lists the next page of location areas",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists the prevoius page of location areas",
			callback:    callbackMapb,
		},
	}
}

func cleanInput(input string) []string {
	loweredInput := strings.ToLower(input)
	words := strings.Fields(loweredInput)
	return words
}
