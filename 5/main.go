package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func pop[T any](arr *[]T) T {
	l := len(*arr)
	li := (*arr)[l-1]
	*arr = (*arr)[:l-1]
	return li
}

func main() {
	inputSplit := strings.Split(input, "\n\n")
	crateRows := strings.Split(inputSplit[0], "\n")
	instructions := strings.Split(strings.TrimSpace(inputSplit[1]), "\n")

	stackMap := make(map[int][]string)

	for _, crateRow := range crateRows[:len(crateRows)-1] {
		for i := 0; i < 9*4; i += 4 {
			crate := strings.TrimSpace((crateRow + " ")[i : i+4])
			if crate != "" {
				stackMap[i/4+1] = append([]string{crate}, stackMap[i/4+1]...)
			}
		}
	}

	for _, instruction := range instructions {
		instructionSplit := strings.Split(instruction, " ")
		numCratesToMove, _ := strconv.Atoi(instructionSplit[1])
		fromStack, _ := strconv.Atoi(instructionSplit[3])
		toStack, _ := strconv.Atoi(instructionSplit[5])

		// PART ONE
		// for i := 0; i < numCratesToMove; i++ {
		// 	fromStackCrates := stackMap[fromStack]
		// 	crateToMove := pop(&fromStackCrates)
		// 	stackMap[toStack] = append(stackMap[toStack], crateToMove)
		// 	stackMap[fromStack] = fromStackCrates
		// }

		// PART TWO
		fromStackCrates := stackMap[fromStack]
		cratesToMove := fromStackCrates[len(fromStackCrates)-numCratesToMove:]
		stackMap[toStack] = append(stackMap[toStack], cratesToMove...)
		stackMap[fromStack] = fromStackCrates[:len(fromStackCrates)-numCratesToMove]
	}

	ans := ""

	for i := 1; i < 10; i++ {
		ans += stackMap[i][len(stackMap[i])-1]
	}

	ans = strings.ReplaceAll(ans, "[", "")
	ans = strings.ReplaceAll(ans, "]", "")

	fmt.Println(ans)
}
