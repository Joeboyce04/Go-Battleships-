package Game

func PlayerTurn(playerOneTurn bool) string {
	if playerOneTurn {
		return "Player One"
	}
	return "Player Two"

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
