package main

import "fmt"

func main() {

	// myArr := [3]int{1, 2, 3}
	// myArr2 := [...]int{1, 2, 3, 4}

	// fmt.Println(myArr)
	// fmt.Println(myArr2)

	// fmt.Println(len(myArr))
	// fmt.Println(len(myArr2))

	// var myArr3 [5]int
	// fmt.Println(myArr3)

	// mySlc := []int{1, 2, 3} // LITERAL
	// fmt.Println(mySlc)
	// fmt.Println(len(mySlc))

	// var myArr4 [4]int
	// fmt.Println(myArr4)
	// myArr4[0] = 5
	// fmt.Println(myArr4)

	// var mySlc2 []int
	// mySlc2 = make([]int, 4) // MAKE METHOD

	// fmt.Println(mySlc)

	// mySlc[0] = 10
	// fmt.Println(mySlc)

	// myArray := [3]int{1, 2, 3}
	// fmt.Println(myArray)

	// myArray2 := myArray
	// fmt.Println(myArray2)

	// myArray2[0] = 100
	// fmt.Println(myArray2)
	// fmt.Println(myArray) // Pass By Value

	mySlice := []int{1, 2, 3}
	fmt.Println(mySlice)

	mySlice2 := mySlice
	fmt.Println(mySlice2)

	mySlice[0] = 33
	fmt.Println(mySlice2)
	fmt.Println(mySlice) // Pass By Reference

	mySlice[0] = 40
	fmt.Println(mySlice2)
	fmt.Println(mySlice)
}
