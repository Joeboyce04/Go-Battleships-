package Player

import (
	table "battleships/Grid"
	"testing"
)

func TestPlaceShip(t *testing.T){
	grid := table.CreateGrid()
	desiredCol:= 3
	desiredRow:= 5
	updateGrid, _:= PlaceShips(grid, desiredCol, desiredRow)

	actual := updateGrid[desiredRow][desiredCol]
	want:= "Ship"

	if actual!=want {
		t.Error("Ship was not placed at col" ,desiredCol, "row", desiredRow )
	}
}

func TestPlayerPlaceShipInTopLeftCorner(t *testing.T) {
    
    grid := table.CreateGrid()
    
    updatedGrid, _ := PlaceShips(grid, 0, 0)
    hasShip := updatedGrid[0][0] == "Ship"
    
    if !hasShip {
        t.Error("Player one ship not found in the top left corner")
    }
}

func TestPlayerPlaceShipInTopRight(t *testing.T) {
    
    grid := table.CreateGrid()
    
    updatedGrid,_ := PlaceShips(grid,6, 0)
    hasShip := updatedGrid[0][6] == "Ship"
    
    if !hasShip {
        t.Error("Player one ship not found in the top Right corner")
    }
}

func TestPlayerPlaceShipInBottomRight(t *testing.T) {
    
    grid := table.CreateGrid()
    
    updatedGrid,_ := PlaceShips(grid, 0,6)
    hasShip := updatedGrid[6][0] == "Ship"
    
    if !hasShip {
        t.Error("Player one ship not found in the Bottom right corner")
    }
}

func TestPlayerPlaceShipInBottomLeft(t *testing.T) {
    
    grid := table.CreateGrid()
    
    updatedGrid,_ := PlaceShips(grid, 6,6)
    hasShip := updatedGrid[6][6] == "Ship"
    
    if !hasShip {
        t.Error("Player one ship not found in the bottom left corner")
    }
}

func TestPlayerPlaceShipOutsideGrid(t *testing.T) {
    grid:= table.CreateGrid()

	updatedGrid,_ := PlaceShips(grid,8,8)

	if updatedGrid != grid{
		t.Error("Player one ship is outside grid")
	}
}

func TestPlayerPlaceShipOutsideLeft(t *testing.T){
	grid := table.CreateGrid()
	
	updateGrid,_ := PlaceShips(grid, -1, 0)

	if updateGrid !=grid{
		t.Error("Player One placed ship outside grid to the left")
	}
}

func TestPlayerPlaceShipOutsideRight(t *testing.T){
	grid := table.CreateGrid()

	updateGrid,_ :=PlaceShips(grid, 7, 0)

	if updateGrid !=grid{
		t.Error("Player One placed ship outside grid to the right")
	}

} 

func TestPlayerPlaceShipOutsideBottom(t *testing.T){
	grid := table.CreateGrid()

	updateGrid,_ := PlaceShips(grid, 0, 7)

	if updateGrid !=grid{
		t.Error("Player One placed ship outside grid to the bottom")
	}
}

func TestPlayerPlaceShipOutsideTop(t *testing.T){
	grid := table.CreateGrid()

	updateGrid,_ := PlaceShips(grid, 0, -1)

	if updateGrid !=grid{
		t.Error("Player One placed ship outside grid to the top")
	}
}

func TestPlayerCannotPlaceShipOnTopOfAnother(t *testing.T) {
	grid := table.CreateGrid()
	grid,_= PlaceShips(grid, 1,2)

	col:=1
	row:=2
	_,  err := PlaceShips(grid, col, row)				

	if err==nil{
		t.Error("Cannot Place ship on top of a ship")
	}
	
} 

func TestPlayerCannotPlaceTenthShip(t *testing.T) {
	grid:= table.CreateGrid()
	
	grid,_ = PlaceShips(grid, 1, 2)
    grid, _ =  PlaceShips(grid, 3, 4)
    grid, _ = PlaceShips(grid, 6, 6)

    grid, _ = PlaceShips(grid, 0, 0)
    grid, _ = PlaceShips(grid, 0, 6)
    grid, _ = PlaceShips(grid, 6, 0)

    grid, _ = PlaceShips(grid, 1, 1)
    grid, _ = PlaceShips(grid, 5, 4)
    grid, _= PlaceShips(grid, 3, 2)

	col:= 5
	row:= 5

    _, err:= PlaceShips(grid, col, row)   

    if err==nil{
        t.Error("Player One Can place ten ships")

	}
}
    
func TestPlayerHit(t *testing.T) {
	grid := table.CreateGrid()
	grid, _ = PlaceShips(grid,5,6)
	
	col := 5
	row := 6
	
	ShotResult:= PlayerTakeShots(grid, col, row)
	


	if ShotResult.Result !="Hit" {
		t.Error("Player 1's shot did not result in a hit")
	}
}

func TestPlayerMiss(t *testing.T) {
	grid := table.CreateGrid()
	grid,_ = PlaceShips(grid, 5,5)
	
	col := 4
	row := 4
	
	ShotResult:= PlayerTakeShots(grid, col, row)
	


	if ShotResult.Result !="Miss" {
		t.Error("Player 1's shot should have missed")
	}
}

func TestPlayerShootingOutsideGrid(t *testing.T) {
	grid := table.CreateGrid()
    
    col := 8
    row := 8
    
    ShotResult := PlayerTakeShots(grid, col, row)          
    

   
		if ShotResult.Result !="Error outside grid" {
			t.Error("Player 1's shot was outside the grid")
		}
} 


func TestPlayerCannotShootSameShipLocationTwice(t *testing.T) {
		
	grid := table.CreateGrid()
	updatedGrid, _ := PlaceShips(grid, 5, 5) 
	
	firstShotResult := PlayerTakeShots(updatedGrid, 5, 5)
	secondShotResult := PlayerTakeShots(updatedGrid, 5, 5)
	
	if firstShotResult.Result != "Hit" {
		t.Error("Player's first shot should be a hit")
	}
	
	if secondShotResult.Result == "Already Hit" {
		t.Error("Player should not be able to shoot the same ship location twice")
	} 
}

func TestPlayerShipsSunkTwice(t *testing.T){
	grid:= table.CreateGrid()
	col:= 2
	row:= 3
	grid, _  = PlaceShips(grid, col, row)

	PlayerTakeShots(grid, col, row)

	shotResult:= PlayerTakeShots(grid, col, row)

	if shotResult.Result=="Already Hit"{
		t.Error("Player One should not be able to have its ship sank twice")
	}
} 


