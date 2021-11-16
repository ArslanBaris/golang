package main

import "fmt"

func main() { // GO pass by value

	x := 5
	fmt.Println(x)

	double(x)
	fmt.Println(x)

	mySlc := []int{1, 10, 100}
	fmt.Println(mySlc)

	doubleSlc(mySlc)
	fmt.Println(mySlc)

	myArr := [3]int{1, 10, 100}
	fmt.Println(myArr)

	doubleArr(myArr)
	fmt.Println(myArr)

	y := 5
	fmt.Println(y)
	double2(&y)
	fmt.Println(y)
}

func double(num int) {
	num *= 2
	fmt.Println(num)
}

func doubleSlc(slc []int) {

	for i := 0; i < len(slc); i++ {
		slc[i] *= 2
	}
	fmt.Println(slc)
}

func doubleArr(arr [3]int) {

	for i := 0; i < len(arr); i++ {
		arr[i] *= 2
	}
	fmt.Println(arr)
}

func double2(num *int) { // double --> pointer method
	fmt.Println(num)
	*num *= 2
	fmt.Println(*num)
}
