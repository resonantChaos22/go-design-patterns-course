package bankaccount

import "fmt"

type Memento struct {
	Balance int
}
type BankAccount struct {
	balance int
	changes []*Memento
	current int
}

func (b *BankAccount) String() string {
	return fmt.Sprint("Balance = $", b.balance, ", current = ", b.current)
}
func (b *BankAccount) Deposit(ammount int) *Memento {
	b.balance += ammount
	m := Memento{
		Balance: b.balance,
	}
	b.changes = append(b.changes, &m)
	b.current++
	fmt.Println("Deposited ", ammount, "\b, balance is now", b.balance)
	return &m
}
func (b *BankAccount) Restore(m *Memento) {
	if m == nil {
		return
	}

	b.balance = m.Balance
	b.changes = append(b.changes, m)
	b.current = len(b.changes) - 1
}
func (b *BankAccount) Undo() *Memento {
	if b.current == 0 {
		return nil
	}

	b.current--
	m := b.changes[b.current]
	b.balance = m.Balance
	return m
}
func (b *BankAccount) Redo() *Memento {
	if len(b.changes) <= b.current+1 {
		return nil
	}

	b.current++
	m := b.changes[b.current]
	b.balance = m.Balance
	return m
}

func NewBankAccount(balance int) *BankAccount {
	b := &BankAccount{
		balance: balance,
	}
	b.changes = append(b.changes, &Memento{
		Balance: balance,
	})

	return b
}

func TestBAUndoRedo() {
	ba := NewBankAccount(100)
	ba.Deposit(50)
	ba.Deposit(25)
	fmt.Println(ba)

	ba.Undo()
	fmt.Println("Undo 1", ba)

	ba.Undo()
	fmt.Println("Undo 2", ba)

	ba.Redo()
	fmt.Println("Redo", ba)
}
