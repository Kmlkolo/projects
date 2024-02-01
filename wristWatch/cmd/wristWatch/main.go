//main.go
package main

import (
	"fmt"
	"time"
)

type Mode uint8
var m Mode
const (
	ModeClock     Mode = iota
	ModeStopwatch Mode = 1
	ModeTimer     Mode = 2
)

type Watch struct {
	Clock     Display
	Stopwatch Display
	Timer     Display
}	
//display.go
type Display struct {
	ModeName string
	TimeVal time.Time //Time?
}	

// 	Mode: [3]string{"Clock", "Stopwatch", "Timer"},	
// 	Time: time.Time{},
// }

const Watch.Clock.Display.ModeName = "Clock"
var tv0 Watch.Clock.Display.TimeVal = "00:00:00"
const Watch.Stopwatch.Display.ModeName = "Stopwatch"
var tv1 Watch.Stopwatch.Display.TimeVal = "00:00:00"
const Watch.Timer.Display.ModeName = "Timer"
var tv2 Watch.Timer.Display.TimeVal = "00:00:00"

//actions.go
func (w Watch) Mode() {
	//gdzie ustawić wartości początkowe/fabryczne Mode:=0 - nie na początku funkcji przecież
	Mode++
	switch Mode {
		case 0:
			Watch.Clock.Display.ModeName = "Clock"
			Watch.Clock.Display.TimeVal = ""
		case 1:
			Watch.Stopwatch.Display.ModeName = "Stopwatch"
			Watch.Stopwatch.Display.TimeVal = "" //zachowana wartość z fmt.Sprint lub global var
		case 2:
			Watch.Timer.Display.ModeName = "Timer"
			Watch.Timer.Display.TimeVal = ""
		}
	}
	
	func (d *Watch.Stopwatch.Display) Start() {
		time.Since() //Start counting
		
	}
	
	func (d *Display) Start() {
		time.Start()
	}
	
	type Starter interface { //Start robi zupełnie co innego zależnie od trybu, co jest wspólne?
		Start()
	}
	
	type Stopwatch interface {
		Starter()
		Stopper()
		Resetter()
	}
	
	type Timer interface {
		Starter()
		Stopper()
		Resetter()
		Setter()
	}
	
// var display Display = Display{
// func (t Time) Clock() (hour, min, sec int)

//buttons.go
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

func main() {
	fmt.Println()
}
