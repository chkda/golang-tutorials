package main
import fm "fmt"

func main(){

	LABEL1:
	for i:= 0; i<5; i++ {
		for j:= 0; j<10; j++ {

			if j == 4 {
				continue LABEL1
			}
			fm.Printf("i:%d j:%d\n",i,j)
		}
		// fm.Printf("\n")
	}
	
}