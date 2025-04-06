package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main_() {
	// n := 5
	fmt.Println("Subset of string -->", subsetStringCount("abc", ""))
	fmt.Println("----------- Main Completed--------------")
}

func SumOfNaturalNo(n int) int {
	if n <= 0 {
		return 0
	}
	sum := n + SumOfNaturalNo(n-1)
	fmt.Println(n)
	return sum
}

//First uppercase letter in a string (Iterative and Recursive)
// Input : geeksforgeeKs
// Output : K

func firstUpperCase(str []byte, i int) byte {
	if len(str) == i {
		fmt.Println("Not found uppercase")
		return '_'
	}

	if str[i] >= 'A' && str[i] <= 'Z' {
		return str[i]
	}
	return firstUpperCase(str, i+1)
}

// maze problem find no of path including diagonal
// Print all path from matrix using backtracking

func maze(ans string, r, c int) {
	if r == 1 && c == 1 {
		fmt.Println(ans)
		return
	}
	if r > 1 {
		maze(ans+"C", r-1, c)
	}
	if r > 1 && c > 1 {
		maze(ans+"X", r-1, c-1)
	}
	if c > 1 {
		maze(ans+"R", r, c-1)
	}
}

func diceNo(target int, ans string) {
	if target == 0 {
		fmt.Println(ans)
		return
	}
	for i := 1; i <= target; i++ {
		numStr := strconv.Itoa(i)
		diceNo(target-i, ans+numStr)
	}
}

// https://leetcode.com/problems/letter-combinations-of-a-phone-number/
func letterCombinations(digits, ans string) []string {
	latterMap := [][]string{{"0", "0", "0"}, {"0", "0", "0"}, {"a", "b", "c"}, {"d", "e", "f"}, {"g", "h", "i"}, {"j", "k", "l"}, {"m", "n", "o"}, {"p", "q", "r", "s"}, {"t", "u", "v"}, {"w", "x", "y", "z"}}

	if len(digits) == 0 {
		return []string{ans}
	}
	num, _ := strconv.Atoi(string(digits[0]))
	var str []string
	str = append(str, letterCombinations(digits[1:], ans+latterMap[num][0])...)
	str = append(str, letterCombinations(digits[1:], ans+latterMap[num][1])...)
	str = append(str, letterCombinations(digits[1:], ans+latterMap[num][2])...)
	return str
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

// for i=0, ans[0:0] and ans[0:2]
// for i=1, ans[0:1] and ans[1:2]
// for i=2, ans[0:2] and ans[2:2]
// https://go.dev/play/p/llrbUFxW6kJ

func permutation(str, ans string) int {
	if len(str) == 0 {
		return 1
	}
	count := 0
	for i := 0; i <= len(ans); i++ {
		s := ans[0:i]
		f := ans[i:len(ans)]
		count = count + permutation(str[1:], s+string(str[0])+f)
	}
	return count
}

// find all subset of strings Count
func subsetStringCount(str, ans string) int {
	if len(str) == 0 {
		return 1
	}
	// str[1:] =  cut first index of string
	temp := subsetStringCount(str[1:], ans+string(str[0])) + subsetStringCount(str[1:], ans)
	return temp
}

// find all subset of strings
func subsetString(str, ans string) string {
	if len(str) == 0 {
		return ans
	}
	// str[1:] =  cut first index of string
	temp := subsetString(str[1:], ans+string(str[0])) + ", " + subsetString(str[1:], ans)
	return temp
}

// remove specific string eg "ac" from given string
func removeString(str string) string {
	if len(str) == 0 {
		return str
	}
	if strings.HasPrefix(str, "ace") {
		return removeString(str[3:])
	} else {
		return string(str[0]) + removeString(str[1:])
	}
}

// Pending
// https://leetcode.com/problems/remove-all-occurrences-of-a-substring/
func removeOccurrences(str, part, ans string, i int) string {

	if i == len(str)-len(part) {
		return ans
	}
	if str[i:i+len(part)] != part {
		ans = ans + str[i:i+1]
		fmt.Println("checking ------1 ", ans, "-----2 ", str[i:i+len(part)])
		return removeOccurrences(str, part, ans, i+1)
	} else {
		fmt.Println("found -->", ans)
		return removeOccurrences(str, part, ans, i+2)
	}

}

// remove A's from a given strings: Pattern --- 2
func removeAs_2(str string) string {
	if len(str) == 0 {
		return str
	}
	if string(str[0]) == "a" {
		return removeAs_2(str[1:])
	} else {
		return string(str[0]) + removeAs_2(str[1:])
	}
}

// remove A's from a given strings: Pattern --- 1
func removeAs_1(str, ans string) string {
	if len(str) == 0 {
		return ans
	}
	if string(str[0]) == "a" {
		return removeAs_1(str[1:], ans)
	} else {
		return removeAs_1(str[1:], ans+string(str[0]))
	}
}

func linearSearch(arr []int, i, e, key int) bool {

	if i == e && arr[i] != key {
		return false
	}
	return arr[i] == key || linearSearch(arr, i+1, e, key)
}

func isShorted(arr []int, i, e int) bool {

	if i == e {
		return true
	}
	return arr[i] < arr[i+1] && isShorted(arr, i+1, e)
}

func sampleFunction(i int) {
	if i > 5 {
		return
	}
	fmt.Println("Hello, World!--", i)
	sampleFunction(i + 1)
	fmt.Println("functioon call completed-->", i)
}

// fmt.Println("found element", fibbonocci(50))
func fibbonocci(term int) int {
	if term < 2 {
		return term
	}
	return fibbonocci(term-1) + fibbonocci(term-2)
}

// fmt.Println("found element", binarySearch(1, 100, 82))
func binarySearch(s, e, f int) int {
	mid := (s + e) / 2
	if mid == f {
		return mid
	}
	if f < mid {
		return binarySearch(s, mid, f)
	}
	return binarySearch(mid, e, f)
}

// printNo(6)
func printNo(x int) {
	if x <= 0 {
		return
	}
	fmt.Print(x)
	printNo(x - 1)
	fmt.Println("z", x)
}

// fmt.Println(factorial(10))
func sumNo(x int) int {
	if x <= 1 {
		return 1
	}
	return x + sumNo(x-1)
}

// fmt.Println(sumOfDigits(1111111111))
func sumOfDigits(x int) int {
	if x == 0 {
		return x
	}
	return x%10 + sumOfDigits(x/10)
}

// fmt.Println(mulOfDigits(1111111111))
func mulOfDigits(x int) int {
	if x%10 == x {
		return x
	}
	return x % 10 * mulOfDigits(x/10)
}

func zeroNo(n int, count int) {
	if n == 0 {
		fmt.Println("No of count", count)
		return
	}
	r := n % 10
	if r == 0 {
		count++
	}
	zeroNo(n/10, count)
}
