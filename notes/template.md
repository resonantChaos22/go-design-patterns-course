- `Template Pattern` is a skeleton algorithm defined in a function. Function can either use an interface (like `Strategy Pattern`) or can take several functions as arguments.
- This is like `Strategy` but instead of having a struct with the `Strategy` interface, it works based on a function.
- A template method is basically a function which operates on an interface, so it provides a skeleton of how exactly the interface will be used and then we get to define the concrete implementation of it.

## Structural Example -

- We are creating an example based on a `Game` which needs to have 4 functions - `Start()`, `TakeTurn()`, `HaveWinner()` and `WinningPlayer()`

```go
type Game interface {
 Start()
 TakeTurn()
 HaveWinner() bool
 WinningPlayer() int
}

func PlayGame(g Game) {
 g.Start()
 for !g.HaveWinner() {
  g.TakeTurn()
 }
 fmt.Printf("Player %d wins. \n", g.WinningPlayer())
}

type chess struct {
 turn, maxTurns, currentPlayer int
}

func (c *chess) Start() {
 fmt.Println("Starting a new game of chess")
}

func (c *chess) TakeTurn() {
 c.turn++
 fmt.Printf("Turn %d taken by player %d\n", c.turn, c.currentPlayer)
 c.currentPlayer = 1 - c.currentPlayer
}

func (c *chess) HaveWinner() bool {
 return c.turn == c.maxTurns
}

func (c *chess) WinningPlayer() int {
 return c.currentPlayer
}

func NewGameOfChess() Game {
 return &chess{
  turn:          1,
  maxTurns:      10,
  currentPlayer: 0,
 }
}
```

- As we can see here, we have defined the `PlayGame` over a `Game` and it perfectly works with `Chess` or any other turn-based games once we define their functions.

## Functional Template

- Instead of having interfaces with the methods , we can use methods directly like so -

```go
func PlayGame(start, takeTurn func(), haveWinner func() bool, winningPlayer func() int) {
 start()
 for !haveWinner() {
  takeTurn()
 }
 fmt.Printf("Player %d wins!\n", winningPlayer())
}

func TestFunctionalTemplate() {
 turns, maxTurns, currentPlayer := 1, 10, 0

 start := func() {
  fmt.Println("Starting a new game of chess")
 }
 takeTurn := func() {
  turns++
  fmt.Printf("Turn %d taken by player %d\n", turns, currentPlayer)
  currentPlayer = (currentPlayer + 1) % 2
 }
 haveWinner := func() bool {
  return turns == maxTurns
 }
 winningPlayer := func() int {
  return currentPlayer
 }

 PlayGame(start, takeTurn, haveWinner, winningPlayer)
}
```

- This approach is not that readable, but we dont need to use interfaces for template method, we can use it as a Higher Order Function
