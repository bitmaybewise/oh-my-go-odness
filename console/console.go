package console

import "fmt"

const ansiClearScreenChars = "\033[H\033[2J"

func Clear() {
	fmt.Print(ansiClearScreenChars)
}
