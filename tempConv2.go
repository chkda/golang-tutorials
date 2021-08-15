package main
import fm "fmt"

// type aliasing
type Celsius float32
type Fahrenheit float32

func main(){

	var tempC Celsius = 100
	var newTemp Fahrenheit
	newTemp = toFahrenheit(tempC)
	fm.Printf("Temperature in Fahrenheit: %v\n",newTemp)

}

func toFahrenheit(temp Celsius) Fahrenheit {

	var tempF Fahrenheit
	// temp = Fahrenheit(temp)
	tempF = Fahrenheit(((temp * 9) / 5 ) +32)
	return tempF
}