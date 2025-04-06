package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var println = fmt.Println
var printf = fmt.Printf

// ============Tree ====

// type Tnode struct {
// 	Val   int
// 	Right *Tnode
// 	Left  *Tnode
// }
// type Tree struct {
// 	root *Tnode
// }

// ============Linked List==
type Node struct {
	Val  int
	next *Node
}
type LinkedList struct {
	length int
	head   *Node
	tail   *Node
}

func main_() {

	// ================== Linked List ==========================
	list1 := LinkedList{} // 1,2,6,3,4,5,6
	list1.add(7)
	list1.add(7)
	list1.add(7)
	list1.add(7)
	// list1.add(5)
	// list1.add(6)
	list1.display(list1.head)
	list1.display(list1.removeElements(list1.head, 7))
	// list1.deleteNode(list1.head)
	// r := list1.reverseLinkList(list1.head)
	// // print(r.Val)
	// list1.display(r)
	// list2 := LinkedList{}
	// list2.add(1)
	// list2.add(3)
	// list2.add(4)
	// list2.display(list2.head)
	// list3 := LinkedList{}
	// list3.display(list3.mergeTwoList(list1.head, list2.head))
	// ================= Tree ==================================
	// t := Tree{}
	// root := t.addNode(1)
	// root.Left = t.addNode(2)
	// root.Right = t.addNode(3)
	// root.Left.Left = t.addNode(4)
	// root.Left.Right = t.addNode(5)
	// print("Inorder traversal --")
	// t.inOrderTraversal(root)
	// print()
	// print("Level Order traversal is --")
	// t.levelOrderTraversal(root)
	// print()
}

// ============Tree =================================

func (t *Tree) addNode(Val int) *Tnode {
	return &Tnode{Val: Val, Right: nil, Left: nil}
}
func (t *Tree) inOrderTraversal(node *Tnode) {
	if node == nil {
		return
	}
	t.inOrderTraversal(node.Left)
	printf("%v->", node.Val)
	t.inOrderTraversal(node.Right)
}
func (t *Tree) levelOrderTraversal(node *Tnode) {
	// var result [][]int
	if node == nil {
		return
	}
	h := t.treeHeight(node, 0)
	for i := 1; i <= h; i++ {
		// slc := []int{}
		t.printLevelOrder(node, i)
		// result = append(result, t.printLevelOrder(node, i, slc))
		// slc = nil
	}
	// return result
}
func (t *Tree) treeHeight(node *Tnode, level int) int {
	if node == nil {
		return level
	} else {
		l := t.treeHeight(node.Left, level+1)
		r := t.treeHeight(node.Right, level+1)
		if l > r {
			return l
		} else {
			return r
		}
	}

}
func (t *Tree) printLevelOrder(node *Tnode, h int) {
	if node == nil {
		return
	}
	if h == 1 {
		// printf("%v->", node.Val)
		// element = append(element, node.Val)
		// print(element)
		printf("%v->", node.Val)
	} else {
		t.printLevelOrder(node.Left, h-1)
		t.printLevelOrder(node.Right, h-1)
	}
}

// ============Linked List===========================

