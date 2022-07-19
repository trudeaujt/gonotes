package pointers

import (
	"errors"
	"fmt"
)

type GoCoin int

//This is an implementation of the 'Stringer' interface.
//This defines how your custom type is printed when used with the %s format string in prints.
//As we can see, the syntax for creating a method on a type declaration is the same as it is on a struct.
func (g GoCoin) String() string {
	return fmt.Sprintf("%d GC", g)
}

// Wallet A secure wallet.
//We don't want to expose our inner state to the rest of the world, so we will control access via methods.
type Wallet struct {
	//since this variable starts with a lowercase b, it is private outside of the package it is defined in.
	balance GoCoin
}

// Deposit here, w Wallet is a copy of whatever we called the method from -
//that means that we are just accessing methods on the copy, and the real balance will not be updated.
//To fix this, we need to accept a pointer in the form of w *Wallet.
func (w *Wallet) Deposit(amount GoCoin) {
	//fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

var ErrInsufficientFunds = errors.New("cannot withdraw more than you have")

func (w *Wallet) Withdraw(amount GoCoin) error {
	if amount >= w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

// Balance Technically, we do not need a pointer to the actual value here - since we are just reading the value a copy is fine.
//But by convention we keep the method receiver types the same for consistency.
func (w *Wallet) Balance() GoCoin {
	//We could write: return (*w).balance, but that is not needed in Go.
	//Here, struct pointers are automatically dereferenced.
	//Explicit dereferencing is not needed.
	return w.balance
}
