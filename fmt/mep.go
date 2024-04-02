package mep

import "fmt"

type (
	Color   int64
	Mode    int64
	BgColor int64
)

const (
	Red Color = iota + 31
	Green
	Yellow
	Blue
	Magenta
	Cyan
	Gray
	White Color = 97
)

const (
	Normal    Mode = 0
	Bold      Mode = 1
	Italic    Mode = 3
	Underline Mode = 4
	Strike    Mode = 9
)

const (
	BgRed Color = iota + 40
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgGray
	BgWhite
)

const reset string = "\033[0m"

func (clr Color) String() string {
	switch clr {
	case Red:
		return "31"
	case Green:
		return "32"
	case Yellow:
		return "33"
	case Blue:
		return "34"
	case Magenta:
		return "35"
	case Cyan:
		return "36"
	case Gray:
		return "37"
	case White:
		return "97"
	}
	return "unknown"
}

// Ink brings color and formatting to the content provided.
func Ink(content string, color Color, mode Mode) string {
	return fmt.Sprintf("\033[%d;%dm%s%s", mode, color, content, reset)
}

// Test is a function to test that this works.
// It won't be here when this is published and working.
func Test() {
	fmt.Println("Hello mep!!")
}
