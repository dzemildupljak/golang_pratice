// package main

// import "fmt"

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func main() {
// 	l11 := ListNode{
// 		Val: 1,
// 	}
// 	l12 := ListNode{
// 		Val: 2,
// 	}
// 	l13 := ListNode{
// 		Val: 4,
// 	}
// 	l11.Next = &l12
// 	l12.Next = &l13
// 	cur := &l11
// 	for cur != nil {
// 		fmt.Println(cur.Val)
// 		cur = cur.Next
// 	}
// 	reverseList(&l11)
// 	cur = &l13
// 	for cur != nil {
// 		fmt.Println(cur.Val)
// 		cur = cur.Next
// 	}

// }

// /**
//  * Definition for singly-linked list.
//  * type ListNode struct {
//  *     Val int
//  *     Next *ListNode
//  * }
//  */
// func reverseList(head *ListNode) *ListNode {
// 	var tmp *ListNode
// 	curr := head

// 	for curr != nil {
// 		nxt := curr.Next
// 		curr.Next = tmp
// 		tmp = curr
// 		curr = nxt
// 	}

// 	return tmp
// }
