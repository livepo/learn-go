package main

import "fmt"


 type ListNode struct {
     Val int
     Next *ListNode
 }


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {return l2}
	if l2 == nil {return l1}

	c, out := 0, 0
	priv := new(ListNode)
	head := priv.Next
	p1, p2 := l1, l2
	for ; p1 != nil && p2 != nil; p1, p2 = p1.Next, p2.Next {
		out = c + p1.Val + p2.Val
		if out >= 10 {
			c = 1
			out = out - 10
		} else {
			c = 0
		}
		current := new(ListNode)
		current.Val = out
		if (head == nil) {
			priv.Next = current
			head = current
		} else {
			head.Next = current
			head = head.Next
		}
	}
	l := p1
	if p2 != nil {l = p2}
	for p := l; p != nil; p = p.Next {
		out = c + p.Val
		if out >= 10 {
			c = 1
			out = out - 10
		} else {
			c = 0
		}
		current := new(ListNode)
		current.Val = out
		head.Next = current
		head = head.Next
	}
	if c > 0 {
		current := new(ListNode)
		current.Val = c
		head.Next = current
	}
	return priv.Next
}


func main() {
	var n1 *ListNode = &ListNode{Val: 8}
	var n2 *ListNode = &ListNode{Val: 0}
	var n3 *ListNode = &ListNode{Val: 7}
	var n7 *ListNode = &ListNode{Val: 9}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n7
	l1 := n1

	var n4 *ListNode = &ListNode{Val: 3}
	var n5 *ListNode = &ListNode{Val: 2}
	var n6 *ListNode = &ListNode{Val: 9}
	n4.Next = n5
	n5.Next = n6
	l2 := n4

	p := addTwoNumbers(l1, l2)
	head := p
	for ; head != nil; {
		fmt.Println(head.Val)
		head = head.Next
	}
}
