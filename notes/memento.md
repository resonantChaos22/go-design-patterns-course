- `Memento` is a token representing the system state. It lets us roll back to the state when the token was generated. May or may not expose state information.
- This can be used as an alternative to `Command` pattern
- Changing a `Memento` is non-idiomatic

## Simple Memento Pattern

```go
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

ba, m0 := NewBankAccount(100)
m1 := ba.Deposit(25)  //  ba = 125
m2 := ba.Deposit(50)  //  ba = 175
ba.Restore(m1)  //  ba = 125
ba.Restore(m0)  //  ba = 100
```

- We basically have a `Memento` of every time the value of the balance was changes so that we have a state that we can rollback to.

## Undo and Redo

- In the same situation, we can implement an undo and redo functionality.
- This is done via having an array of mementos which contains a chain of all the changes that the bank balance has gone through.
- During `Undo` and `Redo`, we modify the balance based on the current state.

```go
type BankAccount struct {
 balance int
 changes []*Memento
 current int
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
```

- We can only `Redo` an action that has been `Undo`'d. We can not actually "Re-do" an action.
- If you remember, in the `Commands` section, we also implemented undo but that was basically done as a symmetric of the command that was given before.
- `Command` is better suited to handle `Redo` as we only store the changes in state where as for `Memento` we have to store the whole state.

## Memento Vs Flyweight

- Both patterns provide a token that clients can hold onto.
- `Memento` uses the token to be fed back to the system. There are no methods and it's not mutable.
- `Flyweight` actually does business logic on the said tokens
