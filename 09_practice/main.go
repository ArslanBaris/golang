package main

import (
	"fmt"
	"strconv"
)

func main() {

	// 1 : int x, float64 y type conversion sample

	x := 75
	var y float64

	y = float64(x) // type(value)

	fmt.Println(y)

	// 2 : multiple assing sample x, y = y, x

	a := 5
	b := 10

	fmt.Println("A:", a, "B:", b)

	a, b = b, a // a = b , b = a

	fmt.Println("A:", a, "B:", b)

	// 3 : non English variable names

	age := 21

	fmt.Println(age)

	名稱 := "Baris"
	年齡 := 21

	fmt.Println("Name:", 名稱, "Age:", 年齡)

	// 4 : shadowing kavramı? ( Gölgeleme )

	x = 5

	if true {
		x := 10
		fmt.Println(x)
	}

	fmt.Println(x)

	// 5 : 40 as a string

	a = 65

	s := string(a)

	fmt.Printf("%v , %T\n", a, a) // 65
	fmt.Printf("%v , %T\n", s, s) // A

	c := strconv.Itoa(a)
	fmt.Printf("%v , %T\n", c, c) // "A"
}
