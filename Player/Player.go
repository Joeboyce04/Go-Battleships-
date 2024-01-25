package Player

import (
	game "battleships/Game"
	"errors"
)

type Shot struct {
	Result string
}

func PlaceShips(grid [7][7]string, col, row int) ([7][7]string, error){
	if col < 0 || col > 6 || row < 0 || row > 6 {
		return grid, errors.New("Ship was placed outside the grid")
	}

	if grid[row][col] == "Ship" {
		return grid, errors.New("Ship cannot be placed on top of another ship")
	}  

	shipCount := game.CountOfShipsOnBoard(grid)
	if shipCount == 9 {
		return grid, errors.New("Too many ships trying to be placed")
	}

	grid[row][col] = "Ship"
	return grid, nil
}


func PlayerTakeShots(grid [7][7]string, col, row int) Shot {
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