package Player1

import (
	table "battleships/Grid"
	"battleships/Player2"
	"testing"
)



func TestPlayerOneShipInTopLeftCorner(t *testing.T) {
    
    grid := table.CreateGrid()
    
    updatedGrid := PlayerOnePlaceShips(grid, 0, 0)
    hasShip := updatedGrid[0][0] == "Ship"
    
    if !hasShip {
        t.Error("Player one ship not found in the top left corner")
    }
}

func TestPlayerOneShipInTopRight(t *testing.T) {
    
    grid := table.CreateGrid()
    
    updatedGrid := PlayerOnePlaceShips(grid,6, 0)
    hasShip := updatedGrid[0][6] == "Ship"
    
    if !hasShip {
        t.Error("Player one ship not found in the top Right corner")
    }
}

func TestPlayerOneShipInBottomRight(t *testing.T) {
    
    grid := table.CreateGrid()
    
    updatedGrid := PlayerOnePlaceShips(grid, 0,6)
    hasShip := updatedGrid[6][0] == "Ship"
    
    if !hasShip {
        t.Error("Player one ship not found in the Bottom right corner")
    }
}

func TestPlayerOneShipInBottomLeft(t *testing.T) {
    
    grid := table.CreateGrid()
    
    updatedGrid := PlayerOnePlaceShips(grid, 6,6)
    hasShip := updatedGrid[6][6] == "Ship"
    
    if !hasShip {
        t.Error("Player one ship not found in the bottom left corner")
    }
}

func TestPlayerOnePlaceShipOutsideGrid(t *testing.T) {
    grid:= table.CreateGrid()

	updatedGrid := PlayerOnePlaceShips(grid,8,8)

	if updatedGrid != grid{
		t.Error("Player one ship is outside grid")
	}
}

func TestPlayerOnePlaceShipOutsideLeft(t *testing.T){
	grid := table.CreateGrid()
	
	updateGrid := PlayerOnePlaceShips(grid, -1, 0)

	if updateGrid !=grid{
		t.Error("Player One placed ship outside grid to the left")
	}
}

func TestPlayerOnePlaceShipOutsideRight(t *testing.T){
	grid := table.CreateGrid()

	updateGrid:= PlayerOnePlaceShips(grid, 7, 0)

	if updateGrid !=grid{
		t.Error("Player One placed ship outside grid to the right")
	}

} 

func TestPlayerOnePlaceShipOutsideBottom(t *testing.T){
	grid := table.CreateGrid()

	updateGrid := PlayerOnePlaceShips(grid, 0, 7)

	if updateGrid !=grid{
		t.Error("Player One placed ship outside grid to the bottom")
	}
}

func TestPlayerOnePlaceShipOutsideTop(t *testing.T){
	grid := table.CreateGrid()

	updateGrid := PlayerOnePlaceShips(grid, 0, -1)

	if updateGrid !=grid{
		t.Error("Player One placed ship outside grid to the top")
	}
}

func TestPlayerOneCannotPlaceShipOnTopOfAnother(t *testing.T) {
	grid := table.CreateGrid()
	grid= PlayerOnePlaceShips(grid, 1,2)

	col:=1
	row:=2
	_,  err := table.PlaceShip(grid, col, row)				//TWEAK TO PLAYER ONE

	if err==nil{
		t.Error("Cannot Place ship on top of a ship")
	}
	
} 

func TestPlayerOneCannotPlaceTenthShip(t *testing.T) {
	grid:= table.CreateGrid()
	
	grid,_ = table.PlaceShip(grid, 1, 2)
    grid, _ =  table.PlaceShip(grid, 3, 4)
    grid, _ = table.PlaceShip(grid, 6, 6)

    grid, _ = table.PlaceShip(grid, 0, 0)
    grid, _ = table.PlaceShip(grid, 0, 6)
    grid, _ = table.PlaceShip(grid, 6, 0)

    grid, _ = table.PlaceShip(grid, 1, 1)
    grid, _ = table.PlaceShip(grid, 5, 4)
    grid, _= table.PlaceShip(grid, 3, 2)

	col:= 5
	row:= 5

    _, err:= table.PlaceShip(grid, col, row)   //TWEAK to player one

    if err==nil{
        t.Error("Player One Can place ten ships")

	}
}
    
func TestPlayerOneHit(t *testing.T) {
	grid := table.CreateGrid()
	grid = Player2.PlayerTwoPlaceShips(grid,5,6)
	
	col := 5
	row := 6
	
	ShotResult:= PlayerOneTakeShots(grid, col, row)
	


	if ShotResult.Result !="Hit" {
		t.Error("Player 1's shot did not result in a hit")
	}
}

func TestPlayerOneMiss(t *testing.T) {
	grid := table.CreateGrid()
	grid = Player2.PlayerTwoPlaceShips(grid, 5,5)
	
	col := 4
	row := 4
	
	ShotResult:= PlayerOneTakeShots(grid, col, row)
	


	if ShotResult.Result !="Miss" {
		t.Error("Player 1's shot should have missed")
	}
}

func TestPlayerOneShootingOutsideGrid(t *testing.T) {
	grid := table.CreateGrid()
    
    col := 8
    row := 8
    
    ShotResult := PlayerOneTakeShots(grid, col, row)          
    

   
		if ShotResult.Result !="Error outside grid" {
			t.Error("Player 1's shot was outside the grid")
		}
} 


func TestPlayerOneCannotShootSameShipLocationTwice(t *testing.T) {
		
	grid := table.CreateGrid()
	updatedGrid, _ := table.PlaceShip(grid, 5, 5) 
	
	firstShotResult := PlayerOneTakeShots(updatedGrid, 5, 5)
	secondShotResult := PlayerOneTakeShots(updatedGrid, 5, 5)
	
	if firstShotResult.Result != "Hit" {
		t.Error("Player's first shot should be a hit")
	}
	
	if secondShotResult.Result == "Already Hit" {
		t.Error("Player should not be able to shoot the same ship location twice")
	} 
}

func TestPlayerOneShipsSunkTwice(t *testing.T){
	grid:= table.CreateGrid()
	col:= 2
	row:= 3
	grid  = PlayerOnePlaceShips(grid, col, row)

	Player2.PlayerTwoTakeShots(grid, col, row)

	shotResult:= Player2.PlayerTwoTakeShots(grid, col, row)

	if shotResult.Result=="Already Hit"{
		t.Error("Player One should not be able to have its ship sank twice")
	}
} 
