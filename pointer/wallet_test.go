package pointer

import "testing"
import "fmt"

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func assertError(t *testing.T, got error, want error) {
	if got == nil {
		t.Error("wanted an error but didn't get one")
	}

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		fmt.Println("address of balance in test is", &wallet.balance)

		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(10)

		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startBalance := Bitcoin(20)
		wallet := Wallet{balance: startBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startBalance)
		assertError(t, err, InsufficientFundsError)
	})
}
