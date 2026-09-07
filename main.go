package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
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
		if cmd, ok := commandRegister()[command]; ok {
			cmd.callback()
		} else {
			fmt.Println("Unknown command:", command)
		}
	}
}
