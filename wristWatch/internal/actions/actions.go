package actions

import (
	"wristWatch/internal/general"
)

func (W *Watch) Mode() {
	//gdzie ustawić wartości początkowe/fabryczne Mode:=0 - nie na początku funkcji przecież
	general.M++
	switch general.M {
	case general.ModeClock:
		general.Watch.Clock.ModeName = "Clock"
		general.Watch.Clock.TimeVal = ""
	case general.ModeStopwatch:
		general.Watch.Stopwatch.ModeName = "Stopwatch"
		general.Watch.Stopwatch.TimeVal = "" //zachowana wartość z fmt.Sprint lub global var
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
