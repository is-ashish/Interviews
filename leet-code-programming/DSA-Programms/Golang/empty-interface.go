// ```go
// In Go, the empty interface is written as interface{}. It’s a special type that can hold values of any type
// because every type in Go implements at least zero methods,
// and thus satisfies the empty interface.

// ** common use case

// Handling values of unknown type (like generic containers).
// Passing arbitrary data (e.g., in logging, JSON decoding).
// Writing functions that accept any type.

// ** Summary
// The empty interface interface{} is Go’s way of representing a value of any type.
// It’s powerful for generic programming,
// but you usually need type assertions or type switches to work with the actual value inside.

// ```

package main

import "fmt"

func main_() {
	var i interface{}
	i = 42
	callInterface(i)

	i = "hello go"
	callInterface(i)

	i = make([]int, 3)
	callInterface(i)

	i = make(map[string]struct{})
	callInterface(i)

}

func callInterface(i interface{}) {

	fmt.Printf("=========================Type checking===========%T\n", i)

	switch val := i.(type) {
	case string:
		fmt.Println("String Found", val)
	case int:
		fmt.Println("Int Found", val)
	case []int:
		fmt.Println("[]int Found", val)
	case map[string]struct{}:
		fmt.Println("[]map Found", val)
	default:
		fmt.Println("Not Found")
	}
}
