package main

import "fmt"

func main() {

	/*
		underArray := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		fmt.Println(underArray)

		mySlc := underArray[2:5] // 2, 3, 4
		fmt.Println(mySlc)

		mySlc2 := underArray[:6] // 1, 2, 3, 4, 5
		fmt.Println(mySlc2)

		mySlc3 := underArray[3:] // 3, 4, 5, 6, 7, 8, 9
		fmt.Println(mySlc3)

		mySlc4 := underArray[:] // 0, 1, 2, 3, 4, 5, 6, 7, 8, 9
		fmt.Println(mySlc4)

		fmt.Println(mySlc)
		mySlc[0] = 100

		fmt.Println(underArray) // 0, 1, 100, 3, ...
		fmt.Println(mySlc)      // 100, 3, 4
		fmt.Println(mySlc2)     // 0, 1, 100, 3, 4, 5
		fmt.Println(mySlc4)     // 0, 1, 100, 3, 4, 5, 6, 7, 8, 9
	*/
	/*
		mySlice := []int{1, 2, 3} // mySliceUnderArray
		fmt.Println(mySlice)

		mySlice = append(mySlice, 4, 5)
		fmt.Println(mySlc)

		mySlice2 := append(mySlice, 4, 5) // mySlice2UnderArray
		fmt.Println(mySlice2)

		mySlice[0] = 100
		fmt.Println(mySlice)
		fmt.Println(mySlice2)
	*/

	/* mySlc := []int{1, 2, 3}
	mySlc2 := []int{4, 5}
	mySlc = append(mySlc, mySlc2...)
	fmt.Println(mySlc)

	mySlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(mySlice)
	mySlice = mySlice[2:]              // ilk n elemanı silmek [n:]
	mySlice = mySlice[:len(mySlice)-3] //  son n elamanı silmek [:len(mySlice)-n]
	fmt.Println(mySlice) */

	var myArr [4]int
	fmt.Println(myArr)

	var mySlc []int
	mySlc = make([]int, 4) // Zero değerler Slice Elemanlarına
	fmt.Println(mySlc)

	var mySlc2 []bool
	mySlc2 = make([]bool, 2) // Zero değerler Slice Elemanlarına
	fmt.Println(mySlc2)

	var mySlc3 []int
	fmt.Printf("%#v", mySlc3)

	fmt.Println()

	mySlc4 := make([]int, 0)
	fmt.Printf("%#v", mySlc4)

}
