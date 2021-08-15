package main
import fm "fmt"

func main(){

	switch num:= 0; {

	case num > 0:
		fm.Println("This is a positive number")
	case num == 0:
		fm.Println("This number is equal to zero")
	case num < 0:
		fm.Println("This is a negative number")
	}
}