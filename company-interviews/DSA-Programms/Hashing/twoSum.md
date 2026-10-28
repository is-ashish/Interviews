Here you go, Ashish — a **complete, interview‑ready cheat sheet** covering all **15 Two‑Sum–style problems**, each with:

✅ Problem summary  
✅ Pattern used  
✅ Time complexity  
✅ Clean **Golang solution**  

This is designed exactly the way top‑tier interviewers expect you to think — fast, pattern‑driven, and hash‑map‑first.

---

# 🧠 **TWO‑SUM MASTER CHEAT SHEET (15 Variants in Go)**  
*(All solutions in Golang, optimized for interviews)*

---

# ✅ **1. Two Sum (Classic) — Hash Map**
**Pattern:** Hashing
**Time:** O(n)

```go
func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i, num := range nums {
        if idx, ok := m[target-num]; ok {
            return []int{idx, i}
        }
        m[num] = i
    }
    return nil
}
```

---

# ✅ **2. Two Sum II — Sorted Array (Two Pointers)**
**Pattern:** Two pointers  
**Time:** O(n)

```go
func twoSumSorted(nums []int, target int) []int {
    l, r := 0, len(nums)-1
    for l < r {
        sum := nums[l] + nums[r]
        if sum == target {
            return []int{l, r}
        } else if sum < target {
            l++
        } else {
            r--
        }
    }
    return nil
}
```

---

# ✅ **3. Two Sum — Return All Unique Pairs**
**Pattern:** Hash set  
**Time:** O(n)

```go
func twoSumPairs(nums []int, target int) [][2]int {
    seen := make(map[int]bool)
    used := make(map[int]bool)
    res := [][2]int{}

    for _, num := range nums {
        comp := target - num
        if seen[comp] && !used[num] && !used[comp] {
            res = append(res, [2]int{comp, num})
            used[num], used[comp] = true, true
        }
        seen[num] = true
    }
    return res
}
```

---

# ✅ **4. Two Sum — Count Number of Pairs**
**Pattern:** Hash map  
**Time:** O(n)

```go
func countTwoSum(nums []int, target int) int {
    freq := make(map[int]int)
    count := 0

    for _, num := range nums {
        comp := target - num
        count += freq[comp]
        freq[num]++
    }
    return count
}
```

---

# ✅ **5. Two Sum — With Duplicates Allowed**
**Pattern:** Hash map  
**Time:** O(n)

```go
func twoSumWithDuplicates(nums []int, target int) int {
    freq := make(map[int]int)
    pairs := 0

    for _, num := range nums {
        comp := target - num
        pairs += freq[comp]
        freq[num]++
    }
    return pairs
}
```

---

# ✅ **6. Two Sum — Streaming Data**
**Pattern:** Hash map storing frequencies  
**Time:** O(1) per operation

```go
type TwoSumStream struct {
    freq map[int]int
}

func NewTwoSumStream() *TwoSumStream {
    return &TwoSumStream{freq: make(map[int]int)}
}

func (t *TwoSumStream) Add(num int) {
    t.freq[num]++
}

func (t *TwoSumStream) Find(target int) bool {
    for num, count := range t.freq {
        comp := target - num
        if comp == num && count > 1 {
            return true
        }
        if comp != num && t.freq[comp] > 0 {
            return true
        }
    }
    return false
}
```

---

# ✅ **7. Two Sum in a BST**
**Pattern:** DFS + Hash set  
**Time:** O(n)

```go
func findTarget(root *TreeNode, k int) bool {
    seen := make(map[int]bool)
    var dfs func(*TreeNode) bool

    dfs = func(n *TreeNode) bool {
        if n == nil { return false }
        if seen[k-n.Val] { return true }
        seen[n.Val] = true
        return dfs(n.Left) || dfs(n.Right)
    }
    return dfs(root)
}
```
## Find alternate solution-  using in-order traversal + two pointers (which takes advantage of the BST’s sorted property)?



---

# ✅ **8. Two Sum in a Linked List**
**Pattern:** Hash set  
**Time:** O(n)

```go
func twoSumList(head *ListNode, target int) bool {
    seen := make(map[int]bool)
    for head != nil {
        if seen[target-head.Val] {
            return true
        }
        seen[head.Val] = true
        head = head.Next
    }
    return false
}
```

---

# ✅ **9. Two Sum — Closest Pair**
**Pattern:** Sort + two pointers  
**Time:** O(n log n)

```go
func closestTwoSum(nums []int, target int) int {
    sort.Ints(nums)
    l, r := 0, len(nums)-1
    best := math.MaxInt

    for l < r {
        sum := nums[l] + nums[r]
        if abs(sum-target) < abs(best-target) {
            best = sum
        }
        if sum < target {
            l++
        } else {
            r--
        }
    }
    return best
}
```

---

# ✅ **10. Two Sum — O(1) Extra Space**
**Pattern:** Sort + two pointers  
**Time:** O(n log n)

*(Same as #2 but without returning original indices.)*

---

# ✅ **11. Two Sum — Large Dataset (Memory Constraints)**
**Pattern:**  
- External sorting  
- Chunk processing  
- Hash partitioning  

*(No code — system design concept.)*

---

# ✅ **12. Two Sum — Distributed System**
**Pattern:**  
- Shard by hash(key)  
- Solve locally  
- Merge results  

*(System design concept.)*

---

# ✅ **13. Two Sum — Using Bitset**
**Pattern:** Bitset for small integer ranges  
**Time:** O(n)

```go
func twoSumBitset(nums []int, target int) bool {
    seen := make([]bool, 1000001)

    for _, num := range nums {
        comp := target - num
        if comp >= 0 && comp < len(seen) && seen[comp] {
            return true
        }
        seen[num] = true
    }
    return false
}
```

---

# ✅ **14. Two Sum — Negative Numbers Allowed**
**Pattern:** Hash map  
**Time:** O(n)

*(Same as classic Two Sum — hashing handles negatives naturally.)*

---

# ✅ **15. Two Sum — Return All Index Pairs**
**Pattern:** Hash map of lists  
**Time:** O(n)

```go
func allTwoSumIndices(nums []int, target int) [][2]int {
    m := make(map[int][]int)
    res := [][2]int{}

    for i, num := range nums {
        comp := target - num
        if idxs, ok := m[comp]; ok {
            for _, idx := range idxs {
                res = append(res, [2]int{idx, i})
            }
        }
        m[num] = append(m[num], i)
    }
    return res
}
```

---

# ✅ **Your Complete Two‑Sum Hashing Toolkit**
This cheat sheet gives you:

✅ All patterns  
✅ All variations  
✅ All Go implementations  
✅ Interview‑ready explanations  

You now have the **full hashing playbook** that FAANG interviewers expect.

If you want, I can also prepare:

📌 A **PDF version** of this cheat sheet  
📌 A **flashcard set** for spaced repetition  
📌 A **practice problem set** with increasing difficulty  

Just tell me what you’d like next.