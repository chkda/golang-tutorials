package main
import fm "fmt"

func main(){

	var num int = 0

	switch {
	case num > 0:
		fm.Println("The number is positive")
	case num == 0:
		fm.Println("The number is equal to zero")
	case num < 0:
		fm.Println("The number is negative")
	}
}