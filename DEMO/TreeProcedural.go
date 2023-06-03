package main

import (
	"fmt"
)

type Node struct {
	value       int
	left, right *Node
}

type Tree *Node

func CreateCompleteBinaryTree(arr []int) Tree {
	return createCompleteBinaryTree(arr, 0, len(arr))
}

func createCompleteBinaryTree(arr []int, start int, size int) *Node {
	curr := &Node{arr[start], nil, nil}
	left := 2*start + 1
	right := 2*start + 2
	if left < size {
		curr.left = createCompleteBinaryTree(arr, left, size)
	}
	if right < size {
		curr.right = createCompleteBinaryTree(arr, right, size)
	}
	return curr
}

func Add(t *Tree, value int) {
	*t = addUtil(*t, value)
}

func addUtil(n *Node, value int) *Node {
	if n == nil {
		return &Node{value, nil, nil}
	}
	if value < n.value {
		n.left = addUtil(n.left, value)
	} else {
		n.right = addUtil(n.right, value)
	}
	return n
}

func PrintPreOrder(t Tree) {
	fmt.Print("Pre Order:")
	printPreOrder(t)
	fmt.Println()
}

func printPreOrder(n *Node) {
	if n != nil {
		fmt.Print(n.value, " ")
		printPreOrder(n.left)
		printPreOrder(n.right)
	}
}

func PrintPostOrder(t Tree) {
	fmt.Print("Post Order:")
	printPostOrder(t)
	fmt.Println()
}

func printPostOrder(n *Node) {
	if n != nil {
		printPostOrder(n.left)
		printPostOrder(n.right)
		fmt.Print(n.value, " ")
	}
}

func PrintInOrder(t Tree) {
	fmt.Print("In Order:")
	printInOrder(t)
	fmt.Println()
}

func printInOrder(n *Node) {
	if n != nil {
		printInOrder(n.left)
		fmt.Print(n.value, " ")
		printInOrder(n.right)
	}
}

func main1() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := CreateCompleteBinaryTree(arr)
	PrintPreOrder(t)
	PrintPostOrder(t)
	PrintInOrder(t)
	Add(&t, 4)
	Add(&t, 44)
	Add(&t, 40)
	PrintPreOrder(t)
	PrintPostOrder(t)
	PrintInOrder(t)
}

func createBinarySearchTreeUtil(arr []int, start int, end int) *Node {
	if start > end {
		return nil
	}
	mid := (start + end) / 2
	return &Node{arr[mid],
		createBinarySearchTreeUtil(arr, start, mid-1),
		createBinarySearchTreeUtil(arr, mid+1, end)}
}

func CreateBinarySearchTree(arr []int) Tree {
	return createBinarySearchTreeUtil(arr, 0, len(arr)-1)
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := CreateBinarySearchTree(arr)
	PrintPreOrder(t)
	PrintPostOrder(t)
	PrintInOrder(t)
	Add(&t, 4)
	Add(&t, 44)
	Add(&t, 40)
	PrintPreOrder(t)
	PrintPostOrder(t)
	PrintInOrder(t)
}
