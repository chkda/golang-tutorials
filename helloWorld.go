package main 
import "fmt" 
/* main() function is necessary
 It doesn't take any  args */
func main() {
	/*
	To run the code use "go run <fileName>.go".

	To build binaries use "go build <fileName.go". 
	Use "./fileName" to run the binary.
	*/
	fmt.Println("Hello World")
	var input string
	fmt.Scanf("%s",&input)
	fmt.Println(input)
}