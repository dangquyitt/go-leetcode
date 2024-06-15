package main

import (
	"fmt"
	"strings"
)

func restoreString(s string, indices []int) string {
	result := make([]byte, len(indices))
	for i, v := range indices {
		result[v] = s[i]
	}
	return string(result)
}

func main() {
	fmt.Println(restoreString("codeleet", []int{4, 5, 6, 7, 0, 2, 1, 3}))
	fmt.Println(restoreString("abc", []int{0, 1, 2}))
	fmt.Println(strings.ReplaceAll("1528-shuffle-string.go", "-", " "))
}
