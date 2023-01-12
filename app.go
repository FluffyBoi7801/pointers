package main

import "fmt"

type Point struct {
	X int
	Y int
	S string
}

func (p Point) method() {
	fmt.Println(p.X)
	fmt.Println(p.Y)
	fmt.Println(p.S)
}

func main() {
	p1 := Point{
		X: 5,
		Y: 10,
		S: "Go Go GO",
	}
	p1.method()
}

func structs(a int, b int) {
	p1 := Point{
		X: 1,
		Y: 2,
	}
	fmt.Println(p1)
	fmt.Println(p1.X)

	p2 := Point{
		X: 12,
		Y: 24,
		S: "POINTS",
	}

	fmt.Println(p2)
}

func pointers() {
	a := "hello world"
	b := 42

	p := &a

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(p)
	fmt.Println(*p)
}
