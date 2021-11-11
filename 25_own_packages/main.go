package main

import (
	"getcelcius"
	"fmt"
)

func main() {
	fmt.Print("Lütfen Celcius dereceyi giriniz: ")
	celcius, err := getCelcius.GetCelcius()
	if err != nil {
		fmt.Println(err)
	}
	fahrenheit := (celcius * 9 / 5) + 32
	fmt.Println(celcius, "Celcius derecesi ", fahrenheit, "Fahrenheit derecesine eşittir.")
}
