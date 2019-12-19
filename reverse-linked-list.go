package main
import "fmt"


type ListNode struct {
    Val int
    Next *ListNode
}


func reverseList(head *ListNode) *ListNode {
	var priv *ListNode  // 定义一个空指针
	cur := head
	for cur != nil {
		cur.Next, priv, cur = priv, cur, cur.Next
	}
	return priv
}


func main() {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	head := reverseList(n1)
	p := head
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}
}
