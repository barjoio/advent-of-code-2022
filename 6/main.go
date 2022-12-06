package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	// markerLength := 4 // (part one)
	markerLength := 14 // (part two)

	ans := 0

streamLoop:
	for i := range input {
		if i > markerLength-1 {
			seg := []rune(input)[i-markerLength : i]
			for j, c := range seg {
				if strings.ContainsRune(string(seg[j+1:]), c) {
					continue streamLoop
				}
			}
			ans = i
			break
		}
	}

	fmt.Println(ans)
}
