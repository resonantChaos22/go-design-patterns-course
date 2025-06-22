package command

import "fmt"

var overdraftLimit = -500

type BankAccount struct {
	balance int
}

func (b *BankAccount) Deposit(amount int) {
	b.balance += amount
	fmt.Println("Deposited", amount, "\b, balance is now", b.balance)
}
func (b *BankAccount) Withdraw(amount int) bool {
	if b.balance-amount >= overdraftLimit {
		b.balance -= amount
		fmt.Println("Withdrew", amount, "\b, balance is now", b.balance)
		return true
	}

	return false
}

type Command interface {
	Call()
	Undo()
	Succeeded() bool
	SetSucceeded(value bool)
}
type Action int

const (
	Deposit Action = iota
	Withdraw
)

type BankAccountCommand struct {
	account   *BankAccount
	action    Action
	amount    int
	succeeded bool
}

func (b *BankAccountCommand) Call() {
	switch b.action {
	case Deposit:
		b.account.Deposit(b.amount)
		b.succeeded = true
		return

	case Withdraw:
		b.succeeded = b.account.Withdraw(b.amount)
		return

	default:
		panic("Action not recognized")
	}
}

func (b *BankAccountCommand) Undo() {
	if !b.succeeded {
		return
	}

	switch b.action {
	case Deposit:
		b.account.Withdraw(b.amount)
		return
	case Withdraw:
		b.account.Deposit(b.amount)
		return
	}
}

func (b *BankAccountCommand) Succeeded() bool {
	return b.succeeded
}
func (b *BankAccountCommand) SetSucceeded(value bool) {
	b.succeeded = value
}

func NewBankAccountCommand(account *BankAccount, action Action, amount int) *BankAccountCommand {
	return &BankAccountCommand{
		account:   account,
		action:    action,
		amount:    amount,
		succeeded: false,
	}
}

// composite command
type CompositeBankAccountCommand struct {
	commands []Command
}

func (c *CompositeBankAccountCommand) Call() {
	for _, cmd := range c.commands {
		cmd.Call()
		if !cmd.Succeeded() {
			c.Undo()
			c.SetSucceeded(false)
			return
		}
	}
}

func (c *CompositeBankAccountCommand) Undo() {
	//	undoing from last called ( like stack )
	for idx := range c.commands {
		c.commands[len(c.commands)-idx-1].Undo()
	}
}

func (c *CompositeBankAccountCommand) Succeeded() bool {
	for _, cmd := range c.commands {
		if !cmd.Succeeded() {
			return false
		}
	}
	return true
}

func (c *CompositeBankAccountCommand) SetSucceeded(value bool) {
	for _, cmd := range c.commands {
		cmd.SetSucceeded(value)
	}
}

type MoneyTransferCommand struct {
	CompositeBankAccountCommand
	from, to *BankAccount
	amount   int
}

func NewMoneyTransferCommand(from, to *BankAccount, amount int) *MoneyTransferCommand {
	mtc := &MoneyTransferCommand{
		from:   from,
		to:     to,
		amount: amount,
	}
	mtc.commands = append(mtc.commands, NewBankAccountCommand(from, Withdraw, amount))
	mtc.commands = append(mtc.commands, NewBankAccountCommand(to, Deposit, amount))

	return mtc
}

// func (m *MoneyTransferCommand) Call() {
// 	ok := true
// 	for _, cmd := range m.commands {
// 		if ok {
// 			cmd.Call()
// 			ok = cmd.Succeeded()
// 		} else {
// 			cmd.SetSucceeded(false)
// 			return
// 		}
// 	}
// }

func TestCommand() {
	ba := BankAccount{}
	NewBankAccountCommand(&ba, Deposit, 100).Call()
	NewBankAccountCommand(&ba, Deposit, 300).Call()
	NewBankAccountCommand(&ba, Withdraw, 200).Call()
	fmt.Println(ba.balance)
	c := []Command{
		NewBankAccountCommand(&ba, Deposit, 400),
		NewBankAccountCommand(&ba, Withdraw, 200),
		NewBankAccountCommand(&ba, Withdraw, 2000),
	}
	for _, cmd := range c {
		cmd.Call()
	}
	fmt.Println(ba.balance)
	for _, cmd := range c {
		cmd.Undo()
	}
	fmt.Println(ba.balance)

	//	checking composite commands
	fmt.Println("Checking composite commands")
	from := BankAccount{
		balance: 100,
	}
	to := BankAccount{
		balance: 0,
	}
	mtc := NewMoneyTransferCommand(&from, &to, 25)
	mtc.Call()
	fmt.Println(from.balance, to.balance)
	mtc.Undo()
	fmt.Println(from.balance, to.balance)

	fmt.Println()
	fmt.Println(ba.balance)
	compos := CompositeBankAccountCommand{
		commands: c,
	}
	compos.Call()
	fmt.Println(ba.balance)
	if !compos.Succeeded() {
		compos.Undo()
	}
	fmt.Println(ba.balance)
}
