package functional

import "fmt"

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
