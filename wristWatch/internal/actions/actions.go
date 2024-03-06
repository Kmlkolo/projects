package actions

import (
	"wristWatch/internal/general"
)

func (W *general.Watch) Mode() {
	general.M++
	switch general.M {
	case general.ModeClock:
		general.Watch.Clock.ModeName = "Clock"
		general.Watch.Clock.TimeVal = ""
	case general.ModeStopwatch:
		general.Watch.Stopwatch.ModeName = "Stopwatch"
		general.Watch.Stopwatch.TimeVal = ""
	case general.ModeTimer:
		general.Watch.Timer.ModeName = "Timer"
		general.Watch.Timer.TimeVal = ""
	}
}

func (d *display.Display) Start() {
	time.Since() //Start counting

}

func (d *display.Display) Start() {
	time.Start()
}

type Starter interface {
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
