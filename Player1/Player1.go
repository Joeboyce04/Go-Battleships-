package Player1

import table "battleships/Grid"

func PlayerOnePlaceShips(grid [7][7]string, col, row int)  [7][7]string {
	grid, _ = table.PlaceShip(grid, col, row )
	return grid
}

func PlayerOneTakeShots(grid [7][7]string, col, row int) table.Shot {
	return table.PlayerTakeShot(grid, col, row)
}
