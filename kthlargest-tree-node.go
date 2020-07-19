package main


import "fmt"


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}



func kthLargest(root *TreeNode, k int) int {
    ret := -1

    helper(root, &k, &ret)

    return ret
}

func helper(root *TreeNode, k *int, ret *int) {
    if root.Right != nil && *k > 0 {
        helper(root.Right, k, ret)
    }

    *k -= 1
    fmt.Println("k, val", *k, root.Val)
    if *k == 0 {
        *ret = root.Val
        return
    }

    if root.Left != nil && *k > 0 {
        helper(root.Left, k, ret)
    }
}


func main() {
	n1 := &TreeNode{Val: 3}
	n2 := &TreeNode{Val: 1}
	n3 := &TreeNode{Val: 4}
	n4 := &TreeNode{Val: 2}
	n1.Left, n1.Right = n2, n3
	n2.Right = n4
	ret := kthLargest(n1, 1)
	fmt.Println(ret)
}
