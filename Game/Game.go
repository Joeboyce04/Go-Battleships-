package Game

func PlayerTurn(playerOneTurn bool) string {
	var currentPlayerTurn string
	if playerOneTurn {
		currentPlayerTurn = "Player One"
	} else {
		currentPlayerTurn = "Player Two"
	}
	return currentPlayerTurn
}

func passTurnToOtherPlayer(playerOneTurn bool) (bool, string) {
	playerOneTurn = !playerOneTurn
	currentPlayerTurn := PlayerTurn(playerOneTurn)
	return playerOneTurn, currentPlayerTurn
}

func PlayerWon(grid [7][7]string) bool {
	for _, row := range grid {
		for _, location := range row {
			if location == "Ship" {
				return false
			}

		}
	}
	return true
}

func CountOfShipsOnBoard(grid [7][7]string) int {
	count := 0

	for _, row := range grid {
		for _, location := range row {
			if location == "Ship" {
				count++
			}

		}
	}
	return count
} 