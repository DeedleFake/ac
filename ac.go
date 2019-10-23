package ac

import (
	"fmt"
	"io"
	"strings"
)

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
		fmt.Fprint(&buf, p)
	}

	buf.WriteRune(c)
	return buf.String()
}

type Black struct {
	V interface{}
}

func (v Black) Format(f fmt.State, c rune) {
	io.WriteString(f, BlackCode)
	fmt.Fprintf(f, verb(f, c), v.V)
	io.WriteString(f, ResetCode)
}

type Red struct {
	V interface{}
}

func (v Red) Format(f fmt.State, c rune) {
	io.WriteString(f, RedCode)
	fmt.Fprintf(f, verb(f, c), v.V)
	io.WriteString(f, ResetCode)
}

type Green struct {
	V interface{}
}

func (v Green) Format(f fmt.State, c rune) {
	io.WriteString(f, GreenCode)
	fmt.Fprintf(f, verb(f, c), v.V)
	io.WriteString(f, ResetCode)
}

type Yellow struct {
	V interface{}
}

func (v Yellow) Format(f fmt.State, c rune) {
	io.WriteString(f, YellowCode)
	fmt.Fprintf(f, verb(f, c), v.V)
	io.WriteString(f, ResetCode)
}

type Blue struct {
	V interface{}
}

func (v Blue) Format(f fmt.State, c rune) {
	io.WriteString(f, BlueCode)
	fmt.Fprintf(f, verb(f, c), v.V)
	io.WriteString(f, ResetCode)
}

type Magenta struct {
	V interface{}
}

func (v Magenta) Format(f fmt.State, c rune) {
	io.WriteString(f, MagentaCode)
	fmt.Fprintf(f, verb(f, c), v.V)
	io.WriteString(f, ResetCode)
}

type Cyan struct {
	V interface{}
}

func (v Cyan) Format(f fmt.State, c rune) {
	io.WriteString(f, CyanCode)
	fmt.Fprintf(f, verb(f, c), v.V)
	io.WriteString(f, ResetCode)
}

type White struct {
	V interface{}
}

func (v White) Format(f fmt.State, c rune) {
	io.WriteString(f, WhiteCode)
	fmt.Fprintf(f, verb(f, c), v.V)
	io.WriteString(f, ResetCode)
}

type BrightBlack struct {
	V interface{}
}

func (v BrightBlack) Format(f fmt.State, c rune) {
	io.WriteString(f, BrightBlackCode)
	fmt.Fprintf(f, verb(f, c), v.V)
	io.WriteString(f, ResetCode)
}

type BrightRed struct {
	V interface{}
}

func (v BrightRed) Format(f fmt.State, c rune) {
	io.WriteString(f, BrightRedCode)
	fmt.Fprintf(f, verb(f, c), v.V)
	io.WriteString(f, ResetCode)
}

type BrightGreen struct {
	V interface{}
}

func (v BrightGreen) Format(f fmt.State, c rune) {
	io.WriteString(f, BrightGreenCode)
	fmt.Fprintf(f, verb(f, c), v.V)
	io.WriteString(f, ResetCode)
}

type BrightYellow struct {
	V interface{}
}

func (v BrightYellow) Format(f fmt.State, c rune) {
	io.WriteString(f, BrightYellowCode)
	fmt.Fprintf(f, verb(f, c), v.V)
	io.WriteString(f, ResetCode)
}

type BrightBlue struct {
	V interface{}
}

func (v BrightBlue) Format(f fmt.State, c rune) {
	io.WriteString(f, BrightBlueCode)
	fmt.Fprintf(f, verb(f, c), v.V)
	io.WriteString(f, ResetCode)
}

type BrightMagenta struct {
	V interface{}
}

func (v BrightMagenta) Format(f fmt.State, c rune) {
	io.WriteString(f, BrightMagentaCode)
	fmt.Fprintf(f, verb(f, c), v.V)
	io.WriteString(f, ResetCode)
}

type BrightCyan struct {
	V interface{}
}

func (v BrightCyan) Format(f fmt.State, c rune) {
	io.WriteString(f, BrightCyanCode)
	fmt.Fprintf(f, verb(f, c), v.V)
	io.WriteString(f, ResetCode)
}

type BrightWhite struct {
	V interface{}
}

func (v BrightWhite) Format(f fmt.State, c rune) {
	io.WriteString(f, BrightWhiteCode)
	fmt.Fprintf(f, verb(f, c), v.V)
	io.WriteString(f, ResetCode)
}