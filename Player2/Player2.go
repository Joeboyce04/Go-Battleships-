package Player2

import table "battleships/Grid"

func PlayerTwoPlaceShips(grid [7][7]string, col, row int) [7][7]string {
	PlayerTwogrid, _ := table.PlaceShip(grid, col, row)
	return PlayerTwogrid
}

func PlayerTwoTakeShots(grid [7][7]string, col, row int) table.Shot {
	return table.PlayerTakeShot(grid, col, row)
}