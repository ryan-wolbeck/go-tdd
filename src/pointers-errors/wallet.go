package pointererrors

import "fmt"

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
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

//Balance : The method to get the ballance of the struct wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
