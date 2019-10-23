// Package ac provides simple functionality for producing ANSI
// color-code output.
//
// This package is designed to be used in conjunction with the
// standard library's fmt package. For example,
//
//    fmt.Printf("%s %v\n", ac.Red("Error:"), err)
//
// will produce a red "Error:" prefix in the termnial. The other types
// in this package work much the same way. Any formatting verb may be
// used with the types in this package, and it will produce the same
// output as if the type stored in the color type was used directly,
// but colored. Note that this colors the entire output, so if, for
// example, "%q" is the verb used, the quotation marks will also be
// colored.
package ac

import (
	"fmt"
	"io"
	"strings"
)

// Disabled disables all coloration. This allows a client to, for
// example, turn off coloring easily when printing to a non-terminal.
var Disabled bool

// ANSI color codes, exported for convienence. These are rarely
// necessary to use directly, but might be for some custom formatting
// cases.
const (
	BlackCode = "\u001b[3" + string('0'+iota) + "m"
	RedCode
	GreenCode
	YellowCode
	BlueCode
	MagentaCode
	CyanCode
	WhiteCode

	BrightBlackCode = "\u001b[3" + string('0'+iota-8) + ";1m"
	BrightRedCode
	BrightGreenCode
	BrightYellowCode
	BrightBlueCode
	BrightMagentaCode
	BrightCyanCode
	BrightWhiteCode

	ResetCode = "\u001b[0m"
)

var flags = []rune{'+', '_', '#', ' ', '0'}

func verb(f fmt.State, c rune) string {
	var buf strings.Builder
	buf.WriteRune('%')

	for _, flag := range flags {
		if f.Flag(int(flag)) {
			buf.WriteRune(flag)
		}
	}

	if w, ok := f.Width(); ok {
		fmt.Fprint(&buf, w)
	}

	if p, ok := f.Precision(); ok {
		buf.WriteRune('.')
		fmt.Fprint(&buf, p)
	}

	buf.WriteRune(c)
	return buf.String()
}

type colorizer struct {
	v    interface{}
	code string
}

func (col colorizer) Format(f fmt.State, c rune) {
	io.WriteString(f, col.code)
	fmt.Fprintf(f, verb(f, c), col.v)
	io.WriteString(f, ResetCode)
}

func Black(v interface{}) interface{} {
	if Disabled {
		return v
	}

	return colorizer{
		v:    v,
		code: BlackCode,
	}
}

func Red(v interface{}) interface{} {
	if Disabled {
		return v
	}

	return colorizer{
		v:    v,
		code: RedCode,
	}
}

func Green(v interface{}) interface{} {
	if Disabled {
		return v
	}

	return colorizer{
		v:    v,
		code: GreenCode,
	}
}

func Yellow(v interface{}) interface{} {
	if Disabled {
		return v
	}

	return colorizer{
		v:    v,
		code: YellowCode,
	}
}

func Blue(v interface{}) interface{} {
	if Disabled {
		return v
	}

	return colorizer{
		v:    v,
		code: BlueCode,
	}
}

func Magenta(v interface{}) interface{} {
	if Disabled {
		return v
	}

	return colorizer{
		v:    v,
		code: MagentaCode,
	}
}

func Cyan(v interface{}) interface{} {
	if Disabled {
		return v
	}

	return colorizer{
		v:    v,
		code: CyanCode,
	}
}

func White(v interface{}) interface{} {
	if Disabled {
		return v
	}

	return colorizer{
		v:    v,
		code: WhiteCode,
	}
}

func BrightBlack(v interface{}) interface{} {
	if Disabled {
		return v
	}

	return colorizer{
		v:    v,
		code: BrightBlackCode,
	}
}

func BrightRed(v interface{}) interface{} {
	if Disabled {
		return v
	}

	return colorizer{
		v:    v,
		code: BrightRedCode,
	}
}

func BrightGreen(v interface{}) interface{} {
	if Disabled {
		return v
	}

	return colorizer{
		v:    v,
		code: BrightGreenCode,
	}
}

func BrightYellow(v interface{}) interface{} {
	if Disabled {
		return v
	}

	return colorizer{
		v:    v,
		code: BrightYellowCode,
	}
}

func BrightBlue(v interface{}) interface{} {
	if Disabled {
		return v
	}

	return colorizer{
		v:    v,
		code: BrightBlueCode,
	}
}

func BrightMagenta(v interface{}) interface{} {
	if Disabled {
		return v
	}

	return colorizer{
		v:    v,
		code: BrightMagentaCode,
	}
}

func BrightCyan(v interface{}) interface{} {
	if Disabled {
		return v
	}

	return colorizer{
		v:    v,
		code: BrightCyanCode,
	}
}

func BrightWhite(v interface{}) interface{} {
	if Disabled {
		return v
	}

	return colorizer{
		v:    v,
		code: BrightWhiteCode,
	}
}
