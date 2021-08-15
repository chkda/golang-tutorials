package main 
import fm "fmt"

func main(){

	var a int = 34
	fm.Printf("The value is %d and the address is %p\n",a,&a)
	var intP *int 
	intP = &a 
	fm.Printf("The value is %d and the address is %p\n",a,intP)
	fm.Printf("The position is %p and the value is %d\n",intP,*(intP))
	*intP = 26
	fm.Printf("The changed value: %d\n",a)
}