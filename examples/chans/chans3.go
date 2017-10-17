package main

import "fmt"

func main() {
	ints := make(chan int, 1)
	strs := make(chan string, 1)

	ints <- 0
	strs <- "zero"

	go func() { ints <- 1 }()
	go func() { strs <- "one" }()
	go func() { strs <- "two" }()
	go func() { ints <- 2 }()

	// Print everything
forLoopLabel:
	for {
		select {
		case i := <-ints:
			fmt.Println("int", i)
		case s := <-strs:
			fmt.Println("string", s)
		default:
			break forLoopLabel
		}
	}

	fmt.Println("we're done")
}
