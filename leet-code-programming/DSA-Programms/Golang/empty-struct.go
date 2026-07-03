package main

import "fmt"

func main_() {
	// Using empty struct in a map to represent a set
	set := make(map[string]struct{})
	set["apple"] = struct{}{}
	set["banana"] = struct{}{}

	if _, exists := set["apple"]; exists {
		fmt.Println("apple is in the set")
	}
}

/**

Empty Struct (`struct{}`) in Go and Its Use in Maps
An empty struct is a struct type with no fields, written as struct{}.
It’s special because it occupies zero bytes in memory,

making it extremely efficient when you need a placeholder but don’t want to waste space.

Why Use Empty Structs in Maps?
Go doesn’t have a built-in set type, but you can simulate one using a map. Normally, you might use map[string]bool to represent a set of strings:
Go
set := make(map[string]bool)
set["apple"] = true
set["banana"] = true
Here, the bool values are just placeholders — you don’t actually care about them, only about the keys. But bool still consumes memory.
Instead, you can use struct{}:

Go
set := make(map[string]struct{})
set["apple"] = struct{}{}
set["banana"] = struct{}{}
struct{}{} is the zero value of an empty struct.
It takes no space, so the map only stores the keys.
This makes it more memory-efficient than map[string]bool.

Example: Checking Membership
Go
package main

import "fmt"

func main() {
    set := make(map[string]struct{})
    set["apple"] = struct{}{}
    set["banana"] = struct{}{}

    if _, exists := set["apple"]; exists {
        fmt.Println("apple is in the set")
    }

    if _, exists := set["orange"]; !exists {
        fmt.Println("orange is not in the set")
    }
}
Output:
apple is in the set
orange is not in the set

Why It’s Better Than `map[string]bool`
Approach | Memory Use | Purpose
-------------------------------
map[string]bool | Stores a bool for each key (1 byte each) | Works fine, but wastes space
map[string]struct{} | Stores nothing for each key (0 bytes) | Pure set representation


*/
