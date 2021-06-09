package main

import (
	"fmt"
	"time"
)

/*
A race condition arises in software when a computer program, to operate properly, depends on the sequence or
timing of the program's processes or threads. Critical race conditions cause invalid execution and software bugs
Critical race conditions often happen when the processes or threads depend on some shared state. Operations
upon shared states are done in critical sections that must be mutually exclusive. Failure to obey this rule can
corrupt the shared state.
*/

var data int = 1;

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%s -> #%d\n", s, data)
		data++
	}
}

func main() {
	go say("world")
	go say("concurrent")
	say("hello")
}
