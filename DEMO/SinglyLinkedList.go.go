package main

import "fmt"

type Node struct {
	Value int
	Next  List
}
type List *Node

func Insert(plist *List, value int) {
	*plist = &Node{Value: value, Next: *plist}
}

func FindLength(list List) int {
	if list == nil {
		return 0
	}
	return 1 + FindLength(list.Next)
}

func PrintList(list List) {
	fmt.Print("List: ")
	printList(list)
}

func printList(list List) {
	if list == nil {
		fmt.Print("\n")
		return
	}
	fmt.Printf("%d ", list.Value)
	printList(list.Next)
}

func Append(plist *List, value int) {
	if *plist == nil {
		*plist = &Node{value, nil}
	} else {
		Append(&((*plist).Next), value)
	}
}

func SearchList(list List, value int) bool {
	if list == nil {
		return false
	}
	if list.Value == value {
		return true
	}
	return SearchList(list.Next, value)
}

func DeleteFirstNode(plist *List) {
	list := *plist
	if list == nil {
		return
	}
	*plist = list.Next
}

func SortedInsert(plist *List, value int) {
	list := *plist
	if list == nil || list.Value > value {
		*plist = &Node{value, list}
	} else {
		for list.Next != nil && list.Next.Value < value {
			list = list.Next
		}
		list.Next = &Node{value, list.Next}
	}
}

func main1() {
	var list List
	Insert(&list, 1)
	Insert(&list, 2)
	Insert(&list, 3)
	Insert(&list, 4)
	PrintList(list)
	fmt.Printf("Length: %d\n", FindLength(list))
	fmt.Printf("Search: %t\n", SearchList(list, 3))
	fmt.Printf("Search: %t\n", SearchList(list, 5))
}

/*
List: 4 3 2 1
Length: 4
Search: true
Search: false
*/

func main2() {
	var list List
	SortedInsert(&list, 3)
	SortedInsert(&list, 1)
	SortedInsert(&list, 4)
	SortedInsert(&list, 2)
	PrintList(list)
}

/*
List: 1 2 3 4
*/

func DeleteNode(plist *List, delValue int) {
	if *plist == nil {
		return
	}
	list := *plist
	if list.Value == delValue { /* first node */
		*plist = list.Next
		return
	}

	for list != nil {
		if list.Next != nil && list.Next.Value == delValue {
			list.Next = list.Next.Next
			return
		}
		list = list.Next
	}
}

func main3() {
	var list List
	SortedInsert(&list, 3)
	SortedInsert(&list, 1)
	SortedInsert(&list, 4)
	SortedInsert(&list, 2)
	PrintList(list) // Output: 1 2 3 4
	DeleteNode(&list, 3)
	PrintList(list) // Output: 1 2 4
}

/*
List: 1 2 3 4
List: 1 2 4
*/

func DeleteNodes(plist *List, delValue int) {
	curr := *plist
	for curr != nil && curr.Value == delValue { /* first node */
		curr = curr.Next
		*plist = curr
	}

	for curr != nil {
		if curr.Next != nil && curr.Next.Value == delValue {
			curr.Next = curr.Next.Next
			continue
		}
		curr = curr.Next
	}
}

func main4() {
	var list List
	Append(&list, 3)
	Append(&list, 1)
	Append(&list, 4)
	Append(&list, 2)
	Append(&list, 3)
	Append(&list, 1)
	Append(&list, 4)
	PrintList(list)
	DeleteNodes(&list, 3)
	PrintList(list)
}

/*
List: 3 1 4 2 3 1 4
List: 1 4 2 1 4
*/

func FreeList(plist *List) {
	*plist = nil
}

func ReverseList(plist *List) {
	curr := *plist
	var prev *Node
	var next *Node

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	*plist = prev
}

func reverseListUtil(currentNode List, nextNode List) List {
	if currentNode == nil {
		return nil
	}

	if currentNode.Next == nil {
		currentNode.Next = nextNode
		return currentNode
	}

	ret := reverseListUtil(currentNode.Next, currentNode)
	currentNode.Next = nextNode
	return ret
}

func ReverseListRecursive(plist *List) {
	*plist = reverseListUtil(*plist, nil)
}

