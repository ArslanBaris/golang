package main

import (
	"fmt"
)
func main(){

	city1 := "istanbul"
	city2 := "roma"
	city3 := "tahran"
	city4 := "belgrad"

	fmt.Println(city1, city2, city3, city4)

	//cities := [4]string{"istanbul", "roma", "tahran", "belgrad"}
	//cities := []string{"istanbul", "roma", "tahran", "belgrad"}
	cities := [...]string{"istanbul", "roma", "tahran", "belgrad"}

	fmt.Println(cities)
	fmt.Println(cities[0])

	fmt.Println(len(cities))

	var myArray [5] int
	fmt.Println(myArray)

	myArray[0] = 100
	fmt.Println(myArray)

	var myArray2 [4] int

	/*
	if myArray == myArray2 {	// myArray != myArray2
		fmt.Println("MESAJ")
	}
	*/

	for i:=0 ; i<len(cities);i++{
		fmt.Println(i,cities[i])
	}

	cities[0] = "ankara"

	myArr := [10]{1,2,3,4,5,6,7,8,9,10}

	myArr = mySquare(myArr)
	fmt.Println(myArr)

	// FOR --- RANGE
	// for index, value := range myArr

	
	for index, city := range cities {
		fmt.Println(index, city)
	} 

	for _, city := range cities {
		fmt.Println(city)
	}

}

func mySquare(arr [10]int) [10]int {
	for i:=0 ; i<len(arr);i++{
		arr[i] = arr[i]*arr[i]
	}
	return arr
}