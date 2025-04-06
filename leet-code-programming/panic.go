package main

import (
	"fmt"
)

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("recovered:", r)
	}
}

func sum(a int, b int) {
	defer recovery()
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	done := make(chan bool, 1)
	divide(a, b, done)
	<-done
}

func divide(a int, b int, done chan bool) {
	fmt.Println("waiting for goroutines")
	done <- true
	fmt.Printf("%d / %d = %d", a, b, a/b)
}

func main_() {
	sum(5, 0)
	fmt.Println("normally returned from main")
}

// package main

// import (
// 	"fmt"
// 	"runtime/debug"
// )

// func fullName(firstName *string, lastName *string) {
// 	defer recoverFullName()

// 	if firstName == nil {
// 		panic("runtime error: first name cannot be nil")
// 	}
// 	if lastName == nil {
// 		panic("runtime error: last name cannot be nil")
// 	}
// 	fmt.Printf("%s %s\n", *firstName, *lastName)
// 	fmt.Println("returned normally from fullName")
// }
// func recoverFullName() {
// 	println("defered from panic mode")
// 	r := recover()
// 	if r != nil {
// 		fmt.Println("recovered from ", r)
// 		debug.PrintStack()
// 	}
// }

// func main_() {
// 	defer fmt.Println("deferred call in main")
// 	firstName := "Elon"
// 	fullName(&firstName, nil)
// 	fmt.Println("returned normally from main")
// }

// package main

// import (
// 	"fmt"
// )

// func slicePanic() {
// 	defer resoleveIndexError()
// 	n := []int{5, 7, 4}
// 	fmt.Println(n[4])
// 	fmt.Println("normally returned from a")
// }
// func resoleveIndexError() {
// 	fmt.Println("Recovering form defer panic")

// 	if r := recover(); r != nil {
// 		fmt.Println("Recovered from", r)
// 	}
// }
// func main_() {
// 	slicePanic()
// 	fmt.Println("normally returned from main")
// }
