package main
import fm "fmt"

func main(){


	var tempC float32 = 100
	var newTemp float32
	newTemp = toFahrenheit(tempC)
	fm.Printf("Temperatue in Fahrenheit: %v\n ",newTemp)
}

func toFahrenheit( temp float32) float32 {

	var tempF float32 
	tempF = ((temp * 9) / 5 ) +32
	return tempF
}