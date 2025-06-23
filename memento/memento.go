package main

import "fmt"

type Memento struct {
	Balance int
}

type BankAccount struct {
	balance int
}

func NewBankAccount(balance int) (*BankAccount, *Memento) {
	return &BankAccount{
			balance: balance,
		}, &Memento{
			Balance: balance,
		}
}

func (b *BankAccount) Deposit(ammount int) *Memento {
	b.balance += ammount
	return &Memento{
		Balance: b.balance,
	}
}
func (b *BankAccount) Restore(m *Memento) {
	b.balance = m.Balance
}

func TestMemento() {
	ba, m0 := NewBankAccount(100)
	m1 := ba.Deposit(50)
	m2 := ba.Deposit(25)
	fmt.Println(ba.balance)
	ba.Restore(m1)
	fmt.Println(ba.balance)
	ba.Restore(m2)
	fmt.Println(ba.balance)
	ba.Restore(m0)
	fmt.Println(ba.balance)
}
