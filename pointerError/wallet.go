package pointererror

import "fmt"

type Wallet struct {
	ballance float64
}

func (w *Wallet) Deposit(amount float64) {
	fmt.Printf("address of balance in Deposit is %p \n", &w.ballance)
	w.ballance += amount
}

func (w Wallet) Ballance() float64 {
	return w.ballance
}
