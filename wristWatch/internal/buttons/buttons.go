package buttons

import (
	"fmt"
	"strconv"
	"wristWatch/internal/general"
)

type Button uint8

var b Button

// func init() { b = 1 }

const (
	Button1Start Button = iota
	Button1Stop  Button = 1
	Button1Reset Button = 2
	Button2Mode  Button = 3
	Button2Set   Button = 4
	Button3Plus  Button = 5
	Button4Minus Button = 6
)

func (W general.Watch) Snap(b Button) {
	var b1 string
	fmt.Scanln(&b1)
	b2, _ := strconv.ParseUint(b1, 10, 8)
	b = Button(b2)
	switch b {
	case Button1Start:
		fmt.Println("Start")
		Start()
	case Button1Stop:
		fmt.Println("Stop")
		Stop()
	case Button1Reset:
		fmt.Println("Reset")
		Reset()
	case Button2Mode:
		//przeskakuje pomiędzy 0 1 2
		fmt.Println("Mode")
		Mode()
	case Button2Set:
		fmt.Println("Set")
		Set()
	case Button3Plus:
		fmt.Println("Plus")
		Plus()
	case Button4Minus:
		fmt.Println("Minus")
		Minus()
	}
}
