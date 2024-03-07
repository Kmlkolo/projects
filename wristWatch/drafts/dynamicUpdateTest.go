package main

import (
	"fmt"
	"os"
)

func main() {
	var i string
	var j, k, l int
	fmt.Println("send one of the letters: j k l")
	fmt.Scanln(&i)

	for {
		switch i {
		case "j":
			j++
			fmt.Printf("Jeremi: %v", j) //po value\033[2K
			fmt.Scanln(&i)
		case "k":
			k++
			fmt.Fprintf(os.Stdout, "Kacper %v\r", k) //input nadpisuje output, wysłanie inputa powoduje newline
			fmt.Scanln(&i)
		case "l":
			l++
			fmt.Printf("\rLinda: %v\r", l)
			fmt.Scanln(&i)
			// fmt.Fprintf(os.Stdout, "\r")
		default:
			fmt.Printf("\rwrong input")
		}
	}
}