func main5() {
	var list List
	Insert(&list, 1)
	Insert(&list, 2)
	Insert(&list, 3)
	Insert(&list, 4)
	PrintList(list)
	ReverseList(&list)
	PrintList(list)
	ReverseListRecursive(&list)
	PrintList(list)
}

/*
List: 4 3 2 1
List: 1 2 3 4
List: 4 3 2 1
*/

func removeDuplicate(plist *List) {
	list := *plist
	for list != nil {
		if list.Next != nil && list.Value == list.Next.Value {
			list.Next = list.Next.Next
		} else {
			list = list.Next
		}
	}
}

func main6() {
	var list List
	SortedInsert(&list, 3)
	SortedInsert(&list, 1)
	SortedInsert(&list, 4)
	SortedInsert(&list, 2)
	SortedInsert(&list, 3)
	SortedInsert(&list, 1)
	SortedInsert(&list, 4)
	SortedInsert(&list, 2)
	PrintList(list)
	removeDuplicate(&list)
	PrintList(list)
}

/*
List: 1 1 2 2 3 3 4 4
List: 1 2 3 4
*/

func CopyListReversed(list List) List {
	var list2 List
	for list != nil {
		list2 = &Node{list.Value, list2}
		list = list.Next
	}
	return list2
}

func main7() {
	var list List
	Append(&list, 1)
	Append(&list, 2)
	Append(&list, 3)
	Append(&list, 4)
	PrintList(list)
	list2 := CopyListReversed(list)
	PrintList(list2)
}

/*
List: 1 2 3 4
List: 4 3 2 1
List: 1 2 3 4
*/

func CopyList(list List) List {
	if list == nil {
		return nil
	}

	list2 := &Node{list.Value, nil}
	tail := list2
	list = list.Next

	for list != nil {
		tail.Next = &Node{list.Value, nil}
		tail = tail.Next
		list = list.Next
	}
	return list2
}

func CopyList2(list List) List {
	if list == nil {
		return nil
	}
	return copyListRecursive(list)
}

func copyListRecursive(node *Node) *Node {
	if node == nil {
		return nil
	}
	return &Node{
		Value: node.Value,
		Next:  copyListRecursive(node.Next),
	}
}

func main8() {
	var list List
	Append(&list, 1)
	Append(&list, 2)
	Append(&list, 3)
	Append(&list, 4)
	PrintList(list)
	list2 := CopyList(list)
	PrintList(list2)
	list3 := CopyList2(list)
	PrintList(list3)
}

/*
List: 1 2 3 4
List: 1 2 3 4
List: 1 2 3 4
*/

func CompareList(list1 List, list2 List) bool {
	if list1 == nil && list2 == nil {
		return true
	} else if list1 == nil || list2 == nil || list1.Value != list2.Value {
		return false
	} else {
		return CompareList(list1.Next, list2.Next)
	}
}

func CompareList2(list1 List, list2 List) bool {
	for list1 != nil && list2 != nil {
		if list1.Value != list2.Value {
			return false
		}
		list1 = list1.Next
		list2 = list2.Next
	}

	if list1 == nil && list2 == nil {
		return true
	}
	return false
}

func main9() {
	var list List
	Append(&list, 1)
	Append(&list, 2)
	Append(&list, 3)
	Append(&list, 4)
	PrintList(list)
	list2 := CopyList(list)
	PrintList(list2)
	fmt.Println("Comparing list1 and list2:")
	fmt.Println(CompareList(list, list2))

	fmt.Println("Comparing list1 and list2:")
	fmt.Println(CompareList2(list, list2))

	list3 := CopyListReversed(list)
	PrintList(list3)
	fmt.Println("Comparing list1 and list3:")
	fmt.Println(CompareList(list, list3))

	fmt.Println("Comparing list1 and list3:")
	fmt.Println(CompareList2(list, list3))
}

/*
List: 1 2 3 4
List: 1 2 3 4
Comparing list1 and list2:
true
Comparing list1 and list2:
true
*/

func LoopDetect(list List) bool {
	slowPtr := list
	fastPtr := list

	for fastPtr != nil && fastPtr.Next != nil {
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next
		if slowPtr == fastPtr {
			fmt.Println("Loop found")
			return true
		}
	}
	fmt.Println("Loop not found")
	return false
}

