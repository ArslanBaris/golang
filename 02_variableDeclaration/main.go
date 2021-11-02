// Package clause
package main

// Import statement
import "fmt"

// My Code
func main() {

	var firstName, lastName string = "Baris", "Arslan" // var - name of variable - static type

	var age int
	age = 21

	var isMarried bool
	isMarried = false

	fmt.Println(firstName, lastName)
	fmt.Println(age)
	fmt.Println(isMarried)

	name := "Baris"
	age2 := 40
	isMarried2 := true

	age = 41

	fmt.Println(name)
	fmt.Println(age2)
	fmt.Println(isMarried2)

}
