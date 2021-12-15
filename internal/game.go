package game

import (
	"fmt"
	"strings"

	gi "github.com/Dominux/go-squid-game-marble-game/internal/game_interactors"
)

type MarblesAmount uint8

const INIT_MARBLES_AMOUNT MarblesAmount = 100

type Player struct {
	Name          string
	MarblesAmount MarblesAmount
	role          Role
}

func (p *Player) MakeMove(g *Game) {
	g.gameInteractor.Say(fmt.Sprintf("\n%s's move:", p.Name))

	switch p.role {
	case Riddler:
		p.makeMoveAsRiddler(g)
	case Guesser:
		p.makeMoveAsGuesser(g)
	}
}

func (p *Player) makeMoveAsRiddler(g *Game) {
	// Getting amount of marbles the riddler is gonna put
	amount, err := g.gameInteractor.GetNumber("Choose even or odd amount of your marbles:")
	for err != nil || amount < 0 || amount > int(p.MarblesAmount) {
		amount, err = g.gameInteractor.GetNumber("Something went wrong, try again:")
	}

	g.parety = Parety(amount % 2)
}

func (p *Player) makeMoveAsGuesser(g *Game) {
	// Getting guessed parety
loopA:
	for {
		text, err := g.gameInteractor.GetString("Guess if the riddler chosen \"even\" or \"odd\" amount of his marbles:")
		if err != nil {
			continue
		}

		switch strings.ToLower(text) {
		case "even":
			g.guessedParety = Even
			break loopA
		case "odd":
			g.guessedParety = Odd
			break loopA
		default:
			fmt.Println("Ya stupid dumba, write \"even\" or \"odd\"!")
		}
	}

	// Getting bet
	for {
		bet, err := g.gameInteractor.GetNumber("Bet amount of marbles that isn't bigger than yours nor your opponent:")
		if err != nil {
			fmt.Println(err)
		} else if bet < 0 || bet > int(p.MarblesAmount) {
			fmt.Println("Ya stupid dumba, write proper amount of marbles you have!")
		} else {
			g.bet = MarblesAmount(bet)
			break
		}
	}
}

type Role uint8

const (
	Riddler Role = iota
	Guesser
)

type Parety uint8

const (
	Even Parety = iota
	Odd
)

type Game struct {
	Player1        Player
	Player2        Player
	gameInteractor gi.GameInteractor
	parety         Parety
	guessedParety  Parety
	bet            MarblesAmount
}

func NewGame(gi gi.GameInteractor) *Game {
	player1_name, _ := gi.GetString("Write ur dickin' name, looser:")
	player2_name, _ := gi.GetString("Write ur dickin' name, looser:")

	player1 := Player{Name: player1_name, MarblesAmount: INIT_MARBLES_AMOUNT, role: Riddler}
	player2 := Player{Name: player2_name, MarblesAmount: INIT_MARBLES_AMOUNT, role: Guesser}

	return &Game{
		Player1:        player1,
		Player2:        player2,
		gameInteractor: gi,
	}
}

func (g *Game) EndRound() bool {
	// Validation
	if g.bet == 0 {
		panic("bet is zero")
	}

	// Getting guesser and riddler and chaging their roles
	var riddler *Player
	var guesser *Player
	for _, p := range []*Player{&g.Player1, &g.Player2} {
		switch p.role {
		case Riddler:
			riddler = p
			p.role = Guesser
		case Guesser:
			guesser = p
			p.role = Riddler
		}
	}

	var isGameEnded bool
	if g.parety == g.guessedParety {
		// Guesser won the round
		isGameEnded = g.transferBet(guesser, riddler, g.bet)
	} else {
		// Riddler won the round
		isGameEnded = g.transferBet(riddler, guesser, g.bet)
	}

	if !isGameEnded {
		g.SayGameStatus()
	}

	return isGameEnded
}

// Transfer bet from the looser to the winner
// and check if looser is lack of marbles
func (g *Game) transferBet(winner *Player, looser *Player, bet MarblesAmount) bool {
	looser.MarblesAmount -= bet
	winner.MarblesAmount += bet

	isGameEnded := looser.MarblesAmount == 0
	if isGameEnded {
		g.gameInteractor.Say(fmt.Sprintf("The game is over, player %v won!", winner.Name))
	} else {
		g.gameInteractor.Say(fmt.Sprintf("%v won the round", winner.Name))
	}

	return isGameEnded
}

func (g *Game) SayGameStatus() {
	player1_status := fmt.Sprintf("Player %s: %d", g.Player1.Name, g.Player1.MarblesAmount)
	player2_status := fmt.Sprintf("Player %s: %d", g.Player2.Name, g.Player2.MarblesAmount)
	g.gameInteractor.Say(fmt.Sprintf("\n%s\n%s\n\n", player1_status, player2_status))
}
