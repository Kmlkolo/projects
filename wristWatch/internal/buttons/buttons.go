package buttons

import "fmt"

type Button uint8

var b Button

const (
	Button1Start Button = iota
	Button1Stop  Button
	Button1Reset Button
	Button2Mode  Button
	Button2Set   Button
	Button3Plus  Button
	Button4Minus Button
)

func (w Watch) Snap(b Button) {
	b := fmt.Scanln(&b) //? b bez wcześniejszej deklarcji
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
		fmt.Println("2Mode")
		Mode()
	case Button2Set:
		fmt.Println("2Set")
		Set()
	case Button3Plus:
		fmt.Println("3Plus")
		Plus()
	case Button4Minus:
		fmt.Println("Minus")
		Minus()
	}
}
