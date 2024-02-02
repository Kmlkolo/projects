package general

type Mode uint8

var M Mode

const (
	ModeClock     Mode = iota
	ModeStopwatch Mode = 1
	ModeTimer     Mode = 2
)

type Display struct {
	ModeName string
	TimeVal time.Time //Time?
}	

type Watch struct {
	Clock     Display
	Stopwatch Display
	Timer     Display
}



// 	Mode: [3]string{"Clock", "Stopwatch", "Timer"},	
// 	Time: time.Time{},
// }

var defaultWatch Watch

func init() {
	defaultWatch.Clock.ModeName = "Clock"
	defaultWatch.Clock.TimeVal = "00:00:00"
	defaultWatch.Stopwatch.ModeName = "Stopwatch"
	defaultWatch.Stopwatch.TimeVal = "00:00:00"
	defaultWatch.Timer.ModeName = "Timer"
	defaultWatch.Timer.TimeVal = "00:00:00"
}
const mn0 Watch.Clock.ModeName = "Clock" //expected ';', found '.'syntax
var tv0 Watch.Clock.TimeVal = "00:00:00" //syntax error: unexpected . after top level declarationgo-staticcheck

const Watch.Stopwatch.ModeName = "Stopwatch" //syntax error: unexpected . after top level declarationgo-staticcheck
var Watch.Stopwatch.TimeVal = "00:00:00" //syntax error: unexpected ., expected typego-staticcheck

const mn2 Watch.Timer.Display.ModeName = "Timer" //syntax error: unexpected . after top level declarationgo-staticcheck
var tv2 Watch.Timer.Display.TimeVal = "00:00:00" //syntax error: unexpected . after top level declarationgo-staticcheck

