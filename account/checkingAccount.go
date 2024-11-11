package account

import (
	// . "fmt"
	. "local/customer"
)

// Checking Account Interface: Defines the core methods for a checking account struct. Contains all methods in Account
type ICheckingAccount interface{
	IAccount
}
// Checking Account Struct: Represents a checking account struct that contains all variables defined in account 
type CheckingAccount struct {
	Account
}


// Checking Account Class Methods

// NewCheckingAccount initializes and returns a new CheckingAccount instance account number,customer and balance set to parameter values
func NewCheckingAccount(number string,customer Customer, balance float32) (ca *CheckingAccount){
	ca= new(CheckingAccount)
	ca.Init(number,customer,balance)
	return
}

// Init Function
func (ca *CheckingAccount) Init(number string,customer Customer, balance float32){
	ca.number = number
	ca.customer = customer
	ca.balance = balance
}
//? In the CheckingAccount.java Accrue is defined but doesnot do anything? Accure is not defined in Account.java
func (ca *CheckingAccount) Accrue(rate float32) float32{
	return 0
}