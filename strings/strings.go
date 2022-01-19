package strings

import "strings"

func Lines(s string) []string {
	return strings.Split(string(s), "\n")
}

func TruncateAt(value string, index int) string {
	result := strings.Trim(value, " ")
	if len(result) > index {
		return result[:index]
	}

	return result
}

func LeftPad(value string, length int, char rune) string {
	if length <= len(value) {
		return value
	}

	repeat := length - len(value)
	paddedVal := strings.Repeat(string(char), repeat) + value

	return paddedVal
}

func RightPad(value string, length int, char rune) string {
	if length <= len(value) {
		return value
	}

	repeat := length - len(value)
	paddedVal := value + strings.Repeat(string(char), repeat)

	return paddedVal
}
