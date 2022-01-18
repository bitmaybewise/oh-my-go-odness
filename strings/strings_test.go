package strings

import "testing"

func Test_strings_Lines(t *testing.T) {
	lines := Lines("line 1\nline 2\nline 3")
	if len(lines) != 3 {
		t.Errorf("expected %d lines, got %d", 3, len(lines))
	}
}
