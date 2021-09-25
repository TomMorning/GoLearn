package pointer

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (wallet *Wallet) Deposit(amount Bitcoin) {
	fmt.Println("address of balance in Deposit is", &wallet.balance)
	wallet.balance += amount
}

func (wallet *Wallet) Balance() Bitcoin {
	return wallet.balance
}

var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

func (wallet *Wallet) Withdraw(amount Bitcoin) error {
	if amount > wallet.balance {
		return InsufficientFundsError
	}
	wallet.balance -= amount
	return nil
}
