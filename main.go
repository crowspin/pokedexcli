package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/crowspin/pokeapi"
	"github.com/crowspin/pokecache"
)

/*---------------------------------------------------------------\
|	Ideas for project extension from lesson (that we might do)	 |
\----------------------------------------------------------------/
<done>Refactor your code to organize it better and make it more testable</done>
Update the CLI to support the "up/down" arrow to cycle through previous commands
Use the PokeAPI to make exploration more interesting. For example, rather than typing the names of areas, maybe you are given choices of areas and just type "left" or "right"
	For starts: we could use the location api to group location-areas instead of going through all 1000+. Or even go by game or region.
Adding support for different types of balls (Pokeballs, Great Balls, Ultra Balls, etc), which have different chances of catching pokemon
	Currently using `rand(0,baseEXP) < 25`, could look at actual catch rates from games?
Random encounters with wild pokemon
	Add wander command, track current area with AppState

/--------------------------------------------------------------------------------------------\
|	Ideas for project extension from lesson (that we're probably not going to bother with)   |
\--------------------------------------------------------------------------------------------/
Persist a user's Pokedex to disk so they can save progress between sessions
Keep pokemon in a "party" and allow them to level up
Allow for pokemon that are caught to evolve after a set amount of time
Simulate battles between pokemon
Add more unit tests (ew)
*/

func main() {
	initCommands()
	input := bufio.NewScanner(os.Stdin)
	cfg := pokeapi.AppState{
		Pokedex: make(map[string]pokeapi.Pokemon),
		Cache:   pokecache.NewCache(5 * time.Second),
	}
	for {
		fmt.Print("Pokedex > ")
		input.Scan()
		input := cleanInput(input.Text())
		if v, ok := input[1]; ok {
			cfg.Arg = v
		} else {
			cfg.Arg = ""
		}
		if cmd, ok := CommandRegistry[input[0]]; ok {
			cmd.callback(&cfg)
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) map[int]string {
	operableText := strings.ToLower(text)
	stringSlice := strings.Split(operableText, " ")

	rv := make(map[int]string)
	it := 0
	for _, val := range stringSlice {
		if val != "" {
			rv[it] = val
			it++
		}
	}
	return rv
}
