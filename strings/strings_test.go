package strings

import "testing"

func Test_strings_Lines(t *testing.T) {
	lines := Lines("line 1\nline 2\nline 3")
	if len(lines) != 3 {
		t.Errorf("expected %d lines, got %d", 3, len(lines))
	}
}

func Test_strings_TruncateAt(t *testing.T) {
	t.Run("given an empty string", func(t *testing.T) {
		if value := TruncateAt("", 10); value != "" {
			t.Error("value should be empty")
		}
	})

	t.Run("given a small string", func(t *testing.T) {
		if value := TruncateAt("hi", 10); value != "hi" {
			t.Error("value should be untouched")
		}
	})

	t.Run("given a big string", func(t *testing.T) {
		value := TruncateAt("Hochgeschwindigkeitszug", 10)
		expected := "Hochgeschw"
		if value != expected {
			t.Errorf("expected %q, got %q", expected, value)
		}
	})
}

func Test_strings_LeftPad(t *testing.T) {
	t.Run("given an empty string", func(t *testing.T) {
		value := LeftPad("", 3, '0')
		expected := "000"
		if value != expected {
			t.Errorf("expected %q, got %q", expected, value)
		}
	})

	t.Run("given a string with equal size of padding", func(t *testing.T) {
		value := LeftPad("hi", 2, 'A')
		expected := "hi"
		if value != expected {
			t.Errorf("expected %q, got %q", expected, value)
		}
	})

	t.Run("given a string that needs to be padded", func(t *testing.T) {
		value := LeftPad("99", 5, '1')
		expected := "11199"
		if value != expected {
			t.Errorf("expected %q, got %q", expected, value)
		}
	})
}

func Test_strings_RightPad(t *testing.T) {
	t.Run("given an empty string", func(t *testing.T) {
		value := RightPad("", 3, '0')
		expected := "000"
		if value != expected {
			t.Errorf("expected %q, got %q", expected, value)
		}
	})

	t.Run("given a string with equal size of padding", func(t *testing.T) {
		value := RightPad("hi", 2, 'A')
		expected := "hi"
		if value != expected {
			t.Errorf("expected %q, got %q", expected, value)
		}
	})

	t.Run("given a string that needs to be padded", func(t *testing.T) {
		value := RightPad("99", 5, '1')
		expected := "99111"
		if value != expected {
			t.Errorf("expected %q, got %q", expected, value)
		}
	})
}
