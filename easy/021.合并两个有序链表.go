package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func create(e int) *ListNode {
	return &ListNode{
		Val:  e,
		Next: nil,
	}
}

func (ln *ListNode) add(e int) *ListNode {
	newNode := ListNode{
		Val:  e,
		Next: nil,
	}

	tmp := ln

	for tmp.Next != nil {
		tmp = tmp.Next
	}

	tmp.Next = &newNode
	return ln
}

func (ln *ListNode) display() {
	tmp := ln
	fmt.Print("ListNodes:")
	for tmp != nil {
		fmt.Printf(" %d", tmp.Val)
		tmp = tmp.Next
	}
	fmt.Println()
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var retNode *ListNode
	var tmpNode *ListNode
	if l1 == nil {
		retNode = l2
	} else if l2 == nil {
		retNode = l1
	} else {
		for l1 != nil && l2 != nil {
			if retNode == nil {
				if l1.Val <= l2.Val {
					tmpNode = l1
					l1 = l1.Next
				} else {
					tmpNode = l2
					l2 = l2.Next
				}
				retNode = tmpNode
			} else {
				if l1.Val <= l2.Val {
					tmpNode.Next = l1
					l1 = l1.Next
				} else {
					tmpNode.Next = l2
					l2 = l2.Next
				}
				tmpNode = tmpNode.Next
			}
		}
        if l1 != nil {
			tmpNode.Next = l1
		}
		if l2 != nil {
			tmpNode.Next = l2
		}
	}

	return retNode
}

func main() {
	// s1 := create(1).add(2).add(4)
	// s2 := create(1).add(3).add(4)
	var s1 *ListNode
	// var s2 *ListNode
	s2 := create(0)
	// s := create(1)
	s := mergeTwoList(s1, s2)
	s.display()
}
