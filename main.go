package main

import (
	"fmt"
	"math"
)

func _print_board(game *Game) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Print(game.Board[i][j])
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func bot_vs_bot() {
	var game = newGame()
	var cnt = 0
	for !game.isGameEnded() {
		game = BOT.findBestMove(game, cnt, (cnt+1)%2)
		fmt.Println(math.Round(BOT.gameTemp(&game)))
		_print_board(&game)
		cnt++
		cnt %= 2
	}
	fmt.Println(game.whoWin())
}

func main() {
	_init()
	bot_vs_bot()
}
