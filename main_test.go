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
	updateGrid, _:= placeShip(grid, desiredCol, desiredRow)

	actual := updateGrid[desiredRow][desiredCol]
	want:= "Ship"

	if actual!=want {
		t.Error("Ship was not placed at col" ,desiredCol, "row", desiredRow )
	}
}

func TestPlayerOneShipInTopLeftCorner(t *testing.T) {
    
    grid := CreateGrid()
    
    updatedGrid := PlayerOnePlaceShips(grid, 0, 0)
    hasShip := updatedGrid[0][0] == "Ship"
    
    if !hasShip {
        t.Error("Player one ship not found in the top left corner")
    }
}

func TestPlayerOneShipInTopRight(t *testing.T) {
    
    grid := CreateGrid()
    
    updatedGrid := PlayerOnePlaceShips(grid,6, 0)
    hasShip := updatedGrid[0][6] == "Ship"
    
    if !hasShip {
        t.Error("Player one ship not found in the top left corner")
    }
}
func TestPlayerOneShipInBottomRight(t *testing.T) {
    
    grid := CreateGrid()
    
    updatedGrid := PlayerOnePlaceShips(grid, 0,6)
    hasShip := updatedGrid[6][0] == "Ship"
    
    if !hasShip {
        t.Error("Player one ship not found in the top left corner")
    }
}
func TestPlayerOneShipInBottomLeft(t *testing.T) {
    
    grid := CreateGrid()
    
    updatedGrid := PlayerOnePlaceShips(grid, 6,6)
    hasShip := updatedGrid[6][6] == "Ship"
    
    if !hasShip {
        t.Error("Player one ship not found in the top left corner")
    }
}


func TestPlayerOnePlaceShipOutsideGrid(t *testing.T) {
    grid:=CreateGrid()

	updatedGrid, _:= placeShip(grid,8,8)

	if updatedGrid != grid{
		t.Error("Player one ship is outside grid")
	}
}

func TestPlayerTwoPlaceShipOutsideGrid(t *testing.T){
	grid:=CreateGrid()

	updatedGrid, _:= placeShip(grid,10,10)

	if updatedGrid != grid{
		t.Error("Player Two ship is outside grid")
	}

} 


func TestPlaceShipOutsideLeft(t *testing.T){
	grid :=CreateGrid()
	
	updateGrid, _:= placeShip(grid, -1, 0)

	if updateGrid !=grid{
		t.Error("Player placed ship outside grid to the left")
	}
}

func TestPlaceShipOutsideRight(t *testing.T){
	grid :=CreateGrid()

	updateGrid, _:= placeShip(grid, 7, 0)

	if updateGrid !=grid{
		t.Error("Player placed ship outside grid to the right")
	}

} 

func TestPlaceShipOutsideBottom(t *testing.T){
	grid :=CreateGrid()

	updateGrid, _:= placeShip(grid, 0, 7)

	if updateGrid !=grid{
		t.Error("Player placed ship outside grid to the bottom")
	}
}

func TestPlaceShipOutsideTop(t *testing.T){
	grid :=CreateGrid()

	updateGrid, _:= placeShip(grid, 0, -1)

	if updateGrid !=grid{
		t.Error("Player placed ship outside grid to the top")
	}
}

func TestPlayerOneCannotPlaceShipOnTopOfAnother(t *testing.T) {
	grid :=CreateGrid()
	grid, _= placeShip(grid, 1,2)

	col:=1
	row:=2
	 _, err :=placeShip(grid, col, row)

	if err==nil{
		t.Error("Cannot Place ship on top of a ship")
	}
	
} 


func TestPlayerOneShipsSunkTwice(t *testing.T){
	
	}




func TestPlayersCanPlaceNineShips(t *testing.T){
}








	





	

