package main 
import fm "fmt"

func main(){

	str:= "This is a sample."

	for pos,val := range str {

		fm.Printf("Char:%c Pos:%d\n",val,pos)

	}
}