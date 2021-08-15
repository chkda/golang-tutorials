package main
import fm "fmt"

func main(){
	
	var n int = 45
	if n % 5 == 0 {
		fm.Println("Number is divisible by 5")
	} else if n % 3 == 0 {
		fm.Println("Number is divisble by 3")
	} else {
		fm.Println("None of the above")
	}
}