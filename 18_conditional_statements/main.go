package main

import "fmt"

func main() {

	// if <boolean expression> { code }  x%2 == 0 (false)

	x := 27

	if x%2 == 0 {
		fmt.Println(x, "çift sayıdır")
	}

	if x%2 == 1 {
		fmt.Println(x, "tek sayıdır")
	}

	if !true {
		fmt.Println("Mesaj Gosterilecek")
	}

	if 5 > 3 {
		fmt.Println("Mesaj Gosterilecek mi?")
	}

	// if <boolean expression> { code } else { code }

	y := 4

	if y%2 == 0 {
		fmt.Println(y, "çift sayıdır")
	} else {
		fmt.Println(y, "tek sayıdır")
	}

	if false {
		fmt.Println("Mesaj Gosterilecek")
	}

	// if <boolean expression> { code } else if <boolean expression> else { code }

	z := -5

	if z < 0 {
		fmt.Println(z, "negatif sayıdır")
	} else if z%2 == 0 {
		fmt.Println(z, "çift sayıdır")
	} else {
		fmt.Println(z, "tek sayıdır")
	}

}
