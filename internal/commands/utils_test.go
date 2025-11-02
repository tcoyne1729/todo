package commands_test

import (
	"testing"

	"github.com/tcoyne1729/todo/internal/commands"
)

func TestPointString(t *testing.T) {
	got := commands.PointString("abc", "abc")
	want := "-->"
	if got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}
}
