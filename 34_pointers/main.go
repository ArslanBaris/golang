package main

import "fmt"

func main() {

	name := "baris"

	fmt.Println(name)
	fmt.Println(&name) // & --> address operator

	x := 22
	fmt.Println(x)
	fmt.Println(&x)

	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T, %v\n", &x, &x)

	y := &x
	fmt.Printf("%T, %v\n", y, y)

	z := &name
	fmt.Printf("%T, %v\n", z, z)

	fmt.Println(x)     // 22 	   -> değer
	fmt.Println(&x)    // 0xc00... -> adres
	fmt.Println(*(&x)) // 22	   -> adreste tutulan değer   (deferencing)
	fmt.Println(&(*(&x)))
	fmt.Println(*(&(*(&x))))

	// Pass by value
	x1 := 10
	x2 := x1
	fmt.Println(x1, x2)
	x1 = 5
	fmt.Println(x1, x2)

	t1 := 10
	t2 := &t1
	fmt.Println(t1, t2)  // 10 0xc00...
	fmt.Println(t1, *t2) // 10 10

	*t2 = 3
	fmt.Println(t1, *t2) // 3 3

	t3 := &t1
	*t3 = 5
	fmt.Println(t1, *t2, *t3)

	//a1 := [4]int{1, 10, 100, 1000}  // array pass by value

	a1 := []int{1, 10, 100, 1000}
	a2 := a1

	fmt.Println(a1, a2)

	a2[0] = 3
	fmt.Println(a2)
	fmt.Println(a1) // slice pass by reference

}
