package main


import "fmt"


type ListNode struct {
    Val int
    Next *ListNode
}


func detectCycle(head *ListNode) *ListNode {
	first, second := head, head
	has_cycle := false
	for second != nil && second.Next != nil {
		first = first.Next
		second = second.Next.Next
		if first == second {
			has_cycle = true
			break
		}
	}
	if ! has_cycle {
		return nil
	}

	// 一个从起始点出发，一个从相遇点出发，必会在环入口点相遇
	p1, p2 := head, first
	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1
}


func main() {
	n1 := &ListNode{Val: 3}
	n2 := &ListNode{Val: 2}
	// n3 := &ListNode{Val: 0}
	// n4 := &ListNode{Val: -4}
	n1.Next = n2
	n2.Next = n1
	// n2.Next = n3
	// n3.Next = n4
	// n4.Next = n2

	h := detectCycle(n1)
	fmt.Println(h.Val)
}

