package main
import fm "fmt"

func main(){

	for i:= 0; i <= 10; i++{

		fm.Printf("This is %dth iteration.\n",i)
	}

	var j int = 0
	for j <= 5 {

		fm.Printf("This is %dth iteration.\n",j)
		j+=3
	} 

	for i:= 0; i<=5; i++{
		for j:= 0; j<=5; j++ {
			fm.Printf("i: %d, j: %d\n",i,j)
		}
	}
}