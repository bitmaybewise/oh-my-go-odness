package strings

import "strings"

func Lines(s string) []string {
	return strings.Split(string(s), "\n")
}
