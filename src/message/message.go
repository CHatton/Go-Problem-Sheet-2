package message

type Message struct {
	Text, Guess string
}

// simple "constructor" function to create message structs with
// a provided text value.
func New(text, guess string) *Message {
	msg := new(Message)
	msg.Text = text
	msg.Guess = guess
	return msg
}
