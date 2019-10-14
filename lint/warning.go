package lint

import "fmt"

// Warning ...
type Warning struct {
	Messages []string
	Line     int
	Column   int
}

func (w Warning) Error() (s string) {
	for _, msg := range w.Messages {
		s += fmt.Sprintf("%d:%d: %s\n", w.Line, w.Column, msg)
	}
	return
}
