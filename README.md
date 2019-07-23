TicTacToe

A simple tictactoe game written in go. Provides a simple api for playing. Requires two players or at least input for two players.

Aim off this project is to increase my knowledge of the go language and experience with simple ideas about ai.

Rework:
Changing the api and operation mode. The game will now run like a server and listen for incoming moves. The players identify themselves either as player 1 or 2.

API:
Reset()
GetState() # returns playing field and player number who has his turn
MakeMove(id int, field int) 


Improvements:
Currently a new connection is created for every single message sent. Causes a lot of delay.