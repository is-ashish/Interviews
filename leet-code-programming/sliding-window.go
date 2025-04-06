package main

import (
	"fmt"
	"math"
)

func main() {
	str := "abababab"
	ang := "ab"
	fmt.Println(anagram(str, ang))
}

func anagram(str, s1 string) int {
	var strCount, s1Count [26]int
	var totalCount int
	for i := range s1 {
		strCount[str[i]-'a']++
		s1Count[s1[i]-'a']++
	}
	for i := len(s1); i < len(str); i++ {
		if strCount == s1Count {
			totalCount++
		}
		strCount[str[i-len(s1)]-'a']--
		strCount[str[i]-'a']++
	}
	return totalCount
}

//Find First Negative element of each sub-array of size k. (return answer in form of an array or list
//and in case there is no negative element in the sub-array of size k, then put a 0 for that sub-array)
// ans -: [-1 -1 -16 -16 -16 0]

func FirstNegativeInWindowSizeK(arr []int, k int) []int {
	result := []int{}
	deque := []int{}
	for i := 0; i < k; i++ {
		if arr[i] < 0 {
			deque = append(deque, i)
		}
	}
	if len(deque) > 0 {
		result = append(result, arr[deque[0]])
	} else {
		result = append(result, 0)
	}

	for i := k; i < len(arr); i++ {
		// remove the element that are out of this window
		if len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}
		//Add the index of the current element to the deque if it's negative
		if arr[i] < 0 {
			deque = append(deque, i)
		}
		// Append the first negative element (or 0 if there is none) from the deque to the result list.
		if len(deque) > 0 {
			result = append(result, arr[deque[0]])
		} else {
			result = append(result, 0)
		}
	}

	return result
}

func maxSumInWindow(arr []int, k int) int {
	i, j := 0, 0
	max_sum := 0
	sum := 0
	for j = 0; j < k; j++ {
		sum = sum + arr[j]
	}
	max_sum = sum
	for j < len(arr) {
		sum = sum + arr[j] - arr[i]
		if sum > max_sum {
			max_sum = sum
		}
		i++
		j++
	}
	return max_sum
}

// Given an array and an integer K, find the maximum for each and every contiguous subarray of size K.
// Input: arr[] = {1, 2, 3, 1, 4, 5, 2, 3, 6}, K = 3
func findMaxSubArray(nums []int, k int) int {
	max := math.MinInt
	var i, sum int = 0, 0
	for i < k {
		sum = sum + nums[i]
		i++
	}
	max = sum
	j := k
	for j < len(nums) {
		sum = sum + nums[j] - nums[j-k]
		if max < sum {
			max = sum
		}
		j++
	}
	return max
}

// Permutation in string
// Given two strings s1 and s2, return true if s2 contains a
// permutation of s1, or false otherwise.
// https://leetcode.com/problems/permutation-in-string/description/
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	var s1Count, s2Count [26]int
	for i := range s1 {
		s1Count[s1[i]-'a']++
		s2Count[s2[i]-'a']++
	}
	for j := len(s1); j < len(s2); j++ {
		if s1Count == s2Count {
			return true
		}
		s2Count[s2[j]-'a']++
		s2Count[s2[j-len(s1)]-'a']--
	}

	return s1Count == s2Count
}
