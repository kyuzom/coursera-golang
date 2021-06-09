package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	var book = make(map[string]string)
	var items = [...]string{"name", "address"}
	for _, n := range items {
		var r string
		fmt.Printf("Please enter a %s: ", n)
		fmt.Scanln(&r)
		book[n] = r
	}
	data, err := json.Marshal(book)
	if err != nil {
		fmt.Println("Could not marshal map into json data. ", err)
	} else {
		fmt.Println(string(data))
	}
}
