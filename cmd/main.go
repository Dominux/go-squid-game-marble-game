package main

import game "github.com/Dominux/go-squid-game-marble-game/internal"

func main() {
	g := game.NewGame()

	for {
		g.Player1.MakeMove(g)
		g.Player2.MakeMove(g)
		if g.EndMove() {
			break
		}
	}
}
