package main

import ("fmt"
		"math"
)

// Define inteface
type Shape interface {
	area() float64
}

type circle struct {
	x, y, radius float64
}

type rectangle struct {
	width, height float64
}

func (c circle) area() float64{
	return math.Pi * c.radius * c.radius
}

func (r rectangle) area() float64{
	return r.width * r.height
}

func getArea(s Shape) float64{
	return s.area()
}

func main() {
	circle := circle{0,0,5}
	rect := rectangle{10, 5}

	fmt.Printf("Circle Area: %f\n", getArea(circle))
	fmt.Printf("Rectangle Area: %f\n", getArea(rect))

}