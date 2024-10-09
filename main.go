package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Pokedex > ")
	scanner := bufio.NewScanner(os.Stdin)
	commands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}

	for scanner.Scan() {
		text := scanner.Text()

		if cmd, ok := commands[text]; ok {
			fmt.Println("")
			fmt.Println("Welcome to the Pokedex!")
			fmt.Println("Usage:")
			fmt.Println("")

			cmd.callback()
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandHelp() error {
	fmt.Printf("%s: %s\n", c.name, c.description)
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}
