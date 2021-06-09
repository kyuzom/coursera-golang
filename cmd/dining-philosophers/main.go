package main

import (
	"fmt"
	"sync"
	"time"
)

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	id int
	leftCS *ChopS
	rightCS *ChopS
}

func (self Philo) eat(host chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i:=0; i<3; i++ {
		// get permission from the host, if buffer is full wait here
		host <- true

		self.leftCS.Lock()
		self.rightCS.Lock()
		fmt.Println("starting to eat", self.id)

		time.Sleep(1*time.Second)

		fmt.Println("finishing eating", self.id)
		self.leftCS.Unlock()
		self.rightCS.Unlock()

		// let host know that the eating is done
		<- host
	}
}

func main() {
	var wg sync.WaitGroup
	// host allows only 2 philosohers to eat at the same time
	var host = make(chan bool, 2)

	chopsticks := make([]*ChopS, 5)
	for i:=0; i<5; i++ {
		chopsticks[i] = new(ChopS)
	}

	philosophers := make([]*Philo, 5)
	for i:=0; i<5; i++ {
		// use module to close the circle (0-4, always the next one)
		philosophers[i] = &Philo{i+1, chopsticks[i], chopsticks[(i+1)%5]}
	}

	for i:=0; i<5; i++ {
		wg.Add(1)
		go philosophers[i].eat(host, &wg)
	}
	wg.Wait()
}
