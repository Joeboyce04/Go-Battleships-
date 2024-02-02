package Game

import (
	grid "battleships/Grid"
	"battleships/Player"
	"testing"
)

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


func TestPlayerOneWinsAfterSinkingAllShips(t *testing.T) {
		PlayerTwogrid:= grid.CreateGrid()
		
		PlayerTwogrid, _= Player.PlaceShips(PlayerTwogrid, 1, 2)
		PlayerTwogrid, _ = Player.PlaceShips(PlayerTwogrid, 2, 3)
		PlayerTwogrid, _ = Player.PlaceShips(PlayerTwogrid, 3, 4)

		PlayerTwogrid,_  = Player.PlaceShips(PlayerTwogrid, 4, 5)
		PlayerTwogrid,_ = Player.PlaceShips(PlayerTwogrid, 5, 6)
		PlayerTwogrid,_ = Player.PlaceShips(PlayerTwogrid, 6, 4)

		PlayerTwogrid,_ = Player.PlaceShips(PlayerTwogrid, 5, 1)
		PlayerTwogrid,_ = Player.PlaceShips(PlayerTwogrid, 1, 3)
		PlayerTwogrid, _ = Player.PlaceShips(PlayerTwogrid, 2, 4)
											
		Player.PlayerTakeShots(PlayerTwogrid, 1, 2)
		Player.PlayerTakeShots(PlayerTwogrid, 2, 3)					
		Player.PlayerTakeShots(PlayerTwogrid, 3, 4)				

		Player.PlayerTakeShots(PlayerTwogrid, 4, 5)
		Player.PlayerTakeShots(PlayerTwogrid, 5, 6)
		Player.PlayerTakeShots(PlayerTwogrid, 6, 4)

		Player.PlayerTakeShots(PlayerTwogrid, 5, 1)
		Player.PlayerTakeShots(PlayerTwogrid, 1, 3)
		Player.PlayerTakeShots(PlayerTwogrid, 2, 4)

		if PlayerWon(PlayerTwogrid){
			t.Error("Player one should have won")
		}
		} 
	


func TestPlayerOneHasNotWonBeforeSinkingAllShips(t *testing.T) {
    
    PlayerTwogrid := grid.CreateGrid()
		
	PlayerTwogrid, _= Player.PlaceShips(PlayerTwogrid, 1, 2)
	PlayerTwogrid, _ = Player.PlaceShips(PlayerTwogrid, 2, 3)
	PlayerTwogrid, _ = Player.PlaceShips(PlayerTwogrid, 3, 4)

	PlayerTwogrid,_  = Player.PlaceShips(PlayerTwogrid, 4, 5)
	PlayerTwogrid,_ = Player.PlaceShips(PlayerTwogrid, 5, 6)
	PlayerTwogrid,_ = Player.PlaceShips(PlayerTwogrid, 6, 4)

	PlayerTwogrid,_ = Player.PlaceShips(PlayerTwogrid, 5, 1)
	PlayerTwogrid,_ = Player.PlaceShips(PlayerTwogrid, 1, 3)
	PlayerTwogrid, _ = Player.PlaceShips(PlayerTwogrid, 2, 4)
	
	Player.PlayerTakeShots(PlayerTwogrid, 1, 2)
	Player.PlayerTakeShots(PlayerTwogrid, 2, 3)					
	Player.PlayerTakeShots(PlayerTwogrid, 3, 4)				

	Player.PlayerTakeShots(PlayerTwogrid, 4, 5)
	Player.PlayerTakeShots(PlayerTwogrid, 5, 6)
	Player.PlayerTakeShots(PlayerTwogrid, 6, 4)
    
    if PlayerWon(PlayerTwogrid) {
        t.Error("Player should not have won")
    }
}

func TestTurnChangesWhenPlayerMissesShot(t *testing.T) {
    playerOneTurn := true
    grid := grid.CreateGrid()

    Player.PlayerTakeShots(grid, 3, 3)
    
    playerOneTurn, currentPlayer := passTurnToOtherPlayer(playerOneTurn)
    
    want := "Player Two"
    got := currentPlayer
    if playerOneTurn != false || got != want {
        t.Errorf("Players turn should change when missed shot")
    }
}



func TestTurnChangesWhenPlayerHitsShip(t *testing.T) {
    
    playerOneTurn := true
    grid := grid.CreateGrid()
    grid, _ = Player.PlaceShips(grid, 3, 3)
    
    Player.PlayerTakeShots(grid, 3, 3)
    
    playerOneTurn, currentPlayer := passTurnToOtherPlayer(playerOneTurn)
    
    want := "Player Two"
    got := currentPlayer
    if playerOneTurn != false || got != want {
        t.Errorf("Players turn changed when player hit ship")
    }
}

func TestTurnDoesNotChangeIfShotIsInvalid(t *testing.T) {
    
    playerOneTurn := true
    grid := grid.CreateGrid()
    
    Player.PlayerTakeShots(grid, 7, 7)
    
    currentPlayer := PlayerTurn(playerOneTurn)
    want := "Player One"
    got := currentPlayer
    if playerOneTurn != true || got != want {
        t.Errorf("Players turn does not change after invalid shot")
    }
}

func TestPlayerCannotWinWithInvalidShot(t *testing.T){
	grid:= grid.CreateGrid()

	Player.PlayerTakeShots(grid,9,9)

	if PlayerWon(grid){
		t.Error("Player should not have won with an invalid shot")
	}
}

func TestPlayerCannotWinWithInvalidShipPlacement(t *testing.T){
	grid:= grid.CreateGrid()

	updatedGrid,_ :=PlaceShips(grid, 9,9)

	if PlayerWon(updatedGrid){
		
		t.Error("Player should not have won with an invalid ship placement")
	}
}
