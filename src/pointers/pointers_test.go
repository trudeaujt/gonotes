package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		//fmt.Printf("address of balance in test is %v \n", &wallet.balance)
		assertBalance(t, wallet, GoCoin(10))
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: GoCoin(10)}
		err := wallet.Withdraw(5)
		assertNoError(t, err)
		assertBalance(t, wallet, GoCoin(5))
	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := GoCoin(1)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(2)
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
		//Nil is synonymous with null in other languages.
		//Errors can be nil because the return type of Withdraw will be error, which is an interface.
		//If you see a function that takes arguments or returns values that are interfaces, they can be nillable.
		//if err == nil {
		//	t.Error("wanted an error but didn't get one")
		//}
	})
}

func assertBalance(t testing.TB, wallet Wallet, want GoCoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}
func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		//We're using t.fatal here, which will stop the test if this is encountered.
		//This is because if the test were to continue it would carry on to the next test,
		//which would panic because of a nil pointer.
		t.Fatal("wanted an error but didn't get one")
	}
	if got.Error() != want.Error() {
		t.Errorf("got %q, want %q", got, want)
	}
}
