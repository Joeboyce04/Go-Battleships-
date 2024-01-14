package main

import (
	"math/rand"
	"testing"
)

//you can run all you tests by typing
//go test -v
//in the terminal window

//this is a utility function for testing
//it will return a random square on the grid
//it does not keep track of any previously returned grids
func getRandomGridSquare() []int {
	
	row := []int{1,2,3,4,5,6,7,8,9}
	column := []int{1,2,3,4,5,6,7,8,9}

	return []int{rand.Intn(len(row))+1,rand.Intn(len(column))+1}

}

//these are the two tests we have for our functions in main
//the purpose of tests is to mimic interaction with our code
//there is no "user input" - the test is the calling code

//here is an example of a failing test - what do we need to do to fix it?
func TestCreateGrid(t *testing.T){
	grid := CreateGrid()
	assertGrisIsCorrectSize(t, grid, 7, 7)
	}


func assertGrisIsCorrectSize(t *testing.T, grid [7][7]string, expectedRows int, expectedCols int){
	gridSizeCols := len(grid)
	if gridSizeCols != expectedCols{
		t.Error("Grid is wrong size. Expected exact size of 7, got: ", gridSizeCols)
	}

	gridSizeRows:= len(grid[0])
	if gridSizeRows!= expectedRows{
		t.Error("Grid has wrong number of rows. Want 7, but was", gridSizeRows)
	}
}


func TestPlaceShip(t *testing.T){
	grid := CreateGrid()
	desiredCol:= 3
	desiredRow:= 5
	updateGrid:= placeShip(grid, desiredCol, desiredRow)

	actual := updateGrid[desiredRow][desiredCol]
	want:= "Ship"

	if actual!=want {
		t.Error("Ship was not placed at col" ,desiredCol, "row", desiredRow )
	}
}

func TestPlayerOneShipInTopLeftCorner(t *testing.T) {
    
    grid := CreateGrid()
    
    updatedGrid := PlayerOnePlaceShips(grid)
    hasShip := updatedGrid[0][0] == "Ship"
    
    if !hasShip {
        t.Error("Player one ship not found in the top left corner")
    }
}

func TestPlayerOneShipInTopRight(t *testing.T) {
    
    grid := CreateGrid()
    
    updatedGrid := PlayerOnePlaceShips(grid)
    hasShip := updatedGrid[0][6] == "Ship"
    
    if !hasShip {
        t.Error("Player one ship not found in the top left corner")
    }
}
func TestPlayerOneShipInBottomRight(t *testing.T) {
    
    grid := CreateGrid()
    
    updatedGrid := PlayerOnePlaceShips(grid)
    hasShip := updatedGrid[6][0] == "Ship"
    
    if !hasShip {
        t.Error("Player one ship not found in the top left corner")
    }
}
func TestPlayerOneShipInBottomLeft(t *testing.T) {
    
    grid := CreateGrid()
    
    updatedGrid := PlayerOnePlaceShips(grid)
    hasShip := updatedGrid[6][6] == "Ship"
    
    if !hasShip {
        t.Error("Player one ship not found in the top left corner")
    }
}


/*func TestPlayerOnePlaceShipOutsideGrid(t *testing.T) {
    grid:=CreateGrid()

	updatedGrid:= placeShip(grid,8,8)

	if updatedGrid != grid{
		t.Error("Player one ship is outside grid")
	}
}

func TestPlayerTwoPlaceShipOutsideGrid(t *testing.T){
	grid:=CreateGrid()

	updatedGrid:= placeShip(grid,10,10)

	if updatedGrid != grid{
		t.Error("Player Two ship is outside grid")
	}

} */


func TestPlaceShipOutsideLeft(t *testing.T){
	grid :=CreateGrid()
	
	updateGrid:= placeShip(grid, -1, 0)

	if updateGrid !=grid{
		t.Error("Player placed ship outside grid to the left")
	}
}

func TestPlaceShipOutsideRight(t *testing.T){
	grid :=CreateGrid()

	updateGrid:= placeShip(grid, 7, 0)

	if updateGrid !=grid{
		t.Error("Player placed ship outside grid to the right")
	}

} 

func TestPlaceShipOutsideBottom(t *testing.T){
	grid :=CreateGrid()

	updateGrid:= placeShip(grid, 0, 7)

	if updateGrid !=grid{
		t.Error("Player placed ship outside grid to the bottom")
	}
}

func TestPlaceShipOutsideTop(t *testing.T){
	grid :=CreateGrid()

	updateGrid:= placeShip(grid, 0, -1)

	if updateGrid !=grid{
		t.Error("Player placed ship outside grid to the top")
	}
}

func TestPlayerOneCannotPlaceShipOnTopOfAnother(t *testing.T) {
}
func TestPlayerOneShipsSunkTwice(t *testing.T){
	
	}




func TestWinCondition(t *testing.T){} 

func TestPlayersCanPlaceNineShips(t *testing.T){
	grid:= CreateGrid()
	updateGrid:= PlayerOnePlaceShips(grid)

	shipCount:= CountOfShipOnBoard(updateGrid)
	if shipCount!=9{
		t.Error("Expected nine ships to be placed")
	}
}
func TestPlacingA10thShipDoesntChangeGrid(t *testing.T){}
func TurnManagement(t *testing.T){}








	





	

func TestPlayerOneCannotPlaceTenthShip(t *testing.T) {
	grid:= CreateGrid()
	updateGrid:= PlayerOnePlaceShips(grid)

	shipCount:= CountOfShipOnBoard(updateGrid)
	if shipCount>9{
		t.Error("Cant place more than 9 ships")
	}
}
    



	
	






	func TestPlayerOneHit(t *testing.T) {
		grid := CreateGrid()
		grid = PlayerTwoPlaceShips(grid)
		
		col := 5
		row := 6
		
		ShotResult:= PlayerOneTakeShots(grid, col, row)
		
	

		if ShotResult.Result !="Hit" {
			t.Error("Player 1's shot did not result in a hit")
		}
	}
	
	func TestPlayerTwoHit(t *testing.T) {
		grid := CreateGrid()
		grid = PlayerOnePlaceShips(grid)
		
		col := 1
		row := 2
		
		ShotResult:= PlayerTwoTakeShots(grid, col, row)
		
	

		if ShotResult.Result !="Hit" {
			t.Error("Player 2's shot did not result in a hit")
		}
	}

	 func TestPlayerOneMiss(t *testing.T) {
		grid := CreateGrid()
		grid = PlayerTwoPlaceShips(grid)
		
		col := 4
		row := 4
		
		ShotResult:= PlayerOneTakeShots(grid, col, row)
		
	

		if ShotResult.Result !="Miss" {
			t.Error("Player 1's shot should have missed")
		}
	}
	

	func TestPlayerTwoMiss(t *testing.T) {
		grid := CreateGrid()
		grid = PlayerOnePlaceShips(grid)
		
		col := 4
		row := 3
		
		ShotResult:= PlayerTwoTakeShots(grid, col, row)
		
	

		if ShotResult.Result !="Miss" {
			t.Error("Player 2's shot should have missed")
		}
	}

func TestPlayerOneShootingOutsideGrid(t *testing.T) {
	grid := CreateGrid()
    
    col := 8
    row := 8
    
    ShotResult := PlayerOneTakeShots(grid, col, row)          
    

   
		if ShotResult.Result !="Error outside grid" {
			t.Error("Player 1's shot was outside the grid")
		}
} 

func TestPlayerTwoShootingOutsideGrid(t *testing.T) {
		grid := CreateGrid()

		col := 8
		row := -6

		ShotResult:= PlayerTwoTakeShots(grid, col, row)
		
		
		if ShotResult.Result !="Error outside grid" {
			t.Error("Player 2's shot was outside the grid")
		}
	}






