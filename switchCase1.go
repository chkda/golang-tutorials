package main
import fm "fmt"

func main(){

	var num int = 23
	switch num {
	case 20:
		fm.Println("The value is 20")
	case 23:
		fm.Println("The value is 23")
	default:
		fm.Println("The default case is being executed")
	}
}