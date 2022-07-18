package main

import (
	"github.com/isdamir/gotype"
)

func Reverse(node *gotype.LNode) {
	if node == nil || node.Next == nil {
		return
	}

	var pre, cur *gotype.LNode // 定义前驱结点 当前结点
	next := node.Next

	for next != nil {
		cur = next.Next
		next.Next = pre
		pre = next
		next = cur
	}
	node.Next = pre
}

func main() {
	head := &gotype.LNode{}
	gotype.CreateNode(head, 8)
	gotype.PrintNode("before: ", head)
	Reverse(head)
	gotype.PrintNode("after: ", head)
}