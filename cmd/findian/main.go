package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func main() {
	// read user input together
	//var str string
	//fmt.Print("Please enter a string: ")
	//fmt.Scan(&str)

	// read user input together with spaces
	fmt.Print("Please enter a string: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()

	// assignment logic
	str = strings.ToLower(str)
	if strings.HasPrefix(str, "i") && strings.HasSuffix(str, "n") && strings.Contains(str, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
