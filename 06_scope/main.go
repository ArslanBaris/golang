package main

import "fmt"

var packVar = "Package Scope"

// packVar := "Package Scope" // paket değişkenlerinde short declaration yapılamaz

var funcVar = "Func(Package) Scope"

func main() {

	var funcVar = "Func Scope"
	fmt.Println(funcVar) // "Func Scope"

	fmt.Println(packVar)

	if true {
		var blockVar = "Block Scope"
		fmt.Println(blockVar)
	}

	// fmt.Println(blockVar)  // blockVar -> undefined

	myFunc()

	var name = "Arin"
	name, surname := "Elis", "Software"
	fmt.Println(name, surname)

}

func myFunc() {
	fmt.Println(packVar)
	// fmt.Println(funcVar)  // funcVar -> undefined
}
