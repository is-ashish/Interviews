
#### 🧠 Question 1: Largest Distance Between Consecutive Peak Indices
Problem Statement:
Given an array of integers, find the largest distance between any two consecutive peak indices.
A peak index is defined as an index i such that:
1 ≤ i ≤ n - 2 (i.e., not the first or last element)
arr[i] > arr[i - 1] and arr[i] > arr[i + 1]
Return the maximum distance between any two consecutive peak indices.
If there are fewer than two peaks, return 0.

✨ Example 1
Go
Input: arr = [1, 5, 2, 3, 1]
Peaks: index 1 (5), index 3 (3)
Distance = 3 - 1 = 2
Output: 2
✨ Example 2
Go
Input: arr = [1, 2, 1, 3, 2, 1, 4, 1]
Peaks: index 1 (2), index 3 (3), index 6 (4)
Distances: 3 - 1 = 2, 6 - 3 = 3
Output: 3


```go
✅ Go Solution
Go
package main

import "fmt"

func solvePeaksProblem(arr []int) int {
    peaks := []int{}
    n := len(arr)

    // Step 1: Find all peak indices
    for i := 1; i < n-1; i++ {
        if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
            peaks = append(peaks, i)
        }
    }

    // Step 2: If fewer than 2 peaks, return 0
    if len(peaks) < 2 {
        return 0
    }

    // Step 3: Find max distance between consecutive peaks
    maxDist := 0
    for i := 1; i < len(peaks); i++ {
        dist := peaks[i] - peaks[i-1]
        if dist > maxDist {
            maxDist = dist
        }
    }

    return maxDist
}

func main() {
    arr := []int{1, 2, 1, 3, 2, 1, 4, 1}
    fmt.Println(solvePeaksProblem(arr)) // Output: 3
}

```


#### 🧠 Question 2: Problem Statement:
You are given an array arr of length n, which is a permutation of integers from 1 to n.
Your task is to determine the maximum length of a contiguous subarray that can be removed such that the remaining array is still a non-empty permutation.
🔹 Definitions:
• A subarray is a contiguous segment of the array.
• A permutation of length k is a sequence containing all integers from 1 to k exactly once.
🔸 Input:
• An array arr of length n where 1 ≤ arr[i] ≤ n and all elements are unique.
🔸 Output:
• Return an integer representing the maximum length of a subarray that can be removed while keeping the rest a valid permutation.
```go
Absolutely! Here’s a clean version of the IBM HackerRank-style coding question you uploaded, along with a complete solution in Go:

🧩 Question: Maximum Removable Subarray from a Permutation
Problem Statement:
You are given an array arr of length n, which is a permutation of integers from 1 to n.
Your task is to determine the maximum length of a contiguous subarray that can be removed such that the remaining array is still a non-empty permutation.
🔹 Definitions:
A subarray is a contiguous segment of the array.
A permutation of length k is a sequence containing all integers from 1 to k exactly once.
🔸 Input:
An array arr of length n where 1 ≤ arr[i] ≤ n and all elements are unique.
🔸 Output:
Return an integer representing the maximum length of a subarray that can be removed while keeping the rest a valid permutation.

✨ Example:
Go
Input: arr = [3, 1, 2]
Possible removals:
- Remove [3] → Remaining [1, 2] → Valid permutation
- Remove [1, 2] → Remaining [3] → Not a valid permutation (missing 1 and 2)
Max removable length = 1

Output: 1

✅ Go Solution
Go
package main

import (
	"fmt"
)

func getMaximumSubarray(arr []int) int {
	n := len(arr)
	maxRemovable := 0

	// Try all possible subarrays [i, j]
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			// Build remaining array after removing arr[i..j]
			remaining := append([]int{}, arr[:i]...)
			remaining = append(remaining, arr[j+1:]...)

			if isValidPermutation(remaining) {
				removalLength := j - i + 1
				if removalLength > maxRemovable {
					maxRemovable = removalLength
				}
			}
		}
	}
	return maxRemovable
}

func isValidPermutation(arr []int) bool {
	n := len(arr)
	if n == 0 {
		return false
	}
	seen := make(map[int]bool)
	for _, val := range arr {
		if val < 1 || val > n || seen[val] {
			return false
		}
		seen[val] = true
	}
	return true
}

func main() {
	arr := []int{3, 1, 2}
	fmt.Println(getMaximumSubarray(arr)) // Output: 1
}

🧠 Time Complexity
Brute-force approach: O(n²) due to nested loops.
Can be optimized using prefix/suffix checks if needed for large n.

Would you like me to rewrite this in a more optimized way for large inputs (up to 2×10⁵), or convert it to Python for practice?

```