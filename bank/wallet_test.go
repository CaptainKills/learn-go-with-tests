package bank

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
		checkBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(Bitcoin(10))
		checkBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw Sufficient Funds", func(t *testing.T) {
		startingBalance := 20
		wallet := Wallet{balance: Bitcoin(startingBalance)}

		err := wallet.Withdraw(Bitcoin(10))
		checkBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("Withdraw Insufficient Funds", func(t *testing.T) {
		startingBalance := 20
		wallet := Wallet{balance: Bitcoin(startingBalance)}

		err := wallet.Withdraw(Bitcoin(30))
		checkBalance(t, wallet, Bitcoin(startingBalance))
		assertError(t, err, ErrInsufficientFunds)
	})
}

func checkBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Fatal("did not want an error but go one")
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
