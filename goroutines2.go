package main
import (
	"fmt"
	"time"
)

func f(n int) {
	for i:=0; i<n ; i++ {
		fmt.Println(n," : ",i)
		time.Sleep(time.Millisecond*1000)
	}
}

func main() {
	for i:=0 ; i<5; i++ {
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
}