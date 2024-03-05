package main

import (
	"fmt"
	"os"
	"time"
)

// GOAL: dynamic change of value in CLI. not skipping to newline triggered by sending the input.

func main() {
	var mode int

	fmt.Println("Send \"4\" to switch mode.")

	for {
		var input string

		// Print the prompt
		// fmt.Print("\r")

		// Read user input
		fmt.Scan(&input)

		// Clear the current line
		fmt.Print("\033[2K\r") //kod asci do czyszczenia linii

		// Process the input
		switch input {
		case "4":
			mode++
			fmt.Print("\r", mode)
			time.Sleep(1 * time.Second)
			fmt.Print("\033[2K\r")
		default:
			fmt.Print("\rCzwórkę wpisz.")
		}

		// Check for exit condition
		if input == "exit" {
			os.Exit(0)
		}
	}
}
