package main

import "fmt"

func main() {

	myMap := map[string]int{ // LITERAL
		"Ahmet":   40,
		"Ayşe":    17,
		"Selim":   14,
		"Mustafa": 70,
	}
	fmt.Println(myMap)
	//fmt.Println(myMap[0])
	fmt.Println(myMap["Ahmet"], myMap["Selim"])
	fmt.Println(myMap["Ahmet Mehmet"]) // 0

	// key - value aynı veri tipinde olmasına gerek yok

	isMarried := map[string]bool{
		"Ahmet":  true,
		"Ayşe":   false,
		"Mahmut": false,
	}
	fmt.Println(isMarried)

	myMap2 := make(map[string]int)
	fmt.Println(myMap2)
	fmt.Println(myMap2["Test"])

	studentGrades := map[string]int{
		"Baris": 80,
		"Ahmet": 29,
		"Selim": 72,
		"Ayşe":  77,
		"Çınar": 0,
	}

	fmt.Println(studentGrades)
	fmt.Println(studentGrades["Baris"])
	fmt.Println(studentGrades["Çınar"])
	//fmt.Println(studentGrades["Elis"])

	//value, ok := studentGrades["Baris"]
	//fmt.Println(value, ok)  // 80 true

	_, ok := studentGrades["Elis"]
	fmt.Println(ok) // false

	fmt.Println(studentGrades)
	studentGrades["Mahmut"] = 55
	fmt.Println(studentGrades)
	fmt.Println(len(studentGrades))

	delete(studentGrades, "Selim")
	fmt.Println(studentGrades)
	fmt.Println(len(studentGrades))

	sg := studentGrades // maps --> pass by reference
	fmt.Println(studentGrades)
	fmt.Println(sg)
	delete(sg, "Arin")
	fmt.Println(sg)
	fmt.Println(studentGrades)

	for k, v := range studentGrades {
		fmt.Println(k, v)
	}
}
