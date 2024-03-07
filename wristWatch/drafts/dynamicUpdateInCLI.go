// GOAL: dynamic change of value in CLI. not skipping to newline triggered by sending the input.
package main

import (
	_ "bufio"
	"fmt"
	"os"
	_ "os"
	"time"
)

var z string = ""

func Spinner(delay time.Duration) {
	for z != "0" {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
			fmt.Fprintf(os.Stdout, "\r \r")
		}
	}
}

func main() {
	d, _ := time.ParseDuration("0.5s")
	Spinner(d)
}

// func main() {
// 	var mode int

// 	scanner := bufio.NewScanner(os.Stdin)

// 	fmt.Println("Send \"4\" to switch mode.")
// 	for {
// 		var input string
// 		scanner.Scan()
// 		input = scanner.Text()

// 		// fmt.Print("\033[2K\r") //kod asci do czyszczenia linii
// 		switch input {
// 		case "4":
// 			mode++
// 			fmt.Print("\r", mode)
// 			// time.Sleep(1 * time.Second)
// 			// fmt.Print("\033[2K\r")
// 		default:
// 			fmt.Print("\rCzwórkę wpisz.")
// 		}
// 		if input == "exit" {
// 			os.Exit(0)
// 		}
// 	}
// }
