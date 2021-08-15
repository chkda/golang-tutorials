package main
import fm "fmt"

// global scope
var number int = 5

func main(){

	var dec bool = true
	fm.Println("Original value of number:",number)
	number = 30// Reassigning value
	fm.Printf("Re-assigned value: %d\n",number)
	fm.Printf("Decision value:%t\n",dec)
	var b float32 = 2.65
	fm.Printf("Float value: %v\n",b)

}