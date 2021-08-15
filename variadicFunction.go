package main
import "fmt"

func add(args...int) int {
	total := 0
	for _,v := range args {
		// fmt.Println(i)
		total += v
	}
	return total
}

func main(){
	sum := add(1,2,3,4,5)
	fmt.Println(sum)

}