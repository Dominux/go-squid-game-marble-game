package game_interactors

// Adapter to interact with players from within the game
type GameInteractor interface {
	GetString(prompt string) (string, error)
	GetNumber(prompt string) (int, error)
	Say(message string)
}
