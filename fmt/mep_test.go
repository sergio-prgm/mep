package mep

import (
	"testing"
)

// TestInk calls mep.Ink with content, color and mode
// checking for a valid return type
func TestInkCorrect(t *testing.T) {
	content := "Testing"
	color := Red
	mode := Italic
	msg := Ink(content, color, mode)
	want := "\x1b[3;31mTesting\x1b[0m"
	if msg != want {
		t.Fatalf(`Ink(%s, %s, %d) = %q, want match for %#q, nil`, content, color, mode, msg, want)
	}
}
