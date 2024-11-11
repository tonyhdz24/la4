package account

import (
	// . "fmt"
	. "local/customer"
)

// Savings Account Interface: Defines the core methods for a savings account strucute. Contains all methods in Account
type ISavingAccount interface{
	IAccount
}
// Savings Account Struct: Represents a savings account struct that contains all variables defined in account plus an intrest variable
type SavingAccount struct {
	Account
	intrest float32
}


// Savings Account Methods

// NewSavingAccount initializes and returns a new SavingsAccount instance account number,customer and balance set to parameter values
func NewSavingAccount(number string,customer Customer, balance float32) (ca *SavingAccount){
	ca= new(SavingAccount)
	ca.Init(number,customer,balance)
	return
}

// Init Function
func (ca *SavingAccount) Init(number string,customer Customer, balance float32){
	ca.number = number
	ca.customer = customer
	ca.balance = balance
	ca.intrest = 0
}
// Accrue: Apples intreste rate to account balance
func (ca *SavingAccount) Accrue(rate float32) float32{
	ca.intrest += ca.balance * rate
	ca.balance += ca.balance * rate

	return ca.intrest

}