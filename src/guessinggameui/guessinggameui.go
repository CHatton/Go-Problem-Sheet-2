package guessinggameui

type GuessingGameUI struct {
	Text, Guess, DisplayMessage string
}

// simple "constructor" function to create message structs with
// a provided text value.
func New(text, guess, displayMessage string) *GuessingGameUI {
	ui := new(GuessingGameUI)
	ui.Text = text
	ui.Guess = guess
	ui.DisplayMessage = displayMessage
	return ui
}
