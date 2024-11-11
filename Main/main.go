package main

import (
	. "fmt"
	. "local/account"
	. "local/bank"
	. "local/customer"
)

func main() {
	{
		// Customer
		ed:=NewCustomer("Ed")
		hank:=NewCustomer("Hank")
		william:=NewCustomer("William")
		ann:=NewCustomer("Ann")

		// TEST customer
		if ed.String() == "Ed" {
			Printf("Customer Test 1 Passed! 🟢\n")
			
		}else{
			Printf("Customer Test 1 Failed! 🔴\n")
		}

		if hank.String() == "Hank" {
			Printf("Customer Test 2 Passed! 🟢\n")
			
		}else{
			Printf("Customer Test 2 Failed! 🔴\n")
		}

		if william.String() == "William" {
			Printf("Customer Test 3 Passed! 🟢\n")
			
		}else{
			Printf("Customer Test 3 Failed! 🔴\n")
		}

		if ann.String() == "Ann" {
			Printf("Customer Test 4 Passed! 🟢\n")
			
		}else{
			Printf("Customer Test 4 Failed! 🔴\n")
		}
		

		// Test Accounts
		edAccount:=NewAccount("1",*ed,21)
		hankCheckingAccount:= NewCheckingAccount("2",*hank,12)
		williamSavingAccount:= NewSavingAccount("3",*william,10)

		if edAccount.String() == "1 : Ed : $21.00" {
			Printf("Account Test 1 Passed! 🟢\n")
			
		}else{
			Printf("Account Test 1 Failed! 🔴\n")
			
		}
		if hankCheckingAccount.String() == "2 : Hank : $12.00" {
			Printf("Account Test 2 Passed! 🟢\n")
			
		}else{
			Printf("Account Test 2 Failed! 🔴\n")

		}
		if williamSavingAccount.String() == "3 : William : $10.00" {
			Printf("Account Test 3 Passed! 🟢\n")
			
		}else{
			Printf("Account Test 3 Failed! 🔴\n")

		}

		// Testing bank
		// Creating new bank
		bank:= NewBank()
		Printf("\nTesting Bank...\n")

		// Adding checking and savings accounts
		bank.Add(NewCheckingAccount("01001",*ann,100.00))
		bank.Add(NewSavingAccount("01002",*ann,200.00))
		Printf("Accounts in Bank:\n%s\nAppling 0.02 intrest rate to accounts...\n",bank.String())

		// Appling intrest rate
		bank.Accrue(0.02)
		// Print the bank's information
		Printf("\nBank accounts:\n%s", bank.String())

	}
}