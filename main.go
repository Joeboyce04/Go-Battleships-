package main

type Shot struct {
	Hit    bool
	Result string
}

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
	grid = placeShip(grid, 1, 1)
	grid = placeShip(grid, 5, 4)
	grid = placeShip(grid, 3, 2)

	return grid
}

func PlayerTwoPlaceShips(grid [7][7]string) [7][7]string {
	grid = placeShip(grid, 2, 3)
	grid = placeShip(grid, 5, 6)

	return grid
}

func PlayerOneTakeShots(grid [7][7]string, col, row int) Shot {
	if col < 0 || col >= 7 || row < 0 || row >= 7 {
		return Shot{Result: "Error shot outside Grid", Hit: false}
	}
	shot := Shot{Result: "", Hit: false}

	if grid[row][col] == "Hit" {
		shot.Result = "Already Hit"
		shot.Hit = false
	} else if grid[row][col] == "Ship" {
		grid[row][col] = "Hit"
		shot.Result = "Hit"
		shot.Hit = true
	} else {
		grid[row][col] = "Miss"
		shot.Result = "Miss"
		shot.Hit = false
	}
	return shot
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
