package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	length := 0
	node := head
	for node != nil {
		length++
		node = node.Next
	}
	dummyHead := &ListNode{0, head}
	for sublen := 1; sublen < length; sublen *= 2 {
		prev := dummyHead
		cur := dummyHead.Next
		for cur != nil {
			head1 := cur
			for i := 1; i < sublen && cur.Next != nil; i++ {
				cur = cur.Next
			}
			head2 := cur.Next
			cur.Next = nil
			cur = head2
			for i := 1; i < sublen && cur != nil && cur.Next != nil; i++ {
				cur = cur.Next
			}
			var next *ListNode
			next = nil
			if cur != nil {
				next = cur.Next
				cur.Next = nil
			}
			merged := merge(head1, head2)
			prev.Next = merged
			for prev.Next != nil {
				prev = prev.Next
			}
			cur = next
		}
	}
	return dummyHead.Next
}
func merge(head1 *ListNode, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{0, nil}
	temp := dummyHead
	temp1 := head1
	temp2 := head2
	for temp1 != nil && temp2 != nil {
		if temp1.Val < temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}
	if temp1 != nil {
		temp.Next = temp1
	} else {
		temp.Next = temp2
	}
	return dummyHead.Next
}
func main() {

}
