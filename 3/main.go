package main

import (
	_ "embed"
	"fmt"
	"strings"
	"unicode"
)

//go:embed input.txt
var input string

func getPriority(c rune) int {
	if unicode.IsUpper(c) {
		return int(c) - 65 + 27
	} else {
		return int(c) - 97 + 1
	}
}

func main() {
	rucksacks := strings.Split(strings.TrimSpace(input), "\n")

	ans := 0
	ans2 := 0

rucksackLoop:
	for _, rucksack := range rucksacks {
		c1 := rucksack[:len(rucksack)/2]
		c2 := rucksack[len(rucksack)/2:]

		for _, c1item := range c1 {
			for _, c2item := range c2 {
				if c1item == c2item {
					ans += getPriority(c1item)
					continue rucksackLoop
				}
			}
		}
	}

rucksackLoop2:
	for i := 0; i < len(rucksacks); i += 3 {
		elf1 := rucksacks[i]
		elf2 := rucksacks[i+1]
		elf3 := rucksacks[i+2]

		for _, elf1item := range elf1 {
			for _, elf2item := range elf2 {
				if elf1item == elf2item {
					for _, elf3item := range elf3 {
						if elf1item == elf3item {
							ans2 += getPriority(elf3item)
							continue rucksackLoop2
						}
					}
				}
			}
		}
	}

	fmt.Println(ans)
	fmt.Println(ans2)
}
