package Player1

func PlayerOnePlaceShips(grid [7][7]string, col, row int)  [7][7]string {
	PlayerOnegrid, _ := grid.PlaceShip(grid, col, row )
	return PlayerOnegrid
}




