package errors

import "strings"

type Messages []string

func (m Messages) Error() string {
	return strings.Join(m, "; ")
}

type Stream chan error

func (s Stream) Error() error {
	var m Messages
	for i := 0; i < cap(s); i++ {
		if err := <-s; err != nil {
			m = append(m, err.Error())
		}
	}

	if len(m) == 0 {
		return nil
	}

	return m
}

func NewStream(capacity int) Stream {
	return make(Stream, capacity)
}
