package main

import "fmt"

func chanPrinter(c chan int) {
	for val := range c {
		fmt.Println(val)
	}
}

func main() {
	c := make(chan int, 1)

	go chanPrinter(c)

	c <- 1
	c <- 2
	c <- 3
}
