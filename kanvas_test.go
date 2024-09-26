package kanvas

import (
	"bytes"
	"os"
	"reflect"
	"testing"
)

func TestKanvas(t *testing.T) {
	right := &Sequences{
		writer:    os.Stdout,
		sequences: []string{"blue", "eyes"},
	}

	left := Kanvas(*right)

	if !reflect.DeepEqual(left, right) {
		t.Errorf("left: %v\tright: %v\n", left, right)
	}
}

// TODO: test Invoke method
func TestInvoke(t *testing.T) {
	right := "blueeyes"

	writer := bytes.NewBuffer(make([]byte, 128))

	left := Kanvas(Sequences{
		writer:    writer,
		sequences: []string{"blue", "eyes"},
	})

	left.Invoke()

	if !reflect.DeepEqual(left, right) {
		t.Errorf("left: %v\tright: %v\n", left, right)
	}
}
