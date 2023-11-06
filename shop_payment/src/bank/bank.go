package bank

type Account struct {
	Number  uint64
	Name    string
	Balance int64
}

type Bank struct {
	Brand   string
	Address string
	Account struct{} //Account nested in Bank or the other way around?
}

func (b Bank) Pay() {
	/*
		1. Receive request from the Card [Sum to pay, Bank and Account details of the Shop]
		2. Find the Account in the Database
		3. Check the balance on the Account
		4. Send approval or refusal
		5. If approved transfer money to the Shop's Bank Account. [Send int through channel]
		6. Update the Balance on Account

		7. Second Bank receives money
	*/
}

func (b Bank) Receive() {
	// 	2. Find the Account in the Database, 6. Update the Balance on Account
}
