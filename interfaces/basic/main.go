package main

import (
	"fmt"
	"math"
)


type Circle struct{
radius float64
}

type Square struct {
side float64
}

func (s Square) area() float64 {
return  s.side*s.side
}

func (r Circle) area() float64 {
return math.Pi * r.radius * r.radius
}


type Shape interface {
	area() float64
}

func squared(z Shape) float64 {
	return z.area()*z.area()
}

func main() {
	n:=Square{28}
	f:=Circle{367}
	fmt.Println(squared(n))
	fmt.Println(squared(f))
	fmt.Print("The area of the two shapes that were each squared time each other is ")
	fmt.Print(squared(n)*squared(f))
}

