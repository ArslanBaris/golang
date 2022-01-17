package main

import "fmt"

type employee struct { // underlying type
	name      string
	age       int
	isMarried bool
}

type manager struct {
	employee
	hasDegree bool
}

// IS A Relation --> Classic OOP
// HAS A Relation

func main() {

	e1 := employee{
		name:      "Gurcan",
		age:       40,
		isMarried: true,
	}

	fmt.Println(e1)
	e2 := e1
	fmt.Println(e2)
	e2.name = "Baris"
	fmt.Println(e2) // e2.name = "Baris"
	fmt.Println(e1) // e1.name = "Gurcan"

	m1 := manager{
		employee: employee{
			name:      "Ayşe",
			age:       28,
			isMarried: false,
		},
		hasDegree: true,
	}

	fmt.Println(m1)

	m2 := manager{}
	m2.name = "Ayşe"
	m2.age = 28
	m2.isMarried = false
	m2.hasDegree = true

	fmt.Println(m2)

	// Anonim Struct

	theBoss := struct {
		name  string
		money bool
	}{name: "THE BOSS", money: true}

	fmt.Println(theBoss)

	e1.save()
}

func (e employee) save() {
	fmt.Println(e.name + " saved.")
}
