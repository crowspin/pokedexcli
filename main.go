package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var CommandRegistry map[string]cliCommand

func main() {
	CommandRegistry = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}

	input := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		input.Scan()
		args := cleanInput(input.Text())
		if cmd, ok := CommandRegistry[args[0]]; ok {
			cmd.callback()
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, v := range CommandRegistry {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	return nil
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func cleanInput(text string) []string {
	operableText := strings.ToLower(text)
	stringSlice := strings.Split(operableText, " ")

	var rv []string
	for _, val := range stringSlice {
		if val != "" {
			rv = append(rv, val)
		}
	}
	return rv
}
