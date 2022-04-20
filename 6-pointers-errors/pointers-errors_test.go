package main

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		balanceHelper(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{10}
		err := wallet.Withdraw(Bitcoin(10))
		assertNotError(t, err)
		balanceHelper(t, wallet, Bitcoin(0))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertError(t, err, ErrInsufficientFunds)
		balanceHelper(t, wallet, startingBalance)
	})
}

func balanceHelper(t testing.TB, w Wallet, want Bitcoin) {
	t.Helper()
	got := w.Balance()
	if got != want {
		t.Errorf("got: %s, want %s", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didnt get an error")
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNotError(t testing.TB, got error) {
	if got != nil {
		t.Errorf("got %q, want nil", got)
	}
}
