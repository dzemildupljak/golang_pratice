package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

// func middleNode(head *ListNode) *ListNode {
// 	currNode := head
// 	counter := 0

// 	for currNode.Next != nil {
// 		counter++
// 		currNode = currNode.Next
// 	}
// 	if counter%2 == 0 {
// 		counter = counter / 2
// 	} else {
// 		counter = counter/2 + 1
// 	}
// 	currNode = head

// 	for i := 0; i < counter; i++ {
// 		currNode = currNode.Next
// 	}

// 	return currNode
// }
