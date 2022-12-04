package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

var shapeScores = map[rune]int{
	'X': 1,
	'Y': 2,
	'Z': 3,
}

func main() {
	rounds := strings.Split(strings.TrimSpace(input), "\n")

	ans := 0
	ans2 := 0

	for _, round := range rounds {
		elfShape := round[0]
		myShape := round[2]

		score := shapeScores[rune(myShape)]

		switch elfShape {
		case 'A':
			switch myShape {
			case 'X':
				score += 3
			case 'Y':
				score += 6
			case 'Z':
				score += 0
			}
		case 'B':
			switch myShape {
			case 'X':
				score += 0
			case 'Y':
				score += 3
			case 'Z':
				score += 6
			}
		case 'C':
			switch myShape {
			case 'X':
				score += 6
			case 'Y':
				score += 0
			case 'Z':
				score += 3
			}
		}

		ans += score
	}

	for _, round := range rounds {
		elfShape := round[0]
		outcome := round[2]

		score := 0

		switch outcome {
		case 'X': // lose
			switch elfShape {
			case 'A':
				score += 0 + 3
			case 'B':
				score += 0 + 1
			case 'C':
				score += 0 + 2
			}
		case 'Y': // draw
			switch elfShape {
			case 'A':
				score += 3 + 1
			case 'B':
				score += 3 + 2
			case 'C':
				score += 3 + 3
			}
		case 'Z': // win
			switch elfShape {
			case 'A':
				score += 6 + 2
			case 'B':
				score += 6 + 3
			case 'C':
				score += 6 + 1
			}
		}

		ans2 += score
	}

	fmt.Println(ans)
	fmt.Println(ans2)
}
