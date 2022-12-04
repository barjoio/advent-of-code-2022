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
	elves := strings.Split(strings.TrimSpace(input), "\n\n")

	ans := 0
	ans2 := 0

	top3 := [3]int{}

	for _, elf := range elves {
		foodItems := strings.Split(elf, "\n")

		calories := 0

		for _, foodItem := range foodItems {
			foodItemCalories, _ := strconv.Atoi(foodItem)
			calories += foodItemCalories
		}

		if calories > ans {
			ans = calories
		}

		if calories > top3[0] && calories > top3[1] && calories > top3[2] {
			top3[1] = top3[0]
			top3[0] = calories
		} else if calories > top3[1] && calories > top3[2] {
			top3[2] = top3[1]
			top3[1] = calories
		} else if calories > top3[2] {
			top3[2] = calories
		}
	}

	ans2 = top3[0] + top3[1] + top3[2]

	fmt.Println(ans)
	fmt.Println(ans2)
}
