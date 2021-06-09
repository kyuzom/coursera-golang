package main

import (
	"fmt"
	"strconv"
)

func main() {
	var f float64
	var str string
	fmt.Print("Enter a floating point number: ")
	fmt.Scan(&f)
	// the instruction do not specify the behavior so
	// this solution rounds down if < .5 and rounds up if >= .5
	str = strconv.FormatFloat(f, 'f', 0, 32)
	fmt.Println(str)
}
