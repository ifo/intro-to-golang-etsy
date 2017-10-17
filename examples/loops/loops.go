package main

import (
	"fmt"
	"strconv"
)

func main() {
	basicLoops()
	loopingArrays()
	loopingSlices()
	loopingMaps()
	fizzbuzz()
}

//

//

//

//

//

func basicLoops() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for {
		break
	}

	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("odd", i)
	}
}

//

//

//

//

//

func loopingArrays() {
	intArr := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(intArr); i++ {
		fmt.Println(intArr[i])
	}

	for index, value := range intArr {
		fmt.Printf("index: %d, value: %d\n", index, value)
	}
}

//

//

//

//

//

func loopingSlices() {
	intSlice := []int{1, 2, 3}

	for _, num := range intSlice {
		fmt.Println(num)
	}

	// fitler
	odds := []int{}
	for _, num := range intSlice {
		if num%2 != 0 {
			odds = append(odds, num)
		}
	}

	fmt.Println(odds)
}

//

//

//

//

//

func loopingMaps() {
	convertToInt := map[string]int{"1": 1, "2": 2}
	for key, value := range convertToInt {
		fmt.Printf("key %s, value %d\n", key, value)
	}

	// Add 3
	convertToInt["3"] = 3

	// Keep it weirds, odds only
	delete(convertToInt, "2")

	fmt.Println(convertToInt)

	if _, ok := convertToInt["2"]; ok {
		panic("We don't get here because `ok` is false")
	}
}

//

//

//

//

//

func fizzbuzz() {
	hundredElement := make([]string, 100)

	for i := range hundredElement {
		switch 0 {
		case i % 15:
			hundredElement[i] = "FizzBuzz"
		case i % 5:
			hundredElement[i] = "Buzz"
		case i % 3:
			hundredElement[i] = "Fizz"
		default:
			hundredElement[i] = strconv.Itoa(i)
		}
	}

	for _, str := range hundredElement {
		fmt.Println(str)
	}
}
