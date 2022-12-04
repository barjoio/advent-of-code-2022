package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	ans := 0
	ans2 := 0

	for _, line := range lines {
		elves := strings.Split(line, ",")
		elf1 := elves[0]
		elf2 := elves[1]

		elf1Range := strings.Split(elf1, "-")
		elf1Start, _ := strconv.Atoi(elf1Range[0])
		elf1End, _ := strconv.Atoi(elf1Range[1])

		elf2Range := strings.Split(elf2, "-")
		elf2Start, _ := strconv.Atoi(elf2Range[0])
		elf2End, _ := strconv.Atoi(elf2Range[1])

		if (elf1Start-elf2Start <= 0 && elf1End-elf2End >= 0) || (elf1Start-elf2Start >= 0 && elf1End-elf2End <= 0) {
			ans++
		}

		if !((elf1Start < elf2Start && elf1End < elf2Start) || (elf1Start > elf2End && elf1End > elf2End)) {
			ans2++
		}
	}

	fmt.Println(ans)
	fmt.Println(ans2)
}
