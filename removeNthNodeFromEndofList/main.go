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
		Val: 3,
	}
	l14 := ListNode{
		Val: 4,
	}
	l15 := ListNode{
		Val: 5,
	}
	l11.Next = &l12
	l12.Next = &l13
	l13.Next = &l14
	l14.Next = &l15
	removeNthFromEnd(&l11, 2)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	ctr := 0
	curr := head

	for curr != nil {
		ctr++
		curr = curr.Next
	}
	fmt.Println(ctr)

	for i := 0; i < ctr-n; i++ {
		fmt.Println(i)
	}
	return head
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//  [1,2,3,4,5]
// func removeNthFromEnd(head *ListNode, n int) *ListNode {
// 	i, listRet, temp1 := 0, head, head
// 	for ; i < n && temp1 != nil; i++ {
// 		temp1 = temp1.Next
// 	}

// 	if i < n {
// 		listRet = nil
// 	} else if temp1 == nil {
// 		listRet = head.Next
// 	} else {
// 		preNth := head
// 		for temp1.Next != nil {
// 			preNth, temp1 = preNth.Next, temp1.Next
// 		}
// 		preNth.Next = preNth.Next.Next
// 	}

// 	return listRet
// }

// func removeNthFromEnd(head *ListNode, n int) *ListNode {
// 	currNode := head
// 	counter := 0

// 	for currNode.Next != nil {
// 		counter++
// 		currNode = currNode.Next
// 	}
// 	rmNth := counter - n
// 	currNode = head

// 	if rmNth < counter {
// 		if rmNth == 0 {
// 			head = nil
// 		} else if n == 1{

// 		}
// 		else {
// 			for i := 0; i < rmNth; i++ {
// 				currNode = currNode.Next
// 			}
// 			currNode.Next = currNode.Next.Next
// 		}
// 	}

// 	return head

// }
