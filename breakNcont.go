package main 
import fm "fmt"

func main(){

	for i:= 0; i<5; i++ {
		for j:= 0; j<10; j++ {
			
			if j > 5 {
				break
			}
			fm.Printf("%d",j)
		}
		fm.Printf("\n")
	}

	for i:= 0; i<10; i++  {

		if i == 5 {
			continue
		}
		fm.Printf("%d",i)
	}
}