```go
Runtime Type Checking in Go (Golang)
Go is a statically typed language, but when working with interfaces (especially interface{} or any), the compiler 
loses information about the underlying type. Runtime type checking is 
the process of recovering or verifying that "type" while the program is running.

(1) Type Assertion
(2) Type Switch

```

package main

import "fmt"

func main_() {
	var k interface{} = "hello go"

	// direct assersion
	val := k.(string)
	fmt.Println(val)

	// safe asserssion by switch
	var i interface{} = make([]int, 5, 50)

	switch v := i.(type) {
	case string:
		fmt.Println("String", v)
	case int:
		fmt.Println("Int", v)
	case bool:
		fmt.Println("bool", v)
	case map[int]int:
		fmt.Println("map", v)
	case []int:
		fmt.Println("slice", v)
	default:
		fmt.Println("type not found", v)
	}

}
