package main
import fm "fmt"

func main(){
	
	// declaring a variable
	var num float32 = 5.6
	// declaring a constant
	const PI = 3.142
	const hello string = "Hello"
	const ln2 = 0.693147180559945309417232121458 
	const HARD_EIGHT = (1 << 100)  >> 99
	const x1, x2, x3 = "Hello", 45, 98.9
	const (
		UNKOWN = 0
		MALE = 1
		FEMALE = 2
	)
	fm.Println(num)
	fm.Println(int8(num) + 3)
	fm.Println(PI)
	fm.Println(hello)
	fm.Println(ln2/3.9)
	fm.Println(HARD_EIGHT)
	fm.Println(x1)
	fm.Println(x2)
	fm.Println(x3)
	fm.Println(MALE)
	fm.Println("Fuck Yeaaaaaaah !!!!!!!!!!!!")
}