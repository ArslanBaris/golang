package main

import "fmt"

type employees struct { // underlying type |  like class
	name      string
	age       int
	isMarried bool
	kids      []string
}

func main() {

	var employee struct { // structure
		name      string
		age       int
		isMarried bool
	}

	fmt.Println(employee) // " 0 false" -> default values
	fmt.Printf("%#v\n", employee)
	fmt.Println(employee.age)

	employee.name = "Baris"
	employee.age = 21
	employee.isMarried = false

	fmt.Printf("%#v\n", employee)
	fmt.Println(employee.name)
	fmt.Println(employee.age)
	fmt.Println(employee.isMarried)

	/* 	type employees struct { // underlying type |  like class
		name      string
		age       int
		isMarried bool
	}	*/

	var e1 employees
	e1.name = "Ahmet"
	e1.age = 35
	e1.isMarried = false

	var e2 employees
	e2.name = "Mehmet"
	e2.age = 40
	e2.isMarried = true

	fmt.Printf("%#v\n", e1)
	fmt.Printf("%#v\n", e2)

	e3 := employees{
		name:      "Ayşe",
		age:       7,
		isMarried: false,
	}

	fmt.Printf("%#v\n", e3)

	e4 := employees{
		name:      "Bircan",
		age:       54,
		isMarried: true,
		kids: []string{
			"Ali",
			"Veli",
			"Fatma",
		},
	}

	fmt.Printf("%#v\n", e4)
	fmt.Println(e4)
	fmt.Println(e4.kids)
	fmt.Println(e4.kids[0])

}
