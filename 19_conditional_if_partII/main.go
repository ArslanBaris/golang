package main

import "fmt"

func main() {

	x := 10

	if x := -5; x < 0 {
		fmt.Println(x, "negatif sayıdır")
	} else if x%2 == 0 {
		fmt.Println(x, "çift sayıdır")
	} else {
		fmt.Println(x, "tek sayıdır")
	}
	fmt.Println(x)

	if z := -25; z < 0 {
		fmt.Println(z, "negatif sayıdır")
		fmt.Println("Benim Adım Arin")
	} else {
		if z%2 == 0 {
			fmt.Println(z, "çift sayıdır")
		} else {
			fmt.Println(z, "tek sayıdır")
		}
	}

}
