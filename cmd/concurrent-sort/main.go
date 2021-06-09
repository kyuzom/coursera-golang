package main

import (
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"os"
	"sync"
	"sort"
)

func NumSort(nums []int, wg *sync.WaitGroup) {
	defer wg.Done()
	sort.Ints(nums)
	fmt.Printf("%v\n", nums)
}

func main() {
	var nums []int
	fmt.Print("Please enter your numbers (separated by space): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numsstr := strings.Split(scanner.Text(), " ")
	for _,v := range numsstr {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		nums = append(nums, num)
	}

	var wg sync.WaitGroup
	var qlen int = len(nums)/4
	for i:=0; i<4; i++ {
		wg.Add(1)
		go NumSort(nums[i*qlen:(i+1)*qlen], &wg)
	}
	wg.Wait()

	sort.Ints(nums)
	fmt.Printf("%v\n", nums)
}
