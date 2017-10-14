package message

type Message struct {
	Text, Guess, DisplayMessage string
}

// simple "constructor" function to create message structs with
// a provided text value.
func New(text, guess, displayMessage string) *Message {
	msg := new(Message)
	msg.Text = text
	msg.Guess = guess
	msg.DisplayMessage = displayMessage
	return msg
}
