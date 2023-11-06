package Card

type Card struct {
	Number uint64
	Name   string
	Brand  string
}

func (c Card) Pay() {
	/*
		1. Receive sum from the Shop and Bank Account details
		2. Send request to the Bank
		3. Receive approval or refusal
	*/
}
