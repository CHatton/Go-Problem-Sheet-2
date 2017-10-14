package message

type Message struct {
	Text string
}

// simple "constructor" function to create message structs with
// a provided text value.
func New(text string) *Message {
	msg := new(Message)
	msg.Text = text
	return msg
}
