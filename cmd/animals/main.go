package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type Animal struct {
	food string
	locomotion string
	noise string
}

func (self *Animal) Init(food, locomotion, noise string) {
	self.food = food
	self.locomotion = locomotion
	self.noise = noise
}

func (self *Animal) Eat() {
	fmt.Println(self.food)
}

func (self *Animal) Move() {
	fmt.Println(self.locomotion)
}

func (self *Animal) Speak() {
	fmt.Println(self.noise)
}

func main() {
	for {
		fmt.Print(">")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		answer := strings.Split(scanner.Text(), " ")
		if len(answer) != 2 {
			fmt.Println("Incorrect number of args. Please try again!")
			continue
		}
		atype, amethod := answer[0], answer[1]
		var animal Animal
		switch atype {
		case "cow":
			animal.Init("grass", "walk", "moo")
		case "bird":
			animal.Init("worms", "fly", "peep")
		case "snake":
			animal.Init("mice", "slither", "hsss")
		default:
			fmt.Printf("Unknown animal %s. Please try again!\n", atype)
			continue
		}
		switch amethod {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Printf("Unknown method %s. Please try again!\n", amethod)
			continue
		}
	}
}
