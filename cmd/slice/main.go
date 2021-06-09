package main

import (
	"fmt"
	"strconv"
	"sort"
)

func main() {
	nums := make([]int, 0, 3)
	for {
		var r string
		fmt.Print("Please enter a number OR 'X' to quit: ")
		fmt.Scanln(&r)
		if r == "X" {
			break
		}
		num, err := strconv.Atoi(r)
		if err != nil {
			fmt.Println("Error: this is not a number. Try again...")
			continue
		}
		nums = append(nums, num)
		sort.Ints(nums)
		fmt.Println("Numbers until now: ", nums)
	}
}
