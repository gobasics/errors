package errors

type Message string

func (m Message) Error() string {
	return string(m)
}

func New(m string) error {
	return Message(m)
}
