package main
import "fmt"

func main(){
	x := make(map[string]int)
	x["key"] = 11
	fmt.Println(x["key"])
	name,ok := x["ke"]
	fmt.Println(name,ok)
}