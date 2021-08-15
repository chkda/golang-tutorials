package main
import "fmt"

type Rect struct {
	height float32
	width float32
}

func main(){
	var x Rect
	y := Rect{height: 4, width: 9}
	z := new(Rect)
	x.height = 5
	x.width = 6
	z.height = 23.78
	z.width = 34
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(*z)

}