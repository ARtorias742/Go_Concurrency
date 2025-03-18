package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	balance int
	mutex   sync.RWMutex // Read-write mutex
}

func (a *Account) Deposit(amount int) {
	a.mutex.Lock() // Exclusive lock for writing
	defer a.mutex.Unlock()
	a.balance += amount
	fmt.Printf("Deposited %d, new balance: %d\n", amount, a.balance)
}

func (a *Account) Withdraw(amount int) {
	a.mutex.Lock() // Exclusive lock for writing
	defer a.mutex.Unlock()
	if a.balance >= amount {
		a.balance -= amount
		fmt.Printf("Withdrew %d, new balance: %d\n", amount, a.balance)
	} else {
		fmt.Println("Insufficient funds")
	}
}

func (a *Account) GetBalance() int {
	a.mutex.RLock() // Shared lock for reading
	defer a.mutex.RUnlock()
	return a.balance
}

func main() {
	account := &Account{balance: 100}
	var wg sync.WaitGroup

	// Writers: Deposit and Withdraw
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			account.Deposit(50)
			time.Sleep(time.Millisecond * 100)
			account.Withdraw(30)
		}()
	}

	// Readers: GetBalance
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			balance := account.GetBalance()
			fmt.Println("Current balance:", balance)
		}()
	}

	wg.Wait()
	fmt.Println("Final balance:", account.GetBalance())
}
