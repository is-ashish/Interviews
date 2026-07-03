package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name          string `json:"name"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"-"` // Internal flag
}

// MarshalJSON is a custom implementation for the User struct
func (u User) MarshalJSON() ([]byte, error) {
	// We create an "Alias" to avoid an infinite loop
	type Alias User

	// Logic: If email isn't verified, we hide it in the JSON output
	displayEmail := u.Email
	if !u.EmailVerified {
		displayEmail = "Hidden (Unverified)"
	}

	return json.Marshal(&struct {
		Email string `json:"email"`
		*Alias
	}{
		Email: displayEmail,
		Alias: (*Alias)(&u),
	})
}

func main_() {
	u1 := User{Name: "Alice", Email: "alice@example.com", EmailVerified: true}
	u2 := User{Name: "Bob", Email: "bob@secret.com", EmailVerified: false}

	j1, _ := json.Marshal(u1)
	j2, _ := json.Marshal(u2)

	fmt.Println(string(j1)) // {"email":"alice@example.com","name":"Alice"}
	fmt.Println(string(j2)) // {"email":"Hidden (Unverified)","name":"Bob"}
}
