package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	if s == "" {
		return 0
	}

	s = strings.TrimSpace(s)

	stringToArr := strings.Split(s, " ")

	result := stringToArr[len(stringToArr)-1]

	result = strings.TrimSpace(result)

	return len(result)
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
	fmt.Println(lengthOfLastWord("luffy is still joyboy"))
}
