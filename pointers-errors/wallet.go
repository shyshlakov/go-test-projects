package errors

import (
	"errors"
	"fmt"
)

// Stringer ...
type Stringer interface {
	String() string
}

// Bitcoin ...
type Bitcoin int

// Wallet ...
type Wallet struct {
	balance Bitcoin
}

// ErrInsufficientFunds ...
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Deposit ...
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance ..
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Withdraw ...
func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
