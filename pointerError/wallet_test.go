package pointererror

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10)
	got := wallet.Ballance()
	//? ballance is unexported field why is it work here?
	//? because it is the same package ?
	// got := wallet.ballance
	want := 10.0
	fmt.Printf("address of balance in test is %p \n", &wallet.ballance)

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
