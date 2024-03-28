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
			fmt.Printf("\033[A\033[2K Jeremi: %v\033[B", *j) //po value\033[2K
		case "k":
			*k++
			fmt.Fprintf(os.Stdout, "Kacper %v\r", *k) //input nadpisuje output, wysłanie inputa powoduje newline
		case "l":
			*l++
			fmt.Printf("\rLinda: %v\r", *l)
			// fmt.Fprintf(os.Stdout, "\r")
		default:
			fmt.Printf("\rwrong input")
		}
	}
}

var pl string
var j, k, l int

func main() {
	var i string
	fmt.Println("send one of the letters: j k l")
	plCh := make(chan string)

	go switchPlayer(plCh, &j, &k, &l)

	for i != "exit" {
		fmt.Scan(&i)
		plCh <- i
	}
}

// 			fmt.Printf("Jeremi: %v\nKacper: %v\nLinda: %v\n", *j, *k, *l)
