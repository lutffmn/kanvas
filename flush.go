package kanvas

import "fmt"

// Flush is a function to flush entire terminal screen
func Flush() string {
	return fmt.Sprintf("%s%s", ESC(), "2J")
}

// FlushLine is a function to flush entire line
func FlushLine() string {
	return fmt.Sprintf("%s%s", ESC(), "2K")
}
