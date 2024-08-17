package main

import "fmt"

func callbackHelp() error {
	fmt.Println("Pokedex Help Menu")
	fmt.Println("Commands:")

	availableCommands := getCommands()
	for _, cmd := range availableCommands {
		fmt.Printf(" - %s: %s\n", cmd.name, cmd.description)
	}

	return nil
}