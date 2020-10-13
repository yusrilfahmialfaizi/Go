package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHello(hasname HasName) {
	fmt.Println("hEllo", hasname.GetName())
}

// implementasi interface 1
type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

// implementasi interface 2
type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var saya Person
	saya.Name = "Saya Yusril"

	sayHello(saya)

	binatang := Animal{Name: "HariMau lewat"}
	sayHello(binatang)
}
