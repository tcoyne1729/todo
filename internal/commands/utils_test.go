package commands_test

import (
	"maps"
	"testing"

	"github.com/tcoyne1729/todo/internal/commands"
)

func TestPointString(t *testing.T) {
	t.Run("padding", func(t *testing.T) {
		got := commands.PointString("abc", "abc")
		want := "-->"
		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})

	t.Run("shorten index example 1", func(t *testing.T) {
		ids := []string{"abc", "acd", "qwe"}
		expected := make(map[string]string)
		expected["abc"] = "ab"
		expected["acd"] = "ac"
		expected["qwe"] = "qw"
		out, err := commands.ShortIds(ids)
		if err != nil {
			t.Fatalf("error processing: %v", err)
		}
		if !maps.Equal(expected, out.LongToShort) {
			t.Errorf("expected: %v\ngot: %v", expected, out)
		}
	})

	t.Run("shorten index example 2", func(t *testing.T) {
		ids := []string{"abc", "vvd", "qwe"}
		expected := make(map[string]string)
		expected["abc"] = "a"
		expected["vvd"] = "v"
		expected["qwe"] = "q"
		out, err := commands.ShortIds(ids)
		if err != nil {
			t.Fatalf("error processing: %v", err)
		}
		if !maps.Equal(expected, out.LongToShort) {
			t.Errorf("expected: %v\ngot: %v", expected, out)
		}
	})

	t.Run("shorten index example 3", func(t *testing.T) {
		ids := []string{"abc", "abd", "abe"}
		expected := make(map[string]string)
		expected["abc"] = "abc"
		expected["abd"] = "abd"
		expected["abe"] = "abe"
		out, err := commands.ShortIds(ids)
		if err != nil {
			t.Fatalf("error processing: %v", err)
		}
		if !maps.Equal(expected, out.LongToShort) {
			t.Errorf("expected: %v\ngot: %v", expected, out)
		}
	})

	t.Run("shorten index with spaces", func(t *testing.T) {
		ids := []string{"", "abc"}
		expected := make(map[string]string)
		expected["abc"] = "a"
		out, err := commands.ShortIds(ids)
		if err != nil {
			t.Fatalf("error processing: %v", err)
		}
		if !maps.Equal(expected, out.LongToShort) {
			t.Errorf("expected: %v\ngot: %v", expected, out)
		}
	})
	t.Run("shorten index empty", func(t *testing.T) {
		ids := []string{}
		expected := make(map[string]string)
		out, err := commands.ShortIds(ids)
		if err != nil {
			t.Fatalf("error processing: %v", err)
		}
		if !maps.Equal(expected, out.LongToShort) {
			t.Errorf("expected: %v\ngot: %v", expected, out)
		}
	})
}
