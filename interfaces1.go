package main
import "fmt"

type Shape interface {
	area() float32
}

type Circle struct {
	radius float32
}

type Rect struct {
	ht float32
	wd float32
}

func (c Circle) area() float32 {
	return 3.14 * c.radius * c.radius
}

func (r Rect) area() float32 {
	return r.ht * r.wd
}

func sumArea(s...Shape) float32 {
	var total float32
	total = 0
	for _, val := range s {
		total += val.area()
	}
	return total
}

func main() {
	circle := Circle{radius: 0.45}
	rect := Rect{ht:34.5,wd:56.4}
	fmt.Println("Area: ",sumArea(circle,rect))
}