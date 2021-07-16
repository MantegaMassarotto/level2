package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	length, width float64
}

type Circle struct {
	radius float64
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

func (r *Rectangle) Area() float64 {
	return r.length * r.width
}

func (r *Rectangle) Perimeter() float64 {
	return 2*r.length + 2*r.width
}

func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(s Shape) (float64, float64) {
	a := s.Area()
	p := s.Perimeter()

	return a, p
}

func main() {
	r := &Rectangle{200, 50}
	ra, rp := measure(r)
	fmt.Println("Rectangle area:", ra)
	fmt.Println("Rectangle perimeter:", rp)

	c := &Circle{5}
	ca, cp := measure(c)
	fmt.Println("Circle area:", ca)
	fmt.Println("Circle perimeter:", cp)
}
