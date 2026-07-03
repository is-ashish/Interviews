package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Address    string `json:"address"`
	BloodGroup string `json:"bloodGroup"`
}

func main_() {
	u := User{Name: "Ashish", Age: 30, Address: "sarjapur, banglore 562125", BloodGroup: "B+"}
	data, _ := json.MarshalIndent(u, "", " ")
	fmt.Println(string(data))
}

// In Go, JSON marshaling means converting a Go data structure (like a struct, map, or slice) into a JSON-encoded byte sequence. It’s the process of serializing Go values
// into JSON format so they can be stored, transmitted, or shared with other systems.
