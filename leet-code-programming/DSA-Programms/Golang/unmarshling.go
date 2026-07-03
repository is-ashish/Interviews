package main

import (
	"encoding/json"
	"fmt"
)

type Customer struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

func main_() {
	// JSON string
	// var arrMp map[string]interface{}

	jsonStr := `{"id":4,"name":"John Doe","country":"USA"}`

	var c Customer

	// Unmarshal JSON into struct
	err := json.Unmarshal([]byte(jsonStr), &c)
	if err != nil {
		fmt.Println("Error unmarshaling:", err)
		return
	}

	// Print Go struct
	fmt.Printf("Struct: %+v\n", c)
	// fmt.Println("Name:", c.Name)
}

// Note - reverse process: JSON unmarshaling in Go. This is how you take a
// JSON string and decode it back into Go structs, maps, or slices.

// Question -:
// Would you like me to also show  how to unmarshal a JSON array (like a list of customers)
// into a slice of structs? That’s very common when dealing with APIs.

package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main__() {
	jsonData := `{"name": "Jane", "age": 28}`
	var user User
	err := json.Unmarshal([]byte(jsonData), &user)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}
	fmt.Println(user.Name, user.Age)
	u, _ := json.MarshalIndent(user, "", " ")
	fmt.Println(string(u))
}
