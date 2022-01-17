package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {

	merhaba("Baris", 21) // Argument Function Run

	fmt.Println(result(40))

	var x int = 10
	fmt.Println(x)

	var moment time.Time = time.Now() // Now ---> method
	fmt.Println(moment)

	//

	fmt.Print("Lütfen Sınav Sonucunuzu Giriniz: ")
	reader := bufio.NewReader(os.Stdin)
	//value, error := reader.ReadString("\n")
	value, _ := reader.ReadString('\n') // _ blank identifier
	fmt.Println(value)

	//

	bölüm, kalan := bölme(104, 5) // bölüm , _ := bölme(104, 5)  // get only first return
	fmt.Println(bölüm, kalan)

	var result = sumVariadic(1, 4, 6, 3, 11)
	fmt.Println(result)

	numbers := []int{4, 6, 8, 10, 12, 14}
	fmt.Println(sumVariadic(numbers...))

}

func merhaba(name string, age int) { // Parameter Function Write
	fmt.Printf("Adım %s, yaşım %d", name, age)
}

func result(grade int) string {
	if grade >= 50 {
		return "geçtiniz"
	}
	return "kaldınız"
	//fmt.Println("FONKSIYON ICINDEYIM")
}

// 104 / 5 =====> 20 - 4

func bölme(bölünen, bölen int) (bölüm, kalan int) {
	bölüm = bölünen / bölen
	kalan = bölünen % bölen

	return bölüm, kalan
}

func sumVariadic(numbers ...int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum = sum + numbers[i]
	}
	return sum
}
