package main

import (
	"fmt"
	_ "wristWatch/internal/actions"
	_ "wristWatch/internal/buttons"
	"wristWatch/internal/display"
)

type Mode uint8

var m Mode

const (
	ModeClock     Mode = iota
	ModeStopwatch Mode = 1
	ModeTimer     Mode = 2
)

type Watch struct {
	Clock     display.Display
	Stopwatch display.Display
	Timer     display.Display
}

// var display Display = Display{
// func (t Time) Clock() (hour, min, sec int)

func main() {
	fmt.Println()
}
