// addition + => 15 + 10 => (15 ve 10) -> operand , + -> operator
// substruction -
// product *
// division /
// remainder %

package main

import "fmt"

func main() {

	x, y := 15, 10

	fmt.Printf("X : %v , %T\n", x, x)
	fmt.Printf("Y : %v , %T\n", y, y)
	fmt.Printf("X + Y : %v , %T\n", (x + y), (x + y))
	fmt.Printf("X - Y : %v , %T\n", (x - y), (x - y))
	fmt.Printf("X * Y : %v , %T\n", (x * y), (x * y))
	fmt.Printf("X / Y : %v , %T\n", (x / y), (x / y))
	fmt.Printf("X % Y : %v , %T\n", (x % y), (x % y))

	fmt.Printf("X / Y : %v , %T\n", (9.0 / 3), (9.0 / 3))

	z := 5 / 7
	fmt.Printf("Z : %v , %T\n", z, z)

	t := 5.0 / 2
	fmt.Printf("T : %v , %T\n", t, t)

	// Increment ++, Decrement -- POSTFIX

	a := 10

	fmt.Println(a)

	a = a - 1

	fmt.Println(a)

	a--

	fmt.Println(a)
}
