package main

import (
	"fmt"
)

// ListNode - узел связного список.
type ListNode struct {
	Next *ListNode
	Val  int
}

// возвращает развернутый связный список
func (head *ListNode) Reverse() *ListNode {
	if head == nil {
		return nil
	}

	var r *ListNode

	curr := head
	for curr != nil {
		r = &ListNode{
			Next: r,
			Val:  curr.Val,
		}
		curr = curr.Next
	}

	return r
}

func main() {

	// связный список вида: 10 -> 20
	list := ListNode{
		Next: &ListNode{
			Next: nil,
			Val:  20,
		},
		Val: 10,
	}

	reversed := list.Reverse() // 20 -> 10

	fmt.Println("reversed")
	var current *ListNode = reversed
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
	// reversed
	// 20
	// 10


	fmt.Println("list")
	current = &list
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
	// list
	// 10
	// 20
	// то есть исходный список не изменился

}