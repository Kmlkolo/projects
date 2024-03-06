package main

import (
	"fmt"
	"time"
)

// var a [3]int

type Mode int

var i string
var a Mode

func (aa *Mode) Mode() {
	*aa++
	fmt.Printf("\r%v", *aa)
	time.Sleep(1 * time.Second)
	*aa++
	fmt.Printf("\r%v", *aa)
	// fmt.Printf("\r%v", *aa)

}

func main() {
	// a = [3]int{0, 1, 2}
	a = 0 //default value
	fmt.Println("Send \"4\" to switch mode.")
	for i != "0" {
		fmt.Scan(&i)
		switch i {
		case "4":
			a.Mode()

		}
	}
	// switch i {
	// case "a":
	// 	a = 0
	// case "b":
	// 	a = 1
	// case "c":
	// 	a = 2
	// }
}

// GOAL: toggling between 3 modes. 1) make dynamic update effect 2) close range: 0-1-2->0...
