package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	initCommands()
	input := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		input.Scan()
		args := cleanInput(input.Text())
		if cmd, ok := CommandRegistry[args[0]]; ok {
			cmd.callback(config{})
		} else {
			fmt.Println("Unknown command")
		}
	}
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