func ReverseListloopDetect(plist *List) bool {
	head2 := *plist
	ReverseList(plist)
	if *plist == head2 {
		ReverseList(plist)
		return true
	} else {
		ReverseList(plist)
		return false
	}
}

func LoopTypeDetect(list List) int {
	slowPtr := list
	fastPtr := list

	for fastPtr != nil && fastPtr.Next != nil {
		if list == fastPtr.Next || list == fastPtr.Next.Next {
			fmt.Println("Circular list detected")
			return 2
		}

		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next

		if slowPtr == fastPtr {
			fmt.Println("Loop detected")
			return 1
		}
	}

	fmt.Println("No loop detected")
	return 0
}

func MakeLoop(list List) {
	temp := list
	for temp != nil {
		if temp.Next == nil {
			temp.Next = list
			return
		}
		temp = temp.Next
	}
}

func RemoveLoop(list List) {
	var slowPtr, fastPtr, head, loopNode *Node
	slowPtr = list
	fastPtr = list
	head = list
	loopNode = nil

	for fastPtr != nil && fastPtr.Next != nil {
		fastPtr = fastPtr.Next.Next
		slowPtr = slowPtr.Next

		if fastPtr == slowPtr || fastPtr.Next == slowPtr {
			loopNode = slowPtr
			break
		}
	}

	if loopNode != nil {
		temp := loopNode.Next
		loopLength := 1
		for temp != loopNode {
			loopLength++
			temp = temp.Next
		}

		temp = head
		breakNode := head

		for i := 1; i < loopLength; i++ {
			breakNode = breakNode.Next
		}

		for temp != breakNode.Next {
			temp = temp.Next
			breakNode = breakNode.Next
		}

		breakNode.Next = nil
	}
}

func main10() {
	var list List
	Append(&list, 1)
	Append(&list, 2)
	Append(&list, 3)
	Append(&list, 4)
	PrintList(list)

	MakeLoop(list)
	LoopDetect(list)

	println("LoopTypeDetect:", LoopTypeDetect(list))

	RemoveLoop(list)
	LoopDetect(list)
}

/*
List: 1 2 3 4
Loop found
Circular list detected
LoopTypeDetect: 2
Loop not found
*/

func findIntersection(list1 List, list2 List) *Node {
	l1 := 0
	tempHead := list1
	for tempHead != nil {
		l1++
		tempHead = tempHead.Next
	}

	l2 := 0
	tempHead2 := list2
	for tempHead2 != nil {
		l2++
		tempHead2 = tempHead2.Next
	}

	head := list1
	head2 := list2

	diff := l1 - l2
	if l1 < l2 {
		temp := head
		head = head2
		head2 = temp
		diff = l2 - l1
	}

	for diff > 0 {
		head = head.Next
		diff--
	}

	for head != head2 {
		head = head.Next
		head2 = head2.Next
	}

	return head
}

func findIntersection2(list1 List, list2 List) *Node {
	head1 := list1
	head2 := list2

	for head1 != nil && head2 != nil {
		head1 = head1.Next
		head2 = head2.Next
	}

	var diff *Node
	if head1 == nil {
		diff = head2
		head1 = list1
		head2 = list2
		for diff != nil {
			diff = diff.Next
			head2 = head2.Next
		}
	} else {
		diff = head1
		head1 = list1
		head2 = list2
		for diff != nil {
			diff = diff.Next
			head1 = head1.Next
		}
	}

	for head1 != head2 {
		head1 = head1.Next
		head2 = head2.Next
	}
	return head1
}

func main11() {
	list1 := &Node{1, &Node{2, &Node{3, &Node{4, nil}}}}
	intersectNode := &Node{5, &Node{6, &Node{7, nil}}}
	list2 := &Node{8, &Node{9, nil}}
	list1.Next.Next.Next.Next = intersectNode
	list2.Next.Next = intersectNode

	result := findIntersection(list1, list2)
	fmt.Println("Intersection:", result.Value)

	result2 := findIntersection2(list1, list2)
	fmt.Println("Intersection2:", result2.Value)
}

/*
Intersection: 5
Intersection2: 5
*/

func main() {

	main1()
	main2()
	main3()
	main4()
	main5()
	main6()
	main7()
	main8()
	main9()
	main10()
	main11()
}
