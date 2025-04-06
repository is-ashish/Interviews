package main

import (
	"fmt"
	"reflect"
)

type Sayer interface {
	Say() string
}
type Cat struct{}

func (c Cat) Say() string { return "meow" }

type Dog struct{}

func (d Dog) Say() string { return "woof" }

type shape interface {
	area() float64
	perimeter() float64
}

// implimentation of rectangle
type rectangle struct {
	length float64
	width  float64
}

func (l rectangle) area() float64 {
	return l.length * l.width
}
func (l rectangle) perimeter() float64 {
	return 2 * (l.length + l.width)
}

// implimentation of circle
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return (2 * 3.14 * c.radius)

}

func main_() {

	r := rectangle{4, 6}
	c := circle{7}
	shape := []shape{r, c}
	for _, s := range shape {
		fmt.Println(reflect.TypeOf(s).Name(), ":Area:", s.area())
		fmt.Println(reflect.TypeOf(s).Name(), ":Perimeter:", s.perimeter())
	}
	// fmt.Println("Area of rectangle is", s.area())
	// fmt.Println("Perimeter of rectangle is", s.perimeter())
	// fmt.Println("Area of circle is ", t.area())
	// fmt.Println("Perimeter of circle is ", t.perimeter())
	println("========================")
	x := Cat{}
	y := Dog{}
	animals := []Sayer{x, y}
	for _, a := range animals {
		fmt.Println(reflect.TypeOf(a).Name(), "says:", a.Say())
	}
	// c := Cat{}
	// fmt.Println("Cat says:", c.Say())
	// d := Dog{}
	// fmt.Println("Dog says:", d.Say())

}
