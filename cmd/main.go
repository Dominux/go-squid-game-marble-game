package main

import (
	game "github.com/Dominux/go-squid-game-marble-game/internal"
	gi "github.com/Dominux/go-squid-game-marble-game/internal/game_interactors"
)

func main() {
	gi := gi.NewTerminalGameInteractor()
	g := game.NewGame(gi)

	for {
		g.Player1.MakeMove(g)
		g.Player2.MakeMove(g)
		isGameEnded := g.EndRound()

		if isGameEnded {
			break
		}
	}
}
