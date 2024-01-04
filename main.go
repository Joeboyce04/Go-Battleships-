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
	return grid
}

func PlayerTwoPlaceShips(grid [7][7]string) [7][7]string {
	grid = placeShip(grid, 2, 3)
	grid = placeShip(grid, 5, 6)
	return grid
}

func PlayerOneTakeShots(grid [7][7]string, col, row int) ([7][7]string, string, bool) {
	if col < 0 || col >= 7 || row < 0 || row >= 7 {
		return grid, "error", false
	}
	result := "*"
	hit := false
	if grid[row][col] == "Hit" {
		result = "Already hit"
	} else if grid[row][col] == "Ship" {
		grid[row][col] = "Hit"
		result = "Hit"
		hit = true
	} else {
		grid[row][col] = "Miss"
		result = "Miss"
	}
	return grid, result, hit
}

func PlayerTwoTakeShots(grid [7][7]string, col, row int) ([7][7]string, string, bool) {
	if col < 0 || col >= 7 || row < 0 || row >= 7 {
		return grid, "error", false
	}
	result := "*"
	hit := false
	if grid[row][col] == "Hit" {
		result = "Already hit"
	} else if grid[row][col] == "Ship" {
		grid[row][col] = "Hit"
		result = "Hit"
		hit = true
	} else {
		grid[row][col] = "Miss"
		result = "Miss"
	}
	return grid, result, hit

}

//func allShipsHit(grid [7][7]string) bool {}
