package game

//API
//
//Usage:
// newGameState = player.makeMove(game)
//

type Ticktacktoe struct {
	//playing field
	//field[x] = 0 : free; = 1: marked by player 1; =  2: marked by player 2
	//treated like
	//[0,1,2]
	//[3,4,5]
	//[6,7,8]
	field       [9]int
	moveCounter int
}

// allows a player to make a move by selecting a field on the board.
func (game Ticktacktoe) MakeMove(fieldNumber int) Ticktacktoe {
	// we assume the players do their moves successively
	game.field[fieldNumber] = (game.moveCounter % 2) + 1
	game.moveCounter++
	return game
}

// Checks the current state of the game.
// returns
// 0 : game not finished;
// 1 : player 1 won;
// 2 : player 2 won;
// 3 : no winner but game finished (tie)
func (game Ticktacktoe) CheckState() int {
	var player = 1

	// check two times
	//once with player = 1 and once with player = 2
	for i := 0; i < 2; i++ {

		//check horizontally
		if game.checkRow(0, player) {
			return player
		} else if game.checkRow(3, player) {
			return player
		} else if game.checkRow(6, player) {
			return player
		}

		//check vertically
		if game.checkColumn(0, player) {
			return player
		} else if game.checkColumn(1, player) {
			return player
		} else if game.checkColumn(2, player) {
			return player
		}

		//check cross
		if game.checkCrossLeftRight(player) {
			return player
		} else if game.checkCrossRightLeft(player) {
			return player
		}

		player = 2
	}

	// are there still free fields ?
	if contains(game.field, 0) {
		return 0
	} else {
		// no fields left and no winner found. Its a tie !
		return 3
	}
}

//Checks if the entire row is marked with one value
func (game Ticktacktoe) checkRow(rowStartField int, valueToCheck int) bool {
	return game.field[rowStartField] == valueToCheck &&
		game.field[rowStartField+1] == valueToCheck &&
		game.field[rowStartField+2] == valueToCheck
}

func (game Ticktacktoe) checkColumn(columnStartField int, valueToCheck int) bool {
	return game.field[columnStartField] == valueToCheck &&
		game.field[columnStartField+3] == valueToCheck &&
		game.field[columnStartField+6] == valueToCheck
}

func (game Ticktacktoe) checkCrossLeftRight(valueToCheck int) bool {
	return game.field[0] == valueToCheck &&
		game.field[4] == valueToCheck &&
		game.field[8] == valueToCheck
}

func (game Ticktacktoe) checkCrossRightLeft(valueToCheck int) bool {
	return game.field[2] == valueToCheck &&
		game.field[4] == valueToCheck &&
		game.field[6] == valueToCheck
}

// checks whether a value is present in an array
func contains(array [9]int, value int) bool {
	for _, i := range array {
		if array[i] == value {
			return true
		}
	}
	return false
}
