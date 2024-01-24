package grid

import (
	"errors"
)


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

func PlaceShip(grid [7][7]string, col, row int) ([7][7]string, error){
	if col < 0 || col > 6 || row < 0 || row > 6 {
		return grid, errors.New("Ship was placed outside the grid")
	}

	if grid[row][col] == "Ship" {
		return grid, errors.New("Ship cannot be placed on top of another ship")
	}  

	shipCount := CountOfShipOnBoard(grid)
	if shipCount == 9 {
		return grid, errors.New("Too many ships trying to be placed")
	}

	grid[row][col] = "Ship"
	return grid, nil
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
