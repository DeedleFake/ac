ac
==

[![GoDoc](http://www.godoc.org/github.com/DeedleFake/ac?status.svg)](http://www.godoc.org/github.com/DeedleFake/ac)

ac provides a Go package that helps with producing ANSI escape sequence compatible colorized terminal output.

Example
-------

```go
type Sum struct {
	Left, Right float64
}

func (s Sum)String() string {
	return fmt.Sprint(s.Left + s.Right)
}

func main() {
	s := Sum{Left: 3.231, Right: 5.561}
	fmt.Printf("%.3v + %.3v = %.3v", ac.Blue(s.Left), ac.Blue(s.Right), ac.Cyan(s))
}
```
