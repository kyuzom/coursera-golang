package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food string
	locomotion string
	noise string
}

type Bird struct {
	food string
	locomotion string
	noise string
}

type Snake struct {
	food string
	locomotion string
	noise string
}

func (self *Cow) Eat() {
	fmt.Println(self.food)
}

func (self *Cow) Move() {
	fmt.Println(self.locomotion)
}

func (self *Cow) Speak() {
	fmt.Println(self.noise)
}

func (self *Bird) Eat() {
	fmt.Println(self.food)
}

func (self *Bird) Move() {
	fmt.Println(self.locomotion)
}

func (self *Bird) Speak() {
	fmt.Println(self.noise)
}

func (self *Snake) Eat() {
	fmt.Println(self.food)
}

func (self *Snake) Move() {
	fmt.Println(self.locomotion)
}

func (self *Snake) Speak() {
	fmt.Println(self.noise)
}

func main() {
	animals := make(map[string]Animal)

	for {
		fmt.Print(">")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		answer := strings.Split(scanner.Text(), " ")
		if len(answer) != 3 {
			fmt.Println("Incorrect number of args. Please try again!")
			continue
		}
		cmdtype, name, joker := answer[0], answer[1], answer[2]

		switch cmdtype {
		case "newanimal":
			switch joker {
			case "cow":
				_, ok := animals[name]
				if ok {
					fmt.Printf("Animal type %s with name %s already existing. Please try again!\n", joker, name)
					continue
				}
				var animal *Cow = new(Cow)
				animal.food = "grass"
				animal.locomotion = "walk"
				animal.noise = "moo"
				animals[name] = animal
			case "bird":
				_, ok := animals[name]
				if ok {
					fmt.Printf("Animal type %s with name %s already existing. Please try again!\n", joker, name)
					continue
				}
				var animal *Bird = new(Bird)
				animal.food = "worms"
				animal.locomotion = "fly"
				animal.noise = "peep"
				animals[name] = animal
			case "snake":
				_, ok := animals[name]
				if ok {
					fmt.Printf("Animal type %s with name %s already existing. Please try again!\n", joker, name)
					continue
				}
				var animal *Snake = new(Snake)
				animal.food = "mice"
				animal.locomotion = "slither"
				animal.noise = "hsss"
				animals[name] = animal
			default:
				fmt.Printf("Unknown animal %s. Please try again!\n", joker)
				continue
			}
		case "query": {
			animal, ok := animals[name]
			if !ok {
				fmt.Printf("Animal with name %s not existing. Please try again!\n", name)
				continue
			}
			switch joker {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Printf("Unknown method %s. Please try again!\n", joker)
				continue
			}
		}
		default:
			fmt.Printf("Unknown command %s. Please try again!\n", cmdtype)
			continue
		}
	}
}
