package pointererror

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		got := wallet.Ballance()
		//? ballance is unexported field why is it work here?
		//? because it is the same package ?
		// got := wallet.ballance
		want := BitCoin(10)
		fmt.Printf("address of balance in test is %p \n", &wallet.ballance)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{BitCoin(20)}
		wallet.Withdraw(BitCoin(10))
		got := wallet.Ballance()
		want := BitCoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
