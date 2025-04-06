package main

import (
	"fmt"
	"math"
)

var println = fmt.Println
var printf = fmt.Printf

type Tnode struct {
	val    int
	left   *Tnode
	right  *Tnode
	height int
}

type RNode struct {
	root *Tnode
}

func main_() {
	r := &RNode{}
	values := []int{64, 55, 45, 67, 65, 68}

	// values := []int{64, 55, 45, 56, 67, 65, 66, 68}

	for _, v := range values {
		r.root = populateBinarySearchTree(r.root, v)
	}
	// println("Height of tree is -->> ", r.root.height)
	// println("Pre-Order traversal is -->>")
	// displayPreOrder(r.root)
	// println()
	println("In-Order traversal is -->>")
	displayInOrder(r.root)
	// println()
	// println("Post-Order traversal is -->>")
	// displayPostOrder(r.root)
	println()
	// arr := []int{}
	// println("Print Root to Node Path", rootToNode(r.root, 56, &arr)) // ans - [64,56,55]
	// println("##########", arr)
	fmt.Println("Total count no of nodes -> ", totalNode(r.root))

	// find all leaves of tree
	// println("Print leave node -->>")
	// leaves(r.root)
	// println()
	// println("Print leave sum -->>", leaveSum(r.root))
}

func totalNode(root *Tnode) int {
	if root == nil {
		return 0
	}
	count := 1
	if root.left != nil {
		count = count + totalNode(root.left)
	}
	fmt.Println("lllll ---", root.val)
	if root.right != nil {
		count = count + totalNode(root.right)
	}
	fmt.Println("rrrrrrrr----", root.val)

	return count

}

func rootToNode(r *Tnode, targetNode int, arr *[]int) bool {
	if r == nil {
		return false
	}
	*arr = append(*arr, r.val)
	if r.val == targetNode {
		return true
	}
	if r.left != nil && rootToNode(r.left, targetNode, arr) || r.right != nil && rootToNode(r.right, targetNode, arr) {
		return true
	}
	*arr = (*arr)[:len(*arr)-1]
	return false
}

// Sum of all leaves node by recursion
func leaveSum(root *Tnode) int {
	var sum int
	if root.left == nil && root.right == nil {
		return root.val
	}
	if root.left != nil {
		sum = sum + leaveSum(root.left)
	}

	if root.right != nil {
		sum = sum + leaveSum(root.right)
	}
	return sum
}

// Print all leaves node by recursion
func leaves(root *Tnode) {
	if root.left == nil && root.right == nil {
		fmt.Print(" ", root.val)
		return
	}
	if root.left != nil {
		leaves(root.left)
	}
	if root.right != nil {
		leaves(root.right)
	}
}

func populateBinarySearchTree(node *Tnode, value int) *Tnode {
	if node == nil {
		return newNode(value)
	}
	if value <= node.val {
		node.left = populateBinarySearchTree(node.left, value)
	}

	if value > node.val {
		node.right = populateBinarySearchTree(node.right, value)
	}

	node.height = int(math.Max(float64(height(node.left)), float64(height(node.right))) + 1)

	return node
}

func populateTree(node *Tnode) *Tnode {
	var isLeft, isRight bool
	var i, j int
	println("Do you want to insert Left node of", node.val)
	fmt.Scanf("%t\n", &isLeft)
	if isLeft {
		println("Enter Left Node")
		fmt.Scanf("%d\n", &i)
		node.left = newNode(i)
		populateTree(node.left)
	}
	println("Do you want to insert Right node of", node.val)
	fmt.Scanf("%t\n", &isRight)

	if isRight {
		println("Enter Right Node")
		fmt.Scanf("%d\n", &j)
		node.right = newNode(j)
		populateTree(node.right)
	}

	node.height = int(math.Max(float64(height(node.left)), float64(height(node.right))) + 1)
	return node
}

func newNode(val int) *Tnode {
	node := &Tnode{}
	node.val = val
	node.left = nil
	node.right = nil
	return node
}

func displayTree(node *Tnode) *Tnode {

	if node != nil {
		println("Current Tree node is ------->> ", node.val)
	}
	if node.left != nil {
		println("Left child of Node", node.val, "is", node.left.val)
		displayTree(node.left)
	}
	if node.right != nil {
		println("Right child of Node", node.val, "is", node.right.val)
		displayTree(node.right)
	}
	return node
}

func displayPreOrder(node *Tnode) *Tnode {
	if node == nil {
		return node
	}
	fmt.Print(" ", node.val)
	displayPreOrder(node.left)
	displayPreOrder(node.right)
	return node
}
func displayInOrder(node *Tnode) *Tnode {
	if node == nil {
		return node
	}
	displayInOrder(node.left)
	fmt.Print(" ", node.val)
	displayInOrder(node.right)
	return node
}
func displayPostOrder(node *Tnode) *Tnode {
	if node == nil {
		return node
	}
	displayPostOrder(node.left)
	displayPostOrder(node.right)
	fmt.Print(" ", node.val)
	return node
}

func height(node *Tnode) int {
	if node == nil {
		return -1
	}
	return node.height
}
