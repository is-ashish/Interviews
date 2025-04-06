package main

import (
	"fmt"
	"os"
)

func main_() {
	f, err := os.Open("./asic.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}
