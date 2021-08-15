package main
import fm "fmt"

func main(){

	i:= 0
	HERE:
	fm.Printf("This is %dth iteration\n",i)
	i++
	if i == 5 {
		return
	}
	goto HERE
}