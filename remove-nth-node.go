package main


import "fmt"


type ListNode struct {
    Val int
    Next *ListNode
}


func removeNthFromEnd(head *ListNode, n int) *ListNode {
	priv := &ListNode{Val: -1, Next: head}
	first := priv
	second := priv
	for i := 0; i < n+1; i ++ {
		second = second.Next
	}
	for second != nil {
		first = first.Next
		second = second.Next
	}
	first.Next = first.Next.Next
	return priv.Next
}


func main() {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5

	head := removeNthFromEnd(n1, 2)
	p := head
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}
}
