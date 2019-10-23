package ac_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/DeedleFake/ac"
)

func TestColors(t *testing.T) {
	tests := []struct {
		name string
		f    string
		in   interface{}
		out  string
	}{
		{
			name: "Black",
			f:    "%s",
			in:   ac.Black("test"),
			out:  ac.BlackCode + "test" + ac.ResetCode,
		},
		{
			name: "Red",
			f:    "%s",
			in:   ac.Red("test"),
			out:  ac.RedCode + "test" + ac.ResetCode,
		},
		{
			name: "Green",
			f:    "%s",
			in:   ac.Green("test"),
			out:  ac.GreenCode + "test" + ac.ResetCode,
		},
		{
			name: "Yellow",
			f:    "%s",
			in:   ac.Yellow("test"),
			out:  ac.YellowCode + "test" + ac.ResetCode,
		},
		{
			name: "Blue",
			f:    "%s",
			in:   ac.Blue("test"),
			out:  ac.BlueCode + "test" + ac.ResetCode,
		},
		{
			name: "Magenta",
			f:    "%s",
			in:   ac.Magenta("test"),
			out:  ac.MagentaCode + "test" + ac.ResetCode,
		},
		{
			name: "Cyan",
			f:    "%s",
			in:   ac.Cyan("test"),
			out:  ac.CyanCode + "test" + ac.ResetCode,
		},
		{
			name: "White",
			f:    "%s",
			in:   ac.White("test"),
			out:  ac.WhiteCode + "test" + ac.ResetCode,
		},
		{
			name: "BrightBlack",
			f:    "%s",
			in:   ac.BrightBlack("test"),
			out:  ac.BrightBlackCode + "test" + ac.ResetCode,
		},
		{
			name: "BrightRed",
			f:    "%s",
			in:   ac.BrightRed("test"),
			out:  ac.BrightRedCode + "test" + ac.ResetCode,
		},
		{
			name: "BrightGreen",
			f:    "%s",
			in:   ac.BrightGreen("test"),
			out:  ac.BrightGreenCode + "test" + ac.ResetCode,
		},
		{
			name: "BrightYellow",
			f:    "%s",
			in:   ac.BrightYellow("test"),
			out:  ac.BrightYellowCode + "test" + ac.ResetCode,
		},
		{
			name: "BrightBlue",
			f:    "%s",
			in:   ac.BrightBlue("test"),
			out:  ac.BrightBlueCode + "test" + ac.ResetCode,
		},
		{
			name: "BrightMagenta",
			f:    "%s",
			in:   ac.BrightMagenta("test"),
			out:  ac.BrightMagentaCode + "test" + ac.ResetCode,
		},
		{
			name: "BrightCyan",
			f:    "%s",
			in:   ac.BrightCyan("test"),
			out:  ac.BrightCyanCode + "test" + ac.ResetCode,
		},
		{
			name: "BrightWhite",
			f:    "%s",
			in:   ac.BrightWhite("test"),
			out:  ac.BrightWhiteCode + "test" + ac.ResetCode,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			var buf strings.Builder
			fmt.Fprintf(&buf, test.f, test.in)
			out := buf.String()

			if out != test.out {
				t.Errorf("expected %q, got %q", test.out, out)
			}
		})
	}
}
