package main
import "fmt"

type Rect struct {
	height float32
	width float32
}

func (r Rect) area() float32 {
	return r.height * r. width
}

func main() {
	coords := Rect{height: 23.5, width: 34.5}
	fmt.Println("The area is ",coords.area())
}