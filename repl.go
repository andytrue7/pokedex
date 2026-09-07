package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(config *ReplStateConfig) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		words := cleanInput(text)
		if len(words) == 0 {
			continue
		}
		command := words[0]
		if cmd, ok := config.commands[command]; ok {
			err := cmd.callback(config)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command:", command)
		}
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(text)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}
	return words
}