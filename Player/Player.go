package Player

import (
	"errors"
)

type Shot struct {
	Result string
}

func validCoordinates(col, row int)bool{
	return col<0 || col>6 || row<0  || row>6
}

func PlaceShips(grid [7][7]string, col, row int) ([7][7]string, error){
	if validCoordinates(col, row){
		return grid, errors.New("ship was placed outside the grid")
	}

	if grid[row][col] == "Ship" {
		return grid, errors.New("ship cannot be placed on top of another ship")
	}  

	shipCount := CountOfShipsOnBoard(grid)
	if shipCount == 9 {
		return grid, errors.New("too many ships trying to be placed")
	}

	grid[row][col] = "Ship"
	return grid, nil
}


func PlayerTakeShots(grid [7][7]string, col, row int) Shot {
	if validCoordinates(col, row){
		return Shot{Result: "Error outside grid"} 
	}
	shot := Shot{Result: ""}

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

func CountOfShipsOnBoard(grid [7][7]string) int {
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