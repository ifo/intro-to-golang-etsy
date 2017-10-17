package main

import "fmt"

type Vector struct {
	X, Y int
}

func (v *Vector) Add(vec Vector) {
	v.X += vec.X
	v.Y += vec.Y
}

type Stringer interface {
	String() string
}

func (v Vector) String() string {
	return fmt.Sprintf("X: %d, Y: %d", v.X, v.Y)
}

func main() {
	vec1 := Vector{X: 1, Y: 1}
	vec2 := Vector{X: 3, Y: -1}

	vec1.Add(vec2)

	var stringerVec Stringer

	stringerVec = vec1

	fmt.Println(vec1)
	fmt.Println(stringerVec)
}
