package main

import (
	"fmt"
	_ "paycard"
	_ "bank"
)

type ATM struct {

}

type Shop struct {
	Name string
	Address string
	// Czy zapisać jako atrybut: konto bankowe?
}

type Client struct {
	Name string
	// Czy zapisać jako atrybuty: kupowany produkt, bank?
}

type product struct {
	name string
	Price uint32
}

func (s Shop) Count() {
	/*
	Loop: Automatically opens new order/basket
	1. Scan a product
	2. Add its price to the sum
	3. Close the order (input, keyboard interrupt) and sum up
	
	XX4.? Send request to the Bank, to open slot for receiving ? XX
	4. Display the sum on a terminal and inform about Bank Account
	*/
}

main() {
	product := 
	fmt.Println("Hello Banking World")
}