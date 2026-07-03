```go
Length vs Capacity in Go Slices
In Go, slices are built on top of arrays and have two important properties: length and capacity. They often confuse beginners, but the distinction is crucial.

Length (`len`)
The number of elements currently present in the slice.
Represents how many items you can directly access.
Returned by the built-in len() function.
Example:
Go
s := []int{1, 2, 3}
fmt.Println(len(s)) // Output: 3

Capacity (`cap`)
The total number of elements the slice can hold before it needs to allocate a new underlying array.
It’s the size of the underlying array starting from the first element of the slice.
Returned by the built-in cap() function.
Example:
Go
s := make([]int, 3, 5) // length = 3, capacity = 5
fmt.Println(len(s)) // Output: 3
fmt.Println(cap(s)) // Output: 5

Illustration
Go
arr := [5]int{10, 20, 30, 40, 50}
s := arr[1:3] // slice from index 1 to 2

fmt.Println(len(s)) // 2 (elements: 20, 30)
fmt.Println(cap(s)) // 4 (from index 1 to end of arr: 20, 30, 40, 50)

Key Differences
Property | Meaning | Example
----------------------------
Length | Number of elements in the slice | len(s)
Capacity | Maximum elements the slice can grow into before reallocation | cap(s)

✅ Summary:
Length = how many elements are in the slice right now.
Capacity = how many elements it can hold before needing a new backing array.


```