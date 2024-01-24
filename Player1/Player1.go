package Player1

import table "battleships/Grid"

func PlayerOnePlaceShips(grid [7][7]string, col, row int)  [7][7]string {
	PlayerOnegrid, _ := table.PlaceShip(grid, col, row )
	return PlayerOnegrid
}


func PlayerOneTakeShots(grid [7][7]string, col, row int) table.Shot {
	return table.PlayerTakeShot(grid, col, row)
}
