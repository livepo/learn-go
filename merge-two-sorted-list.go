package main

import "fmt"


type ListNode struct {
    Val int
    Next *ListNode
}



func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	priv := &ListNode{Val: -1}
	head := priv
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			priv.Next = l1
			l1 = l1.Next
		} else {
			priv.Next = l2
			l2 = l2.Next
		}
		priv = priv.Next
	}
	if l1 != nil {
		priv.Next = l1
	}
	if l2 != nil {
		priv.Next = l2
	}
	return head.Next
}


func main() {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 4}

	n4 := &ListNode{Val: 1}
	n5 := &ListNode{Val: 3}
	n6 := &ListNode{Val: 4}

	n1.Next = n2
	n2.Next = n3

	n4.Next = n5
	n5.Next = n6

	head := mergeTwoLists(n1, n4)

	p := head
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}
}
