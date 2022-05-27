package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l11 := ListNode{
		Val: 1,
	}
	l12 := ListNode{
		Val: 2,
	}
	l13 := ListNode{
		Val: 4,
	}
	l11.Next = &l12
	l12.Next = &l13
	l21 := ListNode{
		Val: 1,
	}
	l22 := ListNode{
		Val: 3,
	}
	l23 := ListNode{
		Val: 4,
	}
	l21.Next = &l22
	l22.Next = &l23
	resLl := mergeTwoLists(&l11, &l21)

	for resLl != nil {
		fmt.Println(resLl.Val)
		resLl = resLl.Next
	}

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoLists(list1, list2.Next)
	return list2
}

// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
// 	var head = new(ListNode)
// 	var ll = head

// 	for list1 != nil && list2 != nil {
// 		if list1.Val <= list2.Val {
// 			ll.Next = list1
// 			list1 = list1.Next
// 		} else {
// 			ll.Next = list2
// 			list2 = list2.Next
// 		}
// 		ll = ll.Next
// 	}
// 	if list1 != nil {
// 		ll.Next = list1
// 	} else {
// 		ll.Next = list2
// 	}

// 	return head.Next
// }
