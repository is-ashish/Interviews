package main

import (
	"fmt"
)

func main_() {
	l := LinkedList{}
	l.add(4)
	l.add(2)
	l.add(1)
	l.add(7)
	l.add(9)
	l.display(l.head)
	l.splitList(l.head)
}

// Say number eg = 134
// ans = one,three,four
func sayNumber(num int) {
	numberName := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	if num == 0 {
		return
	}
	rem := num % 10
	num = num / 10
	sayNumber(num)
	fmt.Printf(numberName[rem] + ",")
}

// sum of array elements = []int{1,2,3,4,5,6,7,8}
// and = 36
func sumOfarrayElements(arr []int, sum int) {
	len := len(arr)
	if len == 0 {
		fmt.Println("Sum of array Elements is -->>", sum)
		return
	}
	sum = sum + arr[len-1]
	arr = arr[:len-1]
	sumOfarrayElements(arr, sum)
}

// check array is sorted or not ??
func IsSortedElements(arr []int) {
	len := len(arr)
	if len == 0 || len == 1 {
		fmt.Println("True")
		return
	}
	if arr[len-2] <= arr[len-1] {
		arr = arr[:len-1]
		IsSortedElements(arr)
	} else {
		fmt.Println("False")
		return
	}
}

// search element in array
func searchElements(arr []int, key int) {
	len := len(arr)
	if len == 0 {
		fmt.Println("Element not found")
		return
	}
	if arr[len-1] == key {
		fmt.Println("Element found at inex ->", len-1)
		return
	}
	arr = arr[:len-1]
	searchElements(arr, key)
}

// Binary search
func binarySearch(arr []int, key int) {
	s := 0
	e := len(arr) - 1
	fmt.Println(isElementFound(arr, s, e, key))
}
func isElementFound(arr []int, s, e, key int) bool {
	mid := s + (e-s)/2
	if s <= e {
		if arr[mid] == key {
			fmt.Println("Element found at index -->>", mid)
			return true
		}
		if arr[mid] < key {
			return isElementFound(arr, mid+1, e, key)
		} else {
			return isElementFound(arr, s, mid-1, key)
		}
	} else {
		return false
	}
}

// reverse string
func reverseGivenString(str string, j int) string {
	if j < 0 {
		return ""
	}
	return string(str[j]) + reverseGivenString(str, j-1)
}

// check palindrome string
func checkPalindromeGivenString(str string, i, j int) bool {
	if i > j {
		return true
	}
	if str[i] == str[j] {
		return checkPalindromeGivenString(str, i+1, j-1)
	} else {
		return false
	}
}

// Optimise 2^1024 power
func OptimisePower(x, y int) int {
	if y == 1 {
		return x
	}
	ans := OptimisePower(x, y/2)
	if y%2 == 0 {
		return ans * ans
	} else {
		return x * ans * ans
	}
}

// Bubble sort
func bubbleSort(arr []int, l int) {
	if l == 0 {
		fmt.Println(arr)
		return
	}
	for j := 0; j < l-1; j++ {
		var temp int
		if arr[j+1] < arr[j] {
			temp = arr[j]
			arr[j] = arr[j+1]
			arr[j+1] = temp
		}
	}
	bubbleSort(arr, l-1)
}

// merge sorting devide
func devide(arr []int) []int {
	l := len(arr)
	if l <= 1 {
		return arr
	}
	mid := l / 2
	s := devide(arr[:mid])
	e := devide(arr[mid:])
	return mergeSort(s, e)
}

// merge sorting conquer
func mergeSort(l, r []int) []int {
	res := make([]int, 0, len(l)+len(r))
	for len(l) > 0 || len(r) > 0 {
		if len(l) == 0 {
			return append(res, r...)
		}
		if len(r) == 0 {
			return append(res, l...)
		}
		if l[0] <= r[0] {
			res = append(res, l[0])
			l = l[1:]
		} else {
			res = append(res, r[0])
			r = r[1:]
		}
	}
	return res
}

// Linked List merge sort --- advance

func (l *LinkedList) add(data int) {
	node := &Node{val: data, next: nil}
	if l.head == nil {
		l.head = node
		l.tail = node
		l.length++
	} else {
		l.tail.next = node
		l.tail = node
		l.length++
		return
	}
}

func (l *LinkedList) display(node *Node) {
	if node == nil {
		fmt.Println("Nil")
	}
	if node.next == nil {
		fmt.Printf("%v -> ", node.val)
	}
	for node != nil {
		fmt.Printf("%v -> ", node.val)
		node = node.next
	}
	fmt.Printf("Nil")
	fmt.Println()
}

func (l *LinkedList) devideList(node *Node) *Node {
	// find middle element of list
	fast_ptr := node
	slow_ptr := node
	if node.next == nil {
		return node
	}
	// middle element
	for fast_ptr != nil && fast_ptr.next != nil {
		fast_ptr = fast_ptr.next.next
		slow_ptr = slow_ptr.next
	}
	fmt.Println("middle of list is -->", slow_ptr.val)
	temp := slow_ptr
	slow_ptr = nil
	s := l.devideList(node)
	e := l.devideList(temp)
	return mergeSortList(s, e)
}

func (l *LinkedList) splitListHalf(node *Node) {
	fast_ptr := node
	slow_ptr := node
	for fast_ptr.next != nil && fast_ptr != nil {
		fast_ptr = fast_ptr.next.next
		slow_ptr = slow_ptr.next
	}
	fmt.Println()
	list2 := slow_ptr.next
	fmt.Println("Fist List -->>")
	for node != slow_ptr.next {
		fmt.Printf("%v -> ", node.val)
		node = node.next
	}
	node = nil
	fmt.Println()
	fmt.Println("Second List -->>")
	for list2 != nil {
		fmt.Printf("%v -> ", list2.val)
		list2 = list2.next
	}

}
