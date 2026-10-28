// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	length int
	head   *Node
	tail   *Node
}

func main_() {
	l := LinkedList{}
	l.add(2)
	l.add(6)
	l.add(7)
	l.add(8)
	l.add(9)
	fmt.Println("First List -->")
	l.display(l.head)
	r := LinkedList{}
	r.add(1)
	r.add(3)
	r.add(4)
	r.add(5)
	r.add(8)
	fmt.Println("Second List -->")
	r.display(r.head)
	fmt.Println("Merge two sorted linked lists")
	l.reverseListRecursion(l.head)
}

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

func (l *LinkedList) deleteFirstNode(node *Node) *Node {
	if node.next == nil {
		node = nil
	} else {
		node = node.next
	}
	return node
}

// Merge two sorted linked lists
func (l *LinkedList) mergeSortedLists(list1 *Node, list2 *Node) {
	head1 := list1
	head2 := list2
	list3 := LinkedList{}
	for head1 != nil && head2 != nil {
		if head1.val < head2.val {
			list3.add(head1.val)
			head1 = l.deleteFirstNode(head1)
		} else {
			list3.add(head2.val)
			head2 = l.deleteFirstNode(head2)
		}
	}

	if head1 != nil {
		list3.tail.next = head1
	}
	if head2 != nil {
		list3.tail.next = head2
	}
	list3.display(list3.head)
}

// Merge two sorted linked lists -- using recursion
func (l *LinkedList) mergeSortedListsRecursion(list1 *Node, list2 *Node) *Node {
	var result *Node = nil
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.val <= list2.val {
		result = list1
		result.next = l.mergeSortedListsRecursion(list1.next, list2)
	} else {
		result = list2
		result.next = l.mergeSortedListsRecursion(list1, list2.next)
	}
	return result
}

// With Iterations --
func (l *LinkedList) reverseList(list1 *Node) *Node {

	var prev_node *Node = nil
	var cur_node *Node = list1
	var next_node *Node = nil

	for cur_node != nil {
		next_node = cur_node.next
		cur_node.next = prev_node
		prev_node = cur_node
		cur_node = next_node
	}
	return prev_node
}

// reverse linkedlist using recursion-: method-1
func (l *LinkedList) reverseListRecursion1(head *Node, cur_node *Node, prev_node *Node) *Node {
	var forward *Node
	if cur_node == nil {
		head = prev_node
		return head
	}
	forward = cur_node.next
	cur_node.next = prev_node
	return l.reverseListRecursion(head, forward, cur_node)
}

// reverse linkedlist using recursion-: method-2

func (l *LinkedList) reverseListRecursion2(head *Node) *Node {

	if head == nil || head.next == nil {
		return head
	}
	var small_head *Node = reverseListRecursion2(head.next)
	head.next.next = head
	head.next = nil
	return small_head
}

// #https://leetcode.com/problems/add-two-numbers/?envType=problem-list-v2&envId=linked-list
