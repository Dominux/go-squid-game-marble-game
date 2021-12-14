package game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MarblesAmount uint8

const INIT_MARBLES_AMOUNT MarblesAmount = 100

type Player struct {
	Name          string
	MarblesAmount MarblesAmount
	role          Role
}

func (p *Player) MakeMove(g *Game) {
	switch p.role {
	case Riddler:
		p.makeMoveAsRiddler(g)
	case Guesser:
		p.makeMoveAsGuesser(g)
	}
}

func (p *Player) makeMoveAsRiddler(g *Game) {
	// Getting amount of marbles the riddler is gonna put
	fmt.Println("Choose even or odd amount of your marbles:")
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		amount, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("Ya stupid dumba, write proper number!")
			continue
		} else if amount < 0 || amount > int(p.MarblesAmount) {
			fmt.Println("Ya stupid dumba, write proper amount of marbles you have!")
			continue
		}

		g.parety = Parety(amount % 2)
		return
	}
}

func (p *Player) makeMoveAsGuesser(g *Game) {
	// Getting guessed parety
	fmt.Println("Guess if the riddler chosen \"even\" or \"odd\" amount of his marbles:")
	reader := bufio.NewReader(os.Stdin)
loopA:
	for {
		text, _ := reader.ReadString('\n')
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
	fmt.Println("Bet amount of marbles that isn't bigger than yours nor your opponent:")
	for {
		text, _ := reader.ReadString('\n')
		bet, err := strconv.Atoi(text)
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
	Player1       Player
	Player2       Player
	parety        Parety
	guessedParety Parety
	bet           MarblesAmount
}

func NewGame() *Game {
	player1_name := getPlayerName()
	player2_name := getPlayerName()

	player1 := Player{Name: player1_name, MarblesAmount: INIT_MARBLES_AMOUNT}
	player2 := Player{Name: player2_name, MarblesAmount: INIT_MARBLES_AMOUNT}

	return &Game{
		Player1: player1,
		Player2: player2,
	}
}

func (g *Game) EndMove() bool {
	// Validation
	if g.parety == 0 || g.guessedParety == 0 || g.bet == 0 {
		panic("LOOOOOL!")
	}

	// Getting guesser and riddler
	var riddler *Player
	var guesser *Player
	for _, p := range []*Player{&g.Player1, &g.Player2} {
		switch p.role {
		case Riddler:
			riddler = p
		case Guesser:
			guesser = p
		}
	}

	if g.parety == g.guessedParety {
		// Guesser won
		return transferBet(guesser, riddler, g.bet)
	} else {
		// Riddler won
		return transferBet(riddler, guesser, g.bet)
	}
}

// Transfer bet from the looser to the winner
// and check if looser is lack of marbles
func transferBet(winner *Player, looser *Player, bet MarblesAmount) bool {
	looser.MarblesAmount -= bet
	winner.MarblesAmount += bet

	fmt.Println("&v won", winner.role)

	return looser.MarblesAmount == 0
}

func getPlayerName() string {
	fmt.Println("Write ur dickin' name, looser:")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	return name
}
