package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

type Account struct {
	userName string
	balance  uint64
	mut      sync.RWMutex
}

func (account *Account) Deposit(amount uint64) {
	wg.Done()
	defer account.mut.Unlock()
	account.mut.Lock()
	account.balance += amount
}

func (account *Account) Withdraw(amount uint64) bool {
	wg.Done()
	defer account.mut.Unlock()
	account.mut.Lock()
	if amount > 0 && amount <= account.balance {
		account.balance -= amount
		return true
	}
	return false
}

func (account *Account) Balance() uint64 {
	defer account.mut.RUnlock()
	account.mut.RLock() //creating a read locker so that writers cant write while reading
	return account.balance
}

func main() {
	acc := Account{userName: "user 1"}
	fmt.Println(acc.Balance())
	for range 100 {
		wg.Add(1)
		go acc.Deposit(1)
	}

	for range 100 {
		wg.Add(1)
		go acc.Withdraw(1)
	}
	wg.Wait()
	fmt.Println(acc.Balance())

}
