package main

import (
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"os"
)

func Swap(nums []int, i int) {
	var tmp int = nums[i]
	nums[i] = nums[i+1]
	nums[i+1] = tmp
}

func BubbleSort(nums []int) {
	var found bool
	for {
		found = false
		var i int
		for i=0; i<len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				Swap(nums, i)
				found = true
			}
		}
		if !found {
			break
		}
	}
}

func main() {
	//var nums = make([]int, 10)
	var nums []int
	fmt.Print("Please enter your numbers (separated by space): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numsstr := strings.Split(scanner.Text(), " ")
	for i, s := range numsstr {
		num, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		//nums[i] = num
		nums = append(nums, num)
		if i == 9 {
			break
		}
	}
	BubbleSort(nums)
	fmt.Printf("%v", nums)
}
