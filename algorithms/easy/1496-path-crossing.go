package main

import (
	"fmt"
	"strconv"
)

// git commit -m "1496 path crossing"
func isPathCrossing(path string) bool {
	x := 0
	y := 0

	visitedNode := make(map[string]bool)
	visitedNode["00"] = true

	for _, p := range path {
		switch p {
		case 'N':
			y++
		case 'S':
			y--
		case 'E':
			x++
		case 'W':
			x--
		}
		key := strconv.Itoa(x) + strconv.Itoa(y)
		_, isVisited := visitedNode[key]
		if isVisited {
			return true
		} else {
			visitedNode[key] = true
		}
	}
	return false
}

func main() {
	fmt.Println(isPathCrossing("NES"))
	fmt.Println(isPathCrossing("NESWW"))
}
