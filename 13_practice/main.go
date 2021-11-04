package main

import "fmt"

const a = 14

func main() {

	// 1-) x = x - 10 vs x -=10

	x := 50
	//x = x - 10  // assignment statement
	//x -=10		// assignment operation
	//x += 10
	//x *= 10
	x /= 10
	fmt.Printf("%v, %T\n", x, x)

	// 2-) K = F - 32 / 1.8 + 273   => -40 F kaç K derecedir?

	F := -40
	K := float64(F-32)/1.8 + 273

	fmt.Printf("%v, %T\n", K, K)

	// 3-)

	/* age := 40
	const myAge = age
	fmt.Printf("%v, %T\n", myAge, myAge)
	*/

	// 4-) Sabitlerde Shadowing Kavramı çalışır mı?

	const a = 24

	fmt.Printf("%v, %T\n", K, K)

	// 5-) const x = 4, const y = 5.4,  x + y?

	const i = 4   // typeless
	const j = 5.4 // typeless

	fmt.Printf("%v, %T\n", i+j, i+j)

	// 6-) const x float64 = 6.4 , y := 4 + x, y?

	const z float64 = 6.4
	y := 4 + z
	fmt.Printf("%v, %T\n", y, y) // 10.4 float64
}