func (l *LinkedList) len() {
	print(l.length)
}
func (l *LinkedList) display(head *Node) {
	if l.length == 0 {
		print("No node found in Link List")
		return
	}
	n := l.head
	for n != nil {
		fmt.Printf("%v->", n.Val)
		n = n.next
	}
	println("NULL")
}
func (l *LinkedList) add(Val int) {
	node := &Node{Val: Val, next: nil}
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
func (l *LinkedList) mergeTwoList(head1 *Node, head2 *Node) *Node {
	h1 := head1
	h2 := head2
	for h1 != nil && h2 != nil {
		if h1.Val < h2.Val {
			l.add(h1.Val)
			h1 = h1.next
		} else {
			l.add(h2.Val)
			h2 = h2.next
		}
	}
	for h1 != nil {
		l.add(h1.Val)
		h1 = h1.next
	}
	for h2 != nil {
		l.add(h2.Val)
		h2 = h2.next
	}
	return l.head
}
func (l *LinkedList) reverseLinkList(head *Node) *Node {
	cur := head
	var prev *Node
	var next *Node
	if cur != nil {
		next = cur.next
	}
	for cur != nil {
		cur.next = prev
		prev = cur
		cur = next
		if next != nil {
			next = cur.next
		}
	}
	head = prev
	return head
}

// https://leetcode.com/problems/remove-linked-list-elements/
func (l *LinkedList) removeElements(head *Node, val int) *Node {
	cur := head

	for cur != nil {
		if cur.next.Val == val {
			cur.next = cur.next.next
		}
		cur = cur.next
	}
	return head
}

// ============String====================
func reverseString(str string) string {
	print("Given String -", str)
	rns := []rune(str) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	print("Reverse of Given String -")
	return string(rns)
}
func reverseStringAppend(str string) string {
	result := ""
	for _, ele := range str {
		result = string(ele) + result
	}
	return result
}
func reverseStringRec(str string, s string) string {
	if len(str) < 1 {
		return s
	}
	ch := string(str[0])
	return reverseStringRec(str[1:], ch+s)
}
func palindromeString(str string) bool {
	n := len(str)
	isPalindrome := true
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		if string(str[i]) != string(str[j]) {
			isPalindrome = false
			break
		}
	}
	return isPalindrome
}
func multiply(num1 string, num2 string) string {
	int1, _ := strconv.ParseInt(num1, 10, 64)
	int2, _ := strconv.ParseInt(num2, 10, 64)
	temp := int1 * int2
	return strconv.Itoa(int(temp))
}
func checkStringAnagram(str1 string, str2 string) bool {
	arr := [26]int{}
	arr1 := [26]int{}
	l1 := len(str1)
	l2 := len(str2)
	if l1 != l2 {
		return false
	}
	for i := 0; i < l1; i++ {
		arr[str1[i]-97] = arr[str1[i]-97] + 1
		arr1[str2[i]-97] = arr1[str2[i]-97] + 1
	}
	for j := 0; j < len(arr); j++ {
		if arr[j] != arr1[j] {
			return false
		}
	}
	return true
}
func findSubString(str string, str1 string) {
	n := len(str)

	for j := 0; j < n; j++ {
		helperSubString(str[j:], str1)
	}
}
func helperSubString(str string, targetString string) {
	n := len(str)
	if n < 1 {
		return
	}
	if targetString == str {
		print("String present")
		return
	}
	print(str)
	s := str[0 : n-1]
	helperSubString(s, targetString)
}
func checkpalindromeString(str string) {
	n := len(str)
	flag := true
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		if string(str[i]) != string(str[j]) {
			flag = false
			break
		}
	}
	print(flag)
}
func checkpalindromeStringRec(str string) {
	if len(str) < 1 {
		print("false")
	} else {
		str1 := getReverseString("", str)
		if str1 == str {
			print("String is palindrome")

		} else {
			print("not palindrome ")
		}
	}
}
func getReverseString(rev string, st string) string {
	n := len(st)
	if n < 1 {
		return rev
	}
	ch := string(st[len(st)-1])
	st = st[0 : n-1]
	return getReverseString(rev+ch, st)
}
func longestUniqueSubString(str string) int { // ashishman, dvdf
	max_str := 0
	temp_str := ""
	if len(str) <= 1 {
		return len(str)
	}
	for i := 0; i < len(str); i++ {
		j := str[i]
		if strings.Contains(temp_str, string(j)) {
			temp_str = temp_str[strings.Index(temp_str, string(j))+1:] + string(j)
		} else {
			temp_str = temp_str + string(j)
		}
		max_str = checkMaxString(max_str, len(temp_str))
	}
	return max_str
}
func checkMaxString(s1 int, s2 int) int {
	if s1 > s2 {
		return s1
	} else {
		return s2
	}
}
func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	res := 0
	for l, i := 0, 0; i < len(s); i++ {
		if index, ok := m[s[i]]; ok {
			l = max(l, index)
		}
		res = max(res, i-l+1)
		m[s[i]] = i + 1
		println(m)
	}
	return res
}
func Max(s int, l int) int {
	if s < l {
		return l
	}
	return s
}

// ==========bubble sorting ===================
func sortString(str string) string {
	rns := []rune(str)
	for i := 0; i < len(rns)-1; i++ {
		for j := 0; j < len(rns)-i-1; j++ {
			if rns[j] > rns[j+1] {
				rns[j], rns[j+1] = rns[j+1], rns[j]
			}
		}
	}
	return string(rns)
}

// ============Arrays==========================
func twoSum(nums []int, target int) []int {
	indexes := make([]int, 0)
	for i, v := range nums {
		for in, va := range nums {
			if i != in && v+va == target && len(indexes) < 2 {
				indexes = append(indexes, i, in)
			}
		}
	}
	return indexes
}

func findPeak(arr []int) int { // NetApp
	if len(arr) == 0 {
		return -1
	}
	if len(arr) == 1 {
		return arr[0]
	}
	if arr[0] > arr[1] {
		return arr[0]
	} else if arr[len(arr)-1] > arr[len(arr)-2] {
		return arr[len(arr)-1]
	} else {
		for i := 1; i < len(arr)-1; i++ {
			if arr[i-1] < arr[i] && arr[i] > arr[i+1] {
				return arr[i]
			}
		}
	}
	print("Not Found Any Peak Point")
	return -1
}

// =====================HashMap===========================
func firstUniqueCharacter(str string) string { // fetch string into unique charater string
	res := ""
	hm := make(map[byte]int)
	var ch byte
	for i := 0; i < len(str); i++ {
		ch = str[i]
		if index, ok := hm[ch]; ok {
			hm[ch] = index + 1
		} else {
			hm[ch] = 1
		}
	}
	for key, v := range hm {
		if v == 1 {
			res = res + string(key)
		}
	}
	return res
}
func softMapByValue(result map[string]int) {
	// marks = map[art:12 bio:10 english:45 hindi:23 math:78 science:90]
	type subject struct {
		name  string
		marks int
	}
	var res []subject
	for k, v := range result {
		res = append(res, subject{name: k, marks: v})
	}
	print("Printing slice before sorted by value")
	print(res)
	sort.Slice(res, func(i, j int) bool {
		return res[i].marks < res[j].marks
	})
	print("Printing slice after sorted by value")
	print(res)
	print("===========OUTPUT=======================")
	for _, kv := range res {
		fmt.Printf("%s, %d\n", kv.name, kv.marks)
	}
}
