package main
import fm "fmt"

func main(){
	var num int = 12
	fm.Println(Seasons(num))

}

func Seasons(month int) string {
	const (
		JANUARY = 1
		FEBRUARY = 2
		MARCH = 3
		APRIL = 4
		MAY = 5
		JUNE = 6 
		JULY = 7 
		AUGUST = 8 
		SEPTEMBER = 9
		OCTOBER  = 10
		NOVEMBER = 11
		DECEMBER = 12
	)

	switch month {
	case JANUARY,FEBRUARY,DECEMBER:
		return "Winter"
	case MARCH,APRIL,MAY:
		return "Spring"
	case JUNE,JULY,AUGUST:
		return "Summer"
	case SEPTEMBER,OCTOBER,NOVEMBER:
		return "Autumn"
	default:
		return "Not a valid input"
	}

}