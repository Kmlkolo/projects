package general

import "time"

type Mode uint8

var M Mode

const (
	ModeClock     Mode = iota
	ModeStopwatch Mode = 1
	ModeTimer     Mode = 2
)

type Display struct {
	ModeName string
	TimeVal  time.Time //Time?
}

type Watch struct {
	Clock     Display
	Stopwatch Display
	Timer     Display
}

var defaultWatch Watch

func init() {
	defaultWatch.Clock.ModeName = "Clock"
	defaultWatch.Clock.TimeVal = time.Date(0, time.January, 0, 0, 0, 0, 0, time.UTC)
	defaultWatch.Stopwatch.ModeName = "Stopwatch"
	defaultWatch.Stopwatch.TimeVal = time.Date(0, time.January, 0, 0, 0, 0, 0, time.UTC)
	defaultWatch.Timer.ModeName = "Timer"
	defaultWatch.Timer.TimeVal = time.Date(0, time.January, 0, 0, 0, 0, 0, time.UTC)
}
