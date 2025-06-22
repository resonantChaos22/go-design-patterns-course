- `Command` is an object which represents an instruction to perform a particular action. It contains all the information necessary to be taken
- This is useful in cases where we might need to redo or undo any set of tasks.
- Or we want to save a list of commands as `Macros

## Bank Account

- Let's take a bank account command into example

```go
type Command interface {
 Call()
}

type BankAccountCommand struct {
 account *BankAccount
 action  Action
 amount  int
}

func (b *BankAccountCommand) Call() {
 switch b.action {
 case Deposit:
  b.account.Deposit(b.amount)
  return

 case Withdraw:
  b.account.Withdraw(b.amount)
  return

 default:
  panic("Action not recognized")
 }
}
```

- Here `Command` is an interface of Command with a `Call()` function.
- This can be used to create a list of commands that need to be called as `Macros` or implement undo function

## Undo Operation

```go
type Command interface {
 Call()
 Undo()
}

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
```

- We do some changes in the code to undo all the commands that have been executed and handle the case in which a command failed (withdraw can fail in case of insufficient balance)

## Macros / Composite Command

- List of commands that need to be executed.

```go
type Command interface {
 Call()
 Undo()
 Succeeded() bool
 SetSucceeded(value bool)
}

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
 // undoing from last called ( like stack )
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
```

- This is a normal `Composite Command` which works on a list of commands.
- As you can see, if the `Call()` function of any previous commands fail, it wont go further and undo any previous changes.
- This could be used to implement transactions as well.
- And `Command` pattern can be used for functional builder as well.
- A money transfer command from one bank to another -

```go
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
```

## Functional Command

- This is a very dumbdown approach if we dont need to make it so robust.

```go
var commands []func()
commands = append(commands, func() {
 Deposit(&ba, 100)
})
commands = append(commands, func() {
 Withdraw(&ba, 25)
})

for _, cmd := range commands {
 cmd()
}
```

- This also works as a list of commands but we lose much of the robustness of the previous approach
