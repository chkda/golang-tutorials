package main
import "fmt"
import "time"

func pinger(c chan string) {
	for i:= 0; ; i ++ {
		c <- "Ping"
	}
}

func ponger(c chan string) {
	for i:= 0; ; i ++ {
		c <- "Pong"
	}
}


func printer(c chan string) {
	for i:=0; ; i++ {
		msg := <- c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main(){
	var c chan string = make(chan string,10)
	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)

}