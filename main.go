package main

import "fmt"

// Node is a singly linked list node.
type Node struct {
	Val  int
	Next *Node
}

// Middle returns the middle node — the second middle when the list has an
// even number of nodes — found with slow/fast pointers in one pass.
func Middle(head *Node) *Node {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func fromSlice(vals []int) *Node {
	dummy := &Node{}
	cur := dummy
	for _, v := range vals {
		cur.Next = &Node{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

func main() {
	fmt.Println(Middle(fromSlice([]int{1, 2, 3, 4, 5})).Val)
}
