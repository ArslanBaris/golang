package main

import "fmt"

func main() {

	var x, y, sum int
	x = 5
	y = 10
	sum = x + y

	fmt.Printf("%d ve %d toplamı %d\n", x, y, sum)

	x = 7
	y = 11
	sum = x + y

	fmt.Printf("%d ve %d toplamı %d\n", x, y, sum)

	// Moduler Programming
	// Readable code
	// From Complex To Simple

	fmt.Println(sumFunc(5, 10))

	merhaba()
	fmt.Println(sumFunc(5, 10))
	fmt.Println(sumFunc(3, 5))
	fmt.Println(sumFunc(2, 7))
	merhaba()
	merhaba()

	// return vs print

	/* 	z := sum(5, 10)
	   	fmt.Println(z)
		   sum2(6, 11) */

	merhaba2("Elis", 4)

	// Fonksiyon İsimlendirme
	// ilk karekter harf
	// camel Case -- mySum, myBestFunction
	// paket dışı --> ilk harf büyük

}

// func funcName(params) return type { code  }

func sumFunc(x, y int) int {
	return x + y
}

func sum2(x, y int) {
	fmt.Println(x + y)
}

func merhaba() {
	fmt.Println("Benim Adım Baris")
}

func merhaba2(name string, age int) {
	fmt.Printf("Adım %s, yaşım %d", name, age)
}
