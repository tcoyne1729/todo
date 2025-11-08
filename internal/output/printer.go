package output

import (
	"fmt"
	"io"
	"os"
)

type Style struct {
	code string
}

func Color(code string) Style { return Style{code: code} }
func Bold() Style             { return Style{"1"} }
func Red() Style              { return Style{"31"} }
func Green() Style            { return Style{"32"} }
func Blue() Style             { return Style{"34"} }

func Combine(styles ...Style) Style {
	concatStyleCode := ""
	for i, s := range styles {
		if i > 0 {
			concatStyleCode += ";"
		}
		concatStyleCode += s.code
	}
	return Style{code: concatStyleCode}
}
func (s Style) Wrap(text string) string {
	return fmt.Sprintf("\033[%sm%s\033[0m", s.code, text)
}

// define our custom printer

type Printer struct {
	enabled bool
	writer  io.Writer
}

func NewPrinter(enabled bool) Printer {
	return Printer{
		enabled: enabled,
		writer:  os.Stdout,
	}
}

func (p Printer) Printf(style Style, format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	if p.enabled {
		msg = style.Wrap(msg)
	}
	// fmt.Fprintf(p.writer, msg)
	fmt.Printf(format, a...)

}
func (p Printer) Println(style Style, a ...any) {
	msg := fmt.Sprintln(a...)
	if p.enabled {
		msg = style.Wrap(msg)
	}
	fmt.Fprintln(p.writer, msg)
}
