package bank

import (
	. "fmt"
	"local/account"
	. "local/account"
	. "time"
)

// Sleep function
func sleep(c chan float32, acc IAccount, rate float32) {
	Sleep(1 * Second) // Simulating a delay
	interest := acc.Accrue(rate)
	c <- interest
}

// Bank Interface: Defines the core methods for a bank structure
type IBank interface{
	// Add: Adds an account to the bank
	Add()
	// Accure: Applies intreset to acounts
	Accrue()
	// String: Returns a string reprsentation of the bank and its account
	String() string
}

// Bank Struct: Bank represents a bank that manages a collection of accounts
type Bank struct{
	// accounts is a map storing accounts in the bank, keyed by account number.
	accounts map[*account.IAccount]account.IAccount
}


// Bank Constructor: NewBank initializes and returns a new Bank instance with an empty account map
func NewBank()(b *Bank){
	b = new(Bank)
	b.accounts = make(map[*IAccount]IAccount)
	return
}

func (b *Bank) Add(account IAccount){
	b.accounts[&account] = account
}


func (b *Bank) Accrue(rate float32){

	totalInterest := float32(0)
	// Using Goroutine/Channel to calculate accrue
	c:=make(chan float32)
	for _, acc := range b.accounts {
		go sleep(c,acc,rate)
	}
	// Collect the interest results from the channel
	for range b.accounts {
		totalInterest += <-c
	}
	
	// Print the total interest accrued
	Printf("Total interest accrued: %.2f\n", totalInterest)

}

func (b *Bank) String() string {
	result := ""
	for _, acc := range b.accounts {
		result += acc.String() + "\n"
	}
	return result
}