package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Method with value receiver
func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

// Method with pointer receiver (can modify struct)
func (p *Person) Birthday() {
	p.Age++
}

func main_() {
	person := Person{Name: "Ashish", Age: 25}
	person.Greet()    // Calls method
	person.Birthday() // Increments age
	person.Greet()    // Shows updated age
}
