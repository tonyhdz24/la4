# CS 354: Language Assignment #4 - GO

## Antonio Hernandez (114225246)

### How to Run Test Suite
1. To run test suite open up your terminal and cd into the directory containing the folders `account`, `bank`,`customer` and `Main`.
    `cd path/to/your/directory`

2. Once in the desired folder run the DOIT script to run test suite (You may need to update permissions on file)
    `./DOIT`

3. Output of test results will appear and no futher input is required.

### Project Overview
The purpose of Language Assignment 4 was to become familiar with and gain a simple understanding of the GO programming language. To achieve this we were tasked with porting simple account/bank Java code into GO. This allowed us to see how classes written in Java can be represented in GO. We also use Goroutine/Channel in this assignment to calculate accrued interest.

### Code Overvew
The code for Language Assignment is fairly straightforward. It is a port of the provided Java code (insert link to code ). In the code we model a savings account and a checking account that both extend a more general account class. The accounts all have an account number, a customer, and a balance. On an account class you are able to get the account number, get the current balance, deposit/withdraw money and finally get a string representation of the account. The checking and savings account classes extend account by adding in an accrue method that accrues interest on the balance of an account.


### Sources used

- [Notes provided in class](https://github.com/BoiseState/CS354-resources/tree/master/buff/classes/354/pub)
- [GO Starter Code](https://github.com/BoiseState/CS354-resources/tree/master/buff/classes/354/pub/la4)