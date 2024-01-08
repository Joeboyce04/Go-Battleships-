package main

func CreateGrid() [7][7]string {

	var grid [7][7]string
	return grid
}

func placeShip(grid [7][7]string, col, row int) [7][7]string {
	if col < 0 || col >= 7 || row < 0 || row >= 7 {
		return grid
	}

	grid[row][col] = "Ship"

	return grid
}

func PlayerOnePlaceShips(grid [7][7]string) [7][7]string {
	grid = placeShip(grid, 1, 2)
	grid = placeShip(grid, 3, 4)
	grid = placeShip(grid, 6, 6)
	grid = placeShip(grid, 0, 0)
	grid = placeShip(grid, 0, 6)
	grid = placeShip(grid, 6, 0)
	grid = placeShip(grid, 8, 0)

	return grid
}

func PlayerTwoPlaceShips(grid [7][7]string) [7][7]string {
	grid = placeShip(grid, 2, 3)
	grid = placeShip(grid, 5, 6)
	grid = placeShip(grid, 10, 10)

	return grid
}

func PlayerOneTakeShots(grid [7][7]string, col, row int) [7][7]string {
	if col < 0 || col >= 7 || row < 0 || row >= 7 {
		return grid
	}
	if grid[row][col] == "Ship" {
		grid[row][col] = "Hit"
	} else if grid[row][col] != "Hit" {
		grid[row][col] = "Miss"
	}
	return grid
}

func PlayerTwoTakeShots(grid [7][7]string, col, row int) [7][7]string {
	if col < 0 || col >= 7 || row < 0 || row >= 7 {
		return grid
	}
	if grid[row][col] == "Ship" {
		grid[row][col] = "Hit"
	} else if grid[row][col] != "Hit" {
		grid[row][col] = "Miss"
	}
	return grid
}

//func allShipsHit(grid [7][7]string) bool {}