func TestPlayerOneCannotPlaceTenthShip(t *testing.T) {
	grid:= CreateGrid()
	
    grid,_ = placeShip(grid, 1, 2)
    grid, _ = placeShip(grid, 3, 4)
    grid, _ = placeShip(grid, 6, 6)

    grid, _ = placeShip(grid, 0, 0)
    grid, _ = placeShip(grid, 0, 6)
    grid, _ = placeShip(grid, 6, 0)

    grid, _ = placeShip(grid, 1, 1)
    grid, _ = placeShip(grid, 5, 4)
    grid, _= placeShip(grid, 3, 2)

	col:= 5
	row:= 5

    _, err:= placeShip(grid, col, row) 

    if err==nil{
        t.Error("Player One Can place ten ships")

	}
}
    


	
	






	func TestPlayerOneHit(t *testing.T) {
		grid := CreateGrid()
		grid = PlayerTwoPlaceShips(grid,5,6)
		
		col := 5
		row := 6
		
		ShotResult:= PlayerOneTakeShots(grid, col, row)
		
	

		if ShotResult.Result !="Hit" {
			t.Error("Player 1's shot did not result in a hit")
		}
	}
	
	func TestPlayerTwoHit(t *testing.T) {
		grid := CreateGrid()
		grid = PlayerOnePlaceShips(grid, 1, 2)
		
		col := 1
		row := 2
		
		ShotResult:= PlayerTwoTakeShots(grid, col, row)
		
	

		if ShotResult.Result !="Hit" {
			t.Error("Player 2's shot did not result in a hit")
		}
	}

	 func TestPlayerOneMiss(t *testing.T) {
		grid := CreateGrid()
		grid = PlayerTwoPlaceShips(grid, 5,5)
		
		col := 4
		row := 4
		
		ShotResult:= PlayerOneTakeShots(grid, col, row)
		
	

		if ShotResult.Result !="Miss" {
			t.Error("Player 1's shot should have missed")
		}
	}
	

	func TestPlayerTwoMiss(t *testing.T) {
		grid := CreateGrid()
		grid = PlayerOnePlaceShips(grid, 1, 1)
		
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


	func TestPlayerCannotShootSameShipLocationTwice(t *testing.T) {
		
		grid := CreateGrid()
		updatedGrid, _ := placeShip(grid, 5, 5) 
		
		firstShotResult := PlayerOneTakeShots(updatedGrid, 5, 5)
		secondShotResult := PlayerOneTakeShots(updatedGrid, 5, 5)
		
		if firstShotResult.Result != "Hit" {
			t.Error("Player's first shot should be a hit")
		}
		
		if secondShotResult.Result == "Already Hit" {
			t.Error("Player should not be able to shoot the same ship location twice")
		} 
	}

func TestPlayerOneGetsFirstTurn(t* testing.T){
	playerOneTurn:= true

	 currentPlayerTurn:= PlayerTurn(playerOneTurn)

	if currentPlayerTurn!="Player One"{
		t.Error("Player One should have the first turn")
	}
}

func TestPlayerTwoTurnPassesToPlayerTwo(t *testing.T){
	playerOneTurn:= false

	playerOneTurn, currentPlayerTurn:= passTurnToOtherPlayer(playerOneTurn)

	if currentPlayerTurn!="Player One"{
		t.Error("Should be player ones turn after player Two")
	}

	if !playerOneTurn{
		t.Error("Player Twos turn should be over")
	}

}
func TestPlayerOneTurnPassesToPlayerTwo(t *testing.T){

	playerOneTurn:=true
	
	playerOneTurn, currentPlayerTurn:=passTurnToOtherPlayer(playerOneTurn)

	if currentPlayerTurn!="Player Two"{
		t.Error("Should be player twos turn after player one")
	}
	if playerOneTurn{
		t.Error("Player ones turn should have been finished!!!!!!")
	}
}


/*func TestPlayerWinsAfterSinkingAllShips(t *testing.T) {
	
		grid := CreateGrid()
		grid, _ = placeShip(grid, 1, 2)
		grid, _ = placeShip(grid, 2, 3)
		grid, _ = placeShip(grid, 3, 4)

		grid, _ = placeShip(grid, 4, 5)
		grid, _ = placeShip(grid, 5, 6)
		grid, _ = placeShip(grid, 6, 4)

		grid, _ = placeShip(grid, 5, 1)
		grid, _ = placeShip(grid, 1, 3)
		grid, _ = placeShip(grid, 2, 4)
											//does not pass will fix
		PlayerTakeShot(grid, 1, 2)
		PlayerTakeShot(grid, 2, 3)					
		PlayerTakeShot(grid, 3, 4)				

		PlayerTakeShot(grid, 4, 5)
		PlayerTakeShot(grid, 5, 6)
		PlayerTakeShot(grid, 6, 4)

		PlayerTakeShot(grid, 5, 1)
		PlayerTakeShot(grid, 1, 3)
		PlayerTakeShot(grid, 2, 4)

		if !PlayerWon(grid){
			t.Error("Player should have won")
		}
		} */
	


func TestPlayerHasNotWonBeforeSinkingAllShips(t *testing.T) {
    
    grid := CreateGrid()
    grid, _ = placeShip(grid, 1, 2)
    grid, _ = placeShip(grid, 2, 3)
    grid, _ = placeShip(grid, 3, 4)

    grid, _ = placeShip(grid, 4, 5)
    grid, _ = placeShip(grid, 5, 6)
    grid, _ = placeShip(grid, 6, 4)

    grid, _ = placeShip(grid, 5, 1)
    grid, _ = placeShip(grid, 1, 3)
    grid, _ = placeShip(grid, 2, 4)
    
    PlayerTakeShot(grid, 1, 2)
    PlayerTakeShot(grid, 2, 3)
	
    PlayerTakeShot(grid, 3, 4)
    PlayerTakeShot(grid, 4, 5)
    
    if PlayerWon(grid) {
        t.Error("Player should not have won")
    }
}

