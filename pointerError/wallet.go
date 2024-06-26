package pointererror

import (
	"errors"
	"fmt"
)

type BitCoin float64

type Stringer interface {
	String() string
}

func (b BitCoin) String() string {
	return fmt.Sprintf("%.2f BTC", b)
}

type Wallet struct {
	ballance BitCoin
}

func (w *Wallet) Deposit(amount BitCoin) {
	fmt.Printf("address of balance in Deposit is %p \n", &w.ballance)
	w.ballance += amount
}

func (w Wallet) Ballance() BitCoin {
	return w.ballance
}

func (w *Wallet) Withdraw(amount BitCoin) error {
	if w.ballance < amount {
		return errors.New("cannot withdraw, insufficient funds")
	}
	w.ballance -= amount
	return nil
}
