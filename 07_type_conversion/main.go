package main

import "fmt"

func main() {

	// 1-)

	/* x := 10
	y := 10.0

	fmt.Printf("%v %T\n", x, x)  // 10   int
	fmt.Printf("%v %T\n", y, y)  // 10.0 float64

	// Type Conversion type(value) => int(y) = 10.0 => 10

	fmt.Println(x + int(y))
	fmt.Printf("%v %T\n", y, y)  */

	// 2-)

	/* 	var x int8 = 10
	   	var y int16 = 10
	   	fmt.Println(x + y) */

	// 3-)

	/* x := 10
	y := 10.0
	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)
	fmt.Println(float64(x) + y) */

	// 4-)

	/* 	var x int16 = 128
	   	var y int8
	   	y = int8(x) // type(value)
	   	fmt.Println(y) */

	// 5-)

	/* 	x := 10
	   	y := "10"
	   	fmt.Printf("%v %T\n", x, x)
	   	fmt.Printf("%v %T\n", y, y)
		   fmt.Println(x + int(y)) */

	// 6-)
	num1 := 106
	str1 := string(num1)

	fmt.Printf("%v %T\n", num1, num1) // 106 int
	fmt.Println()
	fmt.Printf("%v %T\n", str1, str1) // j string

}
