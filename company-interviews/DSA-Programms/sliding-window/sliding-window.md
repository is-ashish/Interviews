# 1. Maximum Sum Subarray of Size K in Go

This problem asks us to find the maximum sum of any contiguous subarray of size `k`.

---

## 🐢 Non-Optimized (Brute Force) Approach

**Idea:**  
Check every possible subarray of size `k` and compute its sum. Keep track of the maximum.

```go
package main

import "fmt"

func maxSumSubarrayBruteForce(arr []int, k int) int {
    n := len(arr)
    if n < k {
        return 0
    }

    maxSum := 0
    for i := 0; i <= n-k; i++ {
        sum := 0
        for j := i; j < i+k; j++ {
            sum += arr[j]
        }
        if sum > maxSum {
            maxSum = sum
        }
    }
    return maxSum
}

func main() {
    arr := []int{2, 1, 5, 1, 3, 2}
    k := 3
    fmt.Println(maxSumSubarrayBruteForce(arr, k)) // Output: 9
}
Complexity:
Time: O(n * k) (nested loops, recomputing sums)
Space: O(1) (constant extra space)


```
## ⚡ Optimized (Sliding Window) Approach**
**Idea:**
Use a sliding window of size k. Compute the sum of the first window, then slide the window by subtracting the element going out and adding the new element coming in.

```go

Go
package main

import "fmt"

func maxSumSubarraySlidingWindow(arr []int, k int) int {
    n := len(arr)
    if n < k {
        return 0
    }

    // Initial window sum
    windowSum := 0
    for i := 0; i < k; i++ {
        windowSum += arr[i]
    }

    maxSum := windowSum

    // Slide the window
    for i := k; i < n; i++ {
        windowSum += arr[i] - arr[i-k]
        if windowSum > maxSum {
            maxSum = windowSum
        }
    }

    return maxSum
}

func main() {
    arr := []int{2, 1, 5, 1, 3, 2}
    k := 3
    fmt.Println(maxSumSubarraySlidingWindow(arr, k)) // Output: 9
}

```

Complexity:
Time: O(n) (each element processed once)
Space: O(1) (constant extra space)

## 📊 Comparison
Approach | Time Complexity | Space Complexity | Notes
-----------------------------------------------------
Brute Force    | O(n * k)    | O(1) | Simple but inefficient

Sliding Window | O(n) | O(1) | Efficient and widely used

---

# 2. Longest Substring Without Repeating Characters in Go

## Problem
Find the length of the longest substring without repeating characters.  
This is a **classic variable-size sliding window problem**.  
**Key Skill:** Expand/shrink window to maintain uniqueness.

---

## 💻 Go Implementation

```go
package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
    seen := make(map[byte]int) // stores last index of each character
    left, maxLen := 0, 0

    for right := 0; right < len(s); right++ {
        if idx, ok := seen[s[right]]; ok && idx >= left {
            // shrink window from left to remove duplicate
            left = idx + 1
        }
        seen[s[right]] = right
        if right-left+1 > maxLen {
            maxLen = right - left + 1
        }
    }

    return maxLen
}

func main() {
    fmt.Println(lengthOfLongestSubstring("abcabcbb")) // Output: 3 ("abc")
    fmt.Println(lengthOfLongestSubstring("bbbbb"))    // Output: 1 ("b")
    fmt.Println(lengthOfLongestSubstring("pwwkew"))   // Output: 3 ("wke")
}

```
---
📊 Complexity Analysis
Time Complexity: O(n)
Each character is processed once by the sliding window.
Space Complexity: O(min(n, charset))
The map stores at most all unique characters in the string.

✨ Key Insight
The sliding window expands with right and shrinks with left whenever a duplicate character is found. This ensures the substring always remains unique.
---