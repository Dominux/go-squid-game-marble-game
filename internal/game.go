package game

type MarblesAmount uint8

const INIT_MARBLES_AMOUNT MarblesAmount = 100

type Player struct {
	Name          string
	MarblesAmount MarblesAmount
	role          Role
}

func (p *Player) MakeMove(ioa *IOAdapter, g *Game) {
	switch p.role {
	case Riddler:
		p.makeMoveAsRiddler(ioa, g)
	case Guesser:
		p.makeMoveAsGuesser(ioa, g)
	}
}

func (p *Player) makeMoveAsRiddler(ioa *IOAdapter, g *Game) {

}
func (p *Player) makeMoveAsGuesser(ioa *IOAdapter, g *Game) {}

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

func NewGame(player1_name string, player2_name string) Game {
	player1 := Player{Name: player1_name, MarblesAmount: INIT_MARBLES_AMOUNT}
	player2 := Player{Name: player2_name, MarblesAmount: INIT_MARBLES_AMOUNT}
	return Game{
		Player1: player1,
		Player2: player2,
	}
}
