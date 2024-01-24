package Game

import ("testing"
"battleships/Grid"
"battleships/Player1"
"battleships/Player2")

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
		PlayerTwogrid:= table.CreateGrid()
		
		PlayerTwogrid = Player2.PlayerTwoPlaceShips(PlayerTwogrid, 1, 2)
		PlayerTwogrid = Player2.PlayerTwoPlaceShips(PlayerTwogrid, 2, 3)
		PlayerTwogrid  = Player2.PlayerTwoPlaceShips(PlayerTwogrid, 3, 4)

		PlayerTwogrid  = Player2.PlayerTwoPlaceShips(PlayerTwogrid, 4, 5)
		PlayerTwogrid = Player2.PlayerTwoPlaceShips(PlayerTwogrid, 5, 6)
		PlayerTwogrid = Player2.PlayerTwoPlaceShips(PlayerTwogrid, 6, 4)

		PlayerTwogrid = Player2.PlayerTwoPlaceShips(PlayerTwogrid, 5, 1)
		PlayerTwogrid = Player2.PlayerTwoPlaceShips(PlayerTwogrid, 1, 3)
		PlayerTwogrid = Player2.PlayerTwoPlaceShips(PlayerTwogrid, 2, 4)
											
		Player1.PlayerOneTakeShots(PlayerTwogrid, 1, 2)
		Player1.PlayerOneTakeShots(PlayerTwogrid, 2, 3)					
		Player1.PlayerOneTakeShots(PlayerTwogrid, 3, 4)				

		Player1.PlayerOneTakeShots(PlayerTwogrid, 4, 5)
		Player1.PlayerOneTakeShots(PlayerTwogrid, 5, 6)
		Player1.PlayerOneTakeShots(PlayerTwogrid, 6, 4)

		Player1.PlayerOneTakeShots(PlayerTwogrid, 5, 1)
		Player1.PlayerOneTakeShots(PlayerTwogrid, 1, 3)
		Player1.PlayerOneTakeShots(PlayerTwogrid, 2, 4)

		if PlayerWon(PlayerTwogrid){
			t.Error("Player one should have won")
		}
		} 
	


func TestPlayerOneHasNotWonBeforeSinkingAllShips(t *testing.T) {
    
    PlayerTwogrid := table.CreateGrid()
	PlayerTwogrid = Player2.PlayerTwoPlaceShips(PlayerTwogrid, 1, 2)
	PlayerTwogrid = Player2.PlayerTwoPlaceShips(PlayerTwogrid, 2, 3)
	PlayerTwogrid  = Player2.PlayerTwoPlaceShips(PlayerTwogrid, 3, 4)

	PlayerTwogrid  = Player2.PlayerTwoPlaceShips(PlayerTwogrid, 4, 5)
	PlayerTwogrid = Player2.PlayerTwoPlaceShips(PlayerTwogrid, 5, 6)
	PlayerTwogrid = Player2.PlayerTwoPlaceShips(PlayerTwogrid, 6, 4)

	PlayerTwogrid = Player2.PlayerTwoPlaceShips(PlayerTwogrid, 5, 1)
	PlayerTwogrid = Player2.PlayerTwoPlaceShips(PlayerTwogrid, 1, 3)
	PlayerTwogrid = Player2.PlayerTwoPlaceShips(PlayerTwogrid, 2, 4)

	Player1.PlayerOneTakeShots(PlayerTwogrid, 1, 2)
	Player1.PlayerOneTakeShots(PlayerTwogrid, 2, 3)	

	Player1.PlayerOneTakeShots(PlayerTwogrid, 3, 4)				
	Player1.PlayerOneTakeShots(PlayerTwogrid, 4, 5)
    
    if PlayerWon(PlayerTwogrid) {
        t.Error("Player should not have won")
    }
}