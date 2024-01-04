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

func TestPlayerOnePlaceShipOutsideGrid(t *testing.T){
	grid:= CreateGrid()

	updatedGrid:= placeShip(grid, 8, 8)
	if updatedGrid!= grid{
		t.Error("Player one Ship placed outside grid")
	}
	
}




func TestPlayerTwoPlaceShipOutsideGrid(t *testing.T){
	grid:= CreateGrid()

	updatedGrid:= placeShip(grid, -1, 8)

	if updatedGrid!= grid{
		t.Error("Player Two Ship placed outside grid")
	}
	
}

//func TestPlayerOneCannotPlaceShipOnTop(t *testing.T){}
//func TestPlayerTwoCannotPlaceShipOnTop(t *testing.T){}


	

	//func TestPlayerOneCantPlaceTenthShip(t *testing.t){}
	//func TestPlayerTwoCantPlaceTenthShip(t *testing.T){}
	
	






	func TestPlayerOneHit(t *testing.T) {
		grid := CreateGrid()
		grid = PlayerOnePlaceShips(grid)
		
		col := 6
		row := 5
		
		_, result, hit := PlayerOneTakeShots(grid, col, row)
		
		wantResult := "Hit"
		wantHit := true
		if result != wantResult || hit != wantHit {
			t.Error("Player 1's shot did not result in a hit")
		}
	}
	
	func TestPlayerTwoHit(t *testing.T) {
		grid := CreateGrid()
		grid = PlayerTwoPlaceShips(grid)
		
		col := 2
		row := 1
		
		_, result, hit := PlayerTwoTakeShots(grid, col, row)
		
		wantResult := "Hit"
		wantHit := true
		if result != wantResult || hit != wantHit{
			t.Error("Player 2's shot did not result in a hit")
		}
	}

	func TestPlayerOneMiss(t *testing.T) {
		grid := CreateGrid()
		
		col := 1
		row := 2
		
		_, result, hit := PlayerOneTakeShots(grid, col, row)
		
		wantResult := "Miss"
		wantHit := false
		if result != wantResult || hit != wantHit {
			t.Error("Player 1's shot did not result in a miss")
		}
	}
	

	func TestPlayerTwoMiss(t *testing.T) {
		grid := CreateGrid()
		
		col := 1
		row := 2
		
		_, result, hit := PlayerTwoTakeShots(grid, col, row)
		
		wantResult := "Miss"
		wantHit := false
		if result != wantResult || hit != wantHit {
			t.Error("Player 2's shot did not result in a miss")
		}
	}

func TestPlayerOneShootingOutsideGrid(t *testing.T) {
	grid := CreateGrid()
    
    col := -1
    row := 8
    
    _, result, _ := PlayerOneTakeShots(grid, col, row)
    
    wantResult := "error"
    if result  != wantResult {
        t.Error("Player 1 shot was out of bounds")
    }
}
func TestPlayerTwoShootingOutsideGrid(t *testing.T) {
		grid := CreateGrid()

		col := 8
		row := -6

		_, result, _ := PlayerTwoTakeShots(grid, col, row)
		
		wantResult := "error"
		if result != wantResult {
			t.Error("Player 2 shot was outside the grid")
		}
	}




//func TestPlayerOneShipsSunkTwice(t *testing.T){}
//func TestPlayerTwoShipsSunkTwice(t *testing.T){}


/*func TestWinCondition(t *testing.T){
	grid:= CreateGrid()
} */





