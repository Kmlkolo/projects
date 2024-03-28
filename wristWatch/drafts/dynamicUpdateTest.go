package main

import (
	"fmt"
	"os"
)

func switchPlayer(plCh <-chan string, j, k, l *int) {
	for {
		pl := <-plCh
		switch pl {
		case "j":
			*j++
			fmt.Printf("\033[A\033[2KJeremi: %v", *j)
		case "k":
			*k++
			fmt.Fprintf(os.Stdout, "\033[A\033[2KKacper: %v", *k)
		case "l":
			*l++
			fmt.Printf("\033[A\033[2KLinda: %v", *l)
			// fmt.Fprintf(os.Stdout, "\r")
		case "points":
			fmt.Printf("\033[A\033[2K----------\nJeremi: %v\nKacper: %v\nLinda: %v\n----------\n", *j, *k, *l)

		default:
			fmt.Printf("\rwrong input")
		}
	}
}

var pl string
var j, k, l int

func main() {
	var i string
	fmt.Println("To give a point to Jeremi, Kacper or Linda, send: j / k / l" +
		"\nTo see the current state of points, send: points" +
		"\nTo quit the program, send: exit")
	plCh := make(chan string)

	go switchPlayer(plCh, &j, &k, &l)

	for i != "exit" {
		fmt.Scan(&i)
		plCh <- i
	}
}
