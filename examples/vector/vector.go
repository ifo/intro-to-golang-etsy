package main

import "fmt"

type Vector struct {
	X int
	Y int
}

func Add(vec1, vec2 Vector) Vector {
	vec1.X += vec2.X
	vec1.Y += vec2.Y
	return vec1
}

func main() {
	vec1 := Vector{X: 1, Y: 1}
	vec2 := Vector{X: 3, Y: -1}

	vec3 := Add(vec1, vec2)

	fmt.Println(vec3)
}
