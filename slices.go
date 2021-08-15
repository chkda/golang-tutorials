package main
import "fmt"

func main(){
	x := make([]int,5,10)
	
	for i:=0; i<len(x); i++ {
		x = append(x,i)
	}
	x = append(x,6)
	fmt.Println(x)
	fmt.Println(cap(x))
}