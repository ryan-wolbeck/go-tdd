package pointererrors

import (
	"errors"
	"fmt"
)

//Bitcoin : Create a type for bitcoin to make the code readable
type Bitcoin int

//Wallet : struct for the wallet
type Wallet struct {
	balance Bitcoin
}

//Stringer : Create an interface to add the bitcoin BTC to prints of the type bitcoin
type Stringer interface {
	String() string
}

//String : Changes the output from 10 to 10 BTC when printed with %s
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

//Deposit : method for making a depost
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

//ErrInsufficientFunds : variable for the insufficient funds message
var ErrInsufficientFunds = errors.New("cannnot withdraw, insufficient funds")

//Withdraw : method for making a depost
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

//Balance : The method to get the ballance of the struct wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
