package pointererrors

import "fmt"

//Wallet : struct for the wallet
type Wallet struct {
	balance int
}

//Deposit : method for making a depost
func (w *Wallet) Deposit(amount int) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

//Balance : The method to get the ballance of the struct wallet
func (w *Wallet) Balance() int {
	return w.balance
}
