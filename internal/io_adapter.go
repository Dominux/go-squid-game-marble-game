package game

// Adapter to perform i/o operations with the game
type IOAdapter interface {
	get(promptMsg string) (string, error)
	set(promptMsg string, value interface{}) error
}
