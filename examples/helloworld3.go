package main

import (
	"fmt"

	"./file"
)

func main() {
	bts, _ := file.Read("helloworld.txt")
	fmt.Println(bts)
}
