package Game

import ("testing"
"battleships/Grid"
"battleships/Player"
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