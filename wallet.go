package wallet

import (
	"errors"
	"fmt"
	"sync"
)

type Bitcoin float64

func (b Bitcoin) String() string {
	return fmt.Sprintf("%.2f BTC", b)
}

type Wallet struct {
	balance Bitcoin
	mu      sync.Mutex
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	w.mu.Lock()
	defer w.mu.Unlock()
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")
