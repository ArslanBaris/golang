// 1-) studentName --> John Doe, grade --> 77, isPassed --> true değişkenlerini
// 3 farklı yöntem ile oluşturup, çıktısını yazdırınız.

package main

import "fmt"

func main() {

	// var studentName = "John Doe"
	// var grade = 77
	// var isPassed = true

	// studentName = "John Doe"
	// grade = 77
	// isPassed = true

	// studentName := "John Doe"
	// grade := 77
	// isPassed := true

	// 2-) yukarıda belirtilen değişkenleri tek satır içerisinde tanımlayınız.

	/* var studentName, grade, isPassed = "John Doe", 77, true

	studentName, grade, isPassed := "John Doe", 77, true */

	// 3-) "Declaration", "Assign", "Initialization", "Initial Value" kavramlarını açıklayınız. (Terminoloji)

	/* var studentName string // Declaration

	studentName = "John Dow" // Assignment

	var studentName string = "Baris Arslan" // Initialization (Initial Value -> "Baris Arslan" ) */

	// 4-) ":=" vs "=" aradaki farkı gösteriniz, double declaration

	var studentName string = "John Doe"
	studentName = "Mahmut Erdem"

	studentName1 := "John Doe"
	studentName1 = "Mahmut Erdem"

	fmt.Println(studentName)
	fmt.Println(studentName1)

}
