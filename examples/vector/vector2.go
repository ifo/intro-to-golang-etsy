package main

import "fmt"

type Vector struct {
	X, Y int
}

func (v Vector) Add(vec Vector) Vector {
	v.X += vec.X
	v.Y += vec.Y
	return v
}

func main() {
	vec1 := Vector{X: 1, Y: 1}
	vec2 := Vector{X: 3, Y: -1}

	vec3 := vec1.Add(vec2)

	fmt.Println(vec3)
}
