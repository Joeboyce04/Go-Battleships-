package Player2

import (
	table "battleships/Grid"

	"testing"

	"battleships/Player1"
)

func TestPlayerTwoShipInTopLeftCorner(t *testing.T) {
    
    grid := table.CreateGrid()
    
    updatedGrid := PlayerTwoPlaceShips(grid, 0, 0)
    hasShip := updatedGrid[0][0] == "Ship"
    
    if !hasShip {
        t.Error("Player Twos ship not found in the top left corner")
    }
}

func TestPlayerTwoShipInTopRight(t *testing.T) {
    
    grid := table.CreateGrid()
    
    updatedGrid := PlayerTwoPlaceShips(grid,6, 0)
    hasShip := updatedGrid[0][6] == "Ship"
    
    if !hasShip {
        t.Error("Player Two ship not found in the top Right corner")
    }
}


func TestPlayerTwoShipInBottomRight(t *testing.T) {
    
    grid := table.CreateGrid()
    
    updatedGrid := PlayerTwoPlaceShips(grid, 0,6)
    hasShip := updatedGrid[6][0] == "Ship"
    
    if !hasShip {
        t.Error("Player Two ship not found in the Bottom right corner")
    }
}

func TestPlayerTwoShipInBottomLeft(t *testing.T) {
    
    grid := table.CreateGrid()
    
    updatedGrid := PlayerTwoPlaceShips(grid, 6,6)
    hasShip := updatedGrid[6][6] == "Ship"
    
    if !hasShip {
        t.Error("Player Two ship not found in the bottom left corner")
    }
}


func TestPlayerTwoPlaceShipOutsideGrid(t *testing.T){
	grid:= table.CreateGrid()

	updatedGrid, _:= table.PlaceShip(grid,10,10)

	if updatedGrid != grid{
		t.Error("Player Two ship is outside grid")
	}

}

func TestPlayerTwoPlaceShipOutsideLeft(t *testing.T){
	grid := table.CreateGrid()

	updateGrid := PlayerTwoPlaceShips(grid, -1, 0)

	if updateGrid !=grid{
		t.Error("Player Two placed ship outside grid to the left")
	}
}

func TestPlayerTwpPlaceShipOutsideRight(t *testing.T){
	grid :=table.CreateGrid()

	updateGrid:= PlayerTwoPlaceShips(grid, 7, 0)

	if updateGrid !=grid{
		t.Error("Player Two placed ship outside grid to the right")
	}

} 


func TestPlayerTwoPlaceShipOutsideBottom(t *testing.T){
	grid :=table.CreateGrid()

	updateGrid := PlayerTwoPlaceShips(grid, 0, 7)

	if updateGrid !=grid{
		t.Error("Player Two placed ship outside grid to the bottom")
	}
}


func TestPlayerTwoPlaceShipOutsideTop(t *testing.T){
	grid :=table.CreateGrid()

	updateGrid := PlayerTwoPlaceShips(grid, 0, -1)

	if updateGrid !=grid{
		t.Error("Player Two placed ship outside grid to the top")
	}
}

func TestPlayerTwoHit(t *testing.T) {
	grid := table.CreateGrid()
	grid = Player1.PlayerOnePlaceShips(grid, 1, 2)
	
	col := 1
	row := 2
	
	ShotResult:= PlayerTwoTakeShots(grid, col, row)
	


	if ShotResult.Result !="Hit" {
		t.Error("Player 2's shot did not result in a hit")
	}
}


func TestPlayerTwoMiss(t *testing.T) {
	grid := table.CreateGrid()
	grid = Player1.PlayerOnePlaceShips(grid, 1, 1)
	
	col := 4
	row := 3
	
	ShotResult:= PlayerTwoTakeShots(grid, col, row)
	


	if ShotResult.Result !="Miss" {
		t.Error("Player 2's shot should have missed")
	}
}

func TestPlayerTwoShootingOutsideGrid(t *testing.T) {
	grid := table.CreateGrid()

	col := 8
	row := -6

	ShotResult:= PlayerTwoTakeShots(grid, col, row)
	
	
	if ShotResult.Result !="Error outside grid" {
		t.Error("Player 2's shot was outside the grid")
	}
}


//ADD REST OF PLAYER TWO CODE