package p

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()

		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

	}

	assetError := func(t testing.TB, got, want error) {
		if got == nil {
			t.Fatal("wanted an error but didnt get one")
		}

		if got != want {
			t.Errorf("got %q wanted %q", got, want)
		}
	}

	assertNoError := func(t testing.TB, got error) {
		t.Helper()

		if got != nil {
			t.Fatal("got an error, but didnt want one")
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10)

		assertBalance(t, wallet, Bitcoin(10))

	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}

		got := wallet.Withdraw(Bitcoin(2))

		assertNoError(t, got)

		assertBalance(t, wallet, Bitcoin(8))

	})

	t.Run("Withdraw Insufficent", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assetError(t, err, ErrInsufficentFunds)

	})

}
