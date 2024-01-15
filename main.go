package main

type Shot struct {
	Result string
}

func CreateGrid() [7][7]string {

	var grid [7][7]string
	return grid
}

func CountOfShipOnBoard(grid [7][7]string) int {
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
func placeShip(grid [7][7]string, col, row int) [7][7]string {
	if col < 0 || col > 6 || row < 0 || row > 6 {
		return grid
	}

	if grid[row][col] == "Ship" {
		return grid
	}

	shipCount := CountOfShipOnBoard(grid)
	if shipCount == 9 {
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
	grid = placeShip(grid, 3, 1)

	return grid
}

func PlayerTwoPlaceShips(grid [7][7]string) [7][7]string {
	grid = placeShip(grid, 2, 3)
	grid = placeShip(grid, 5, 6)

	return grid
}

func PlayerTakeShot(grid [7][7]string, col, row int) Shot {
	if col < 0 || col > 6 || row < 0 || row > 6 {
		return Shot{Result: "Error outside grid"}
	}
	shot := Shot{Result: "Miss"}

	if grid[row][col] == "Hit" {
		shot.Result = "Already Hit"
	} else if grid[row][col] == "Ship" {
		grid[row][col] = "Hit"
		shot.Result = "Hit"
	} else {
		grid[row][col] = "Miss"
		shot.Result = "Miss"
	}
	return shot
}

func PlayerOneTakeShots(grid [7][7]string, col, row int) Shot {
	return PlayerTakeShot(grid, col, row)
}

func PlayerTwoTakeShots(grid [7][7]string, col, row int) Shot {
	return PlayerTakeShot(grid, col, row)
}
func ShipSunk(grid [7][7]string, col, row int) bool {
	return grid[row][col] == "Hit"
}

func allShipsHit(grid [7][7]string) bool {
	for _, row := range grid {
		for _, location := range row {
			if location == "Ship" {
				return false
			}
		}
	}
	return true
}
