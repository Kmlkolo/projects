package main

import (
	"fmt"
	"os"
)

func switchPlayer(usrInput chan string, j, k, l *int) {
	pl := <-usrInput
	switch pl {
	case "j":
		*j++
		fmt.Printf("Jeremi: %v", j) //po value\033[2K
	case "k":
		*k++
		fmt.Fprintf(os.Stdout, "Kacper %v\r", k) //input nadpisuje output, wysłanie inputa powoduje newline
	case "l":
		*l++
		fmt.Printf("\rLinda: %v\r", l)
		// fmt.Fprintf(os.Stdout, "\r")
	default:
		fmt.Printf("\rwrong input")
	}
}

var pl string
var j, k, l int

func main() {
	var i string
	fmt.Println("send one of the letters: j k l")
	plCh := make(chan string)
	for i != "exit" {
		fmt.Scan(&i)
		plCh <- i
		go switchPlayer(plCh, &j, &k, &l)
	}
}
