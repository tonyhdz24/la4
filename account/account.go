package account

import (
	. "fmt"
	. "local/customer"
)

// IAccount interface represents the basic operations for an account.
type IAccount interface{
	Balance()float32
	Deposit(amount float32)
	Withdraw(amount float32)
	String() string
	Accrue(rate float32) float32
}

// Account struct represents an account with a number, a customer, and a balance.
type Account struct{
	number string
	customer Customer
	balance float32
}

// NewAccount creates a new Account instance and initializes it with the provided values.
// @param number the account number.
// @param customer the customer associated with the account.
// @param balance the initial balance of the account.
// @return a pointer to the newly created Account.
func NewAccount(number string,customer Customer, balance float32 ) (a *Account){
	a= new(Account)
	a.Init(number,customer,balance)
	return
}

// Account Methods

// Init initializes an Account with the provided number, customer, and balance.
// @param number the account number.
// @param customer the customer associated with the account.
// @param balance the initial balance of the account
func (a *Account) Init(number string,customer Customer, balance float32){
	a.number = number
	a.customer = customer
	a.balance = balance
}


// AccountNumber returns the account number.
// @return the account number as a string.
func (a *Account) AccountNumber() string{
	return a.number 
}

// Balance returns the current balance of the account.
// @return the current balance as a float32.
func (a *Account) Balance() float32{
	return a.balance
}
// Deposit adds the specified amount to the account balance.
// @param amount the amount to be deposited.
func (a *Account) Deposit(amount float32){
	a.balance += amount
}
// Withdraw subtracts the specified amount from the account balance.
// @param amount the amount to be withdrawn.
func (a *Account) Withdraw(amount float32){
	a.balance -= amount
}
// String returns a string representation of the account.
// @return the string representation of the account.
func (a *Account) String() string{
	return a.number + " : " + a.customer.String() + " : $"+ Sprintf("%.2f",a.balance)
}
