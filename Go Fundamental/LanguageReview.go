package main

import "fmt"

// Package Level Scope
var x int

// Struct
type person struct {
	fname string
	lname string
}

func main() {
	fmt.Println(x)

	x := 7
	fmt.Printf("%T", x)

	xi := []int{2, 4, 7, 9, 42}
	fmt.Println(xi)

	m := map[string]int{
		"Todd": 45,
		"Job":  42,
	}
	fmt.Println(m)

	p1 := person{
		"Miss",
		"Moneypenny",
	}
	fmt.Println(p1)
	p1.speak()

	sa1 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}

	sa1.speak()

	// polymorphism
	saySomething(p1)
	saySomething(sa1)
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, `syas, "Good morning, James."`)
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `says, "Shaken, not stirred."`)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}
