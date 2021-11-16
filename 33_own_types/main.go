package main

import (
	"fmt"
	"strings"
)

// struct --> underlying type, employee ---> Defined Type, Named Type

type mile float64
type kilometer float64
type mystring string

func main() {

	var m1 mile
	m1 = 3.2

	fmt.Println(m1)
	fmt.Printf("%T, %v", m1, m1)

	f1 := float64(4.4)
	fmt.Printf("%T, %v", (m1 + mile(f1)), (m1 + mile(f1)))
	fmt.Println()
	fmt.Printf("%T, %v \n", (float64(m1) + f1), (float64(m1) + f1))

	k1 := kilometer(7.8)
	fmt.Printf("%T, %v", k1, k1)
	//fmt.Println(m1 + k1)

	m2 := mile(4.6)
	fmt.Println(m1 + m2)
	fmt.Printf("%.2f\n", (m1 * m2))
	fmt.Println(m1 + m2 + 2.1)
	fmt.Println()
	//fmt.Println(m1 + m2 + "baris")

	s1 := "baris"
	fmt.Println(strings.ToUpper(s1))

	//m1 := mile(4.6)
	//fmt.Println(strings.ToUpper(m1))

	m4 := mile(10)

	k4 := toKilometer(m4)

	fmt.Println(k4)

	m4 = toMile(k4)

	fmt.Println(m4)
}

func toKilometer(m mile) kilometer {
	return kilometer(m * 1.6)
}

func toMile(k kilometer) mile {
	return mile(k * 0.62)
}
