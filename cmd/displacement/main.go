package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		var displacement float64
		//s = ½ a t2 + vot + so
		displacement = (0.5*a*math.Pow(t,2))+(v0*t)+s0
		return displacement
	}
}

func main() {
	var a, v0, s0 float64
	var t float64
	fmt.Print("Please enter acceleration: ")
	fmt.Scan(&a)
	fmt.Print("Please enter initial velocity: ")
	fmt.Scan(&v0)
	fmt.Print("Please enter initial displacement: ")
	fmt.Scan(&s0)
	fmt.Print("Please enter time: ")
	fmt.Scan(&t)
	fn := GenDisplaceFn(a, v0, s0)
	fmt.Println(fn(t))
}
