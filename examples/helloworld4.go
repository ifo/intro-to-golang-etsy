package main

import (
	"fmt"

	"./file"
)

func main() {
	bts, _ := file.Read("helloworld.txt")

	fileStr := string(bts)

	fmt.Print(fileStr)
}
