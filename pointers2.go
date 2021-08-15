package main
import "fmt"

func main(){
	var x int = 35
	xPtr := new(int)
	fmt.Println("This is value ",x)
	fmt.Println("This is pointer ",xPtr)
	xPtr = &x
	*xPtr = 23
	fmt.Println("This is value ",x)
}