package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	Node := head
	var begin *ListNode
	var end *ListNode
	cnt := 0
	for Node != nil {
		if cnt == left {
			begin = Node
		}
		if cnt == right {
			end = Node
		}
		cnt++
	}
	var before *ListNode
	if begin == head {
		before = nil
	} else {
		Node = head
		for Node != nil {
			if Node.Next == begin {
				before = Node
				break
			}
		}
	}
	after := end.Next
	var prev *ListNode
	var cur *ListNode
	var nex *ListNode
	prev = begin
	cur = begin
	nex = cur.Next
	for i := 1; i <= right-left+1; i++ {
		cur.Next = nex.Next
		nex.Next = prev
		prev = nex
		nex = cur.Next
	}
	before.Next = prev
	cur.Next = after
	if begin == head {
		return end
	}
	return head
}
func main() {

}
