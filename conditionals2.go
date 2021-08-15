package main
import fm "fmt"

func main(){

	var first int = 10
	var cond int
	if first <= 0 {
		fm.Println("The number is less than 0")
	} else if (first > 0 && first <= 10) {
		fm.Println("The number is less than 10")
	} else {
		fm.Println("None of the above")
	}

	if cond = 15; cond > 10 {
		fm.Println("The number is greater than 10")
	} else {
		fm.Println("The number is less than 10")
	}

	if num := 43; num > 10 {
		fm.Println("The number is greater than 10")
	} else {

		fm.Println("The number is less than 10")
	}
}