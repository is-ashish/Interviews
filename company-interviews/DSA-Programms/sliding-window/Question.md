# Sliding Window Problems in Go

Sliding window problems are a favorite in coding interviews, especially for Go developers, because they test your ability to optimize subarray/substring computations without brute force.

---

## 🔑 Common Sliding Window Interview Questions

### 1. Maximum Sum Subarray of Size K
- **Problem:** Find the maximum sum of any contiguous subarray of size `k`.
- **Key Skill:** Fixed-size window logic.

### 2. Sliding Window Maximum (LeetCode 239)
- **Problem:** Given an array and window size `k`, return the maximum in each window.
- **Key Skill:** Efficient use of deque/monotonic queue.

### 3. Longest Substring Without Repeating Characters
- **Problem:** Classic variable-size window problem.
- **Key Skill:** Expand/shrink window to maintain uniqueness.

### 4. Minimum Size Subarray Sum
- **Problem:** Find the smallest subarray with sum ≥ target.
- **Key Skill:** Shrinking window logic.

### 5. Substrings with K Distinct Characters
- **Problem:** Count substrings of length `k` with exactly `k-1` distinct characters.
- **Key Skill:** Hash maps + window sliding.

### 6. Fruit Into Baskets (LeetCode 904)
- **Problem:** Longest subarray with at most 2 distinct integers.
- **Key Skill:** Real-world analogy: picking fruits with limited baskets.

### 7. Check if a Permutation of a String is a Substring
- **Problem:** Given two strings, check if one’s permutation exists in the other.
- **Key Skill:** Frequency counting with sliding window.

---

## 📊 Quick Comparison of Sliding Window Types

| Type              | Example Problem                     | Key Skill Tested       |
|-------------------|-------------------------------------|------------------------|
| Fixed-size window | Max sum subarray of size K          | Efficient updates      |
| Variable-size     | Longest substring without repeats   | Expand/shrink logic    |
| Constraint-based  | Minimum size subarray sum           | Condition handling     |
| Frequency-based   | Permutation substring check         | Hash maps + counts     |

---

## ⚙️ Why These Are Asked
- **Efficiency:** Sliding window reduces `O(n²)` brute force to `O(n)`.  
- **Pattern recognition:** Interviewers want to see if you spot the “window” pattern.  
- **Data structures:** Combines arrays, maps, and queues with algorithmic thinking.  

---

👉 If you’d like, I can write out Go implementations for 2–3 of these (like *Maximum Sum Subarray* and *Longest Substring Without Repeating Characters*) so you have ready-to-practice code snippets.
