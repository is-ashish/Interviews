package main

import (
	"fmt"
	"math"
)

func main_() {
	arr := []int{2, 3, 1, 1}
	t := 7
	dp_arr := [][]bool{}
	fmt.Println(IsSubsequence(arr, dp_arr, len(arr)-1, t))
}

// using tabulation approach
//
//	Time Complexity: O(n)
//
// Space Complexity: O(1)
func tabulationFindMaxSumNonAdjacent(arr []int) int {
	prev2 := 0
	prev1 := arr[0]
	for i := 1; i < len(arr); i++ {
		pick := arr[i]
		if i > 1 {
			pick = pick + prev2
		}
		nonpick := prev1
		if nonpick > pick {
			prev2 = prev2
		} else {
			prev2 = prev1
		}
		prev1 = pick
	}
	return prev1
}

// using dyanamic memoization approach
// Time Complexity: O(n)
// Space Complexity: O(n)
// non circular
func find_NonSubSequenceMaxSum(arr, dp_arr []int, index int) int {

	if index == 0 {
		return arr[index]
	}
	if index == -1 {
		return 0
	}
	if dp_arr[index] != 0 {
		return dp_arr[index]
	}
	pick := arr[index] + find_NonSubSequenceMaxSum(arr, dp_arr, index-2)
	notpick := find_NonSubSequenceMaxSum(arr, dp_arr, index-1)
	if pick > notpick {
		dp_arr[index] = pick
	} else {
		dp_arr[index] = notpick
	}
	return dp_arr[index]
}

func findMaxSumArray(nums []int) int {

	max_no := math.MinInt
	sum := 0
	start := -1
	end := -1
	for i := 0; i < len(nums); i++ {
		if sum == 0 {
			start = i
		}
		sum = sum + nums[i]
		if sum >= max_no {
			max_no = sum
			end = i
		}
		if sum < 0 {
			sum = 0
		}
	}
	fmt.Println("start", start, end)
	return max_no
}

// Check if there exists a subsequence with sum K
func IsSubsequence(arr []int, dp_arr [][]bool, index, target int) bool {

	if target == 0 {
		return true
	}
	if index == 0 {
		return arr[index] == target
	}
	if dp_arr[index][target] != false {
		return dp_arr[index][target]
	}
	nonTake := IsSubsequence(arr, dp_arr, index-1, target)
	var take bool = false
	if target >= arr[index] {
		take = IsSubsequence(arr, dp_arr, index-1, target-arr[index])
	}
	res := take || nonTake
	dp_arr[index][target] = res
	return res

}
