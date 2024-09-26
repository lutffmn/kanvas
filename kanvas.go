package kanvas

import (
	"fmt"
	"io"
	"strings"
)

// Sequences is a struct to hold message and options
type Sequences struct {
	writer    io.Writer
	sequences []string
}

// Kanvas is a constructor function.
// It returns the exact sequence that passed as argument.
// So it can directly call Invoke method afterward.
func Kanvas(s Sequences) *Sequences {
	return &s
}

// Invoke method is use to process the message through all the options
func (s *Sequences) Invoke() {
	result := strings.Join(s.sequences, ",")
	fmt.Fprintf(s.writer, result)
}

func ESC() string {
	return fmt.Sprintf("%s", "\x1B")
}
