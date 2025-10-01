package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")
	for input.Scan() {
		args := cleanInput(input.Text())
		if args[0] == "q" {
			fmt.Println("Exiting...")
			break
		}
		fmt.Printf("Command entered: %s\n", args[0])
		fmt.Print("Pokedex > ")
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
