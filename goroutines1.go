package main
import "fmt"

func f(n int) {
	for i:=0; i<n ; i++ {
		fmt.Println(n," : ",i)
	}
}

func main() {
	go f(46)
	var input string
	fmt.Scanln(&input)
}