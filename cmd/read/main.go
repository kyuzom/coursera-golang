package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

type Book struct {
	fname string
	lname string
}

func main() {
	// init slice of Book with size 0
	books := []Book{}

	var fpath string
	fmt.Printf("Please enter your filename: ")
	fmt.Scanln(&fpath)

	// open file and defer its closing until the end of this function
	fd, err := os.Open(fpath)
	if err != nil {
		fmt.Printf("Could not open file: %s\n", err)
		os.Exit(1)
	}
	defer fd.Close()

	// read file line-by-line
	scanner := bufio.NewScanner(fd)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		var names = strings.Split(scanner.Text(), " ")
		if len(names) != 2 {
			fmt.Printf("Error: %+v has invalid format\n", names)
			os.Exit(2)
		}
		for _, n := range names {
			if len(n) > 20 {
				fmt.Printf("Error: %s is too long (max supported length is 20 chars)\n", n)
				os.Exit(3)
			}
		}
		// append new Struct to the slice
		books = append(books, Book{
			fname: names[0],
			lname: names[1],
		})
	}

	// print out Struct both field and value
	fmt.Printf("%+v\n", books)
}
