package actions

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
	
