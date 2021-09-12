package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	cnt := 0
	cur := head
	var l *ListNode
	for cur != nil {
		cnt++
		if cur.Next == nil {
			l = cur
		}
		cur = cur.Next
	}
	h := head
	k = k % cnt
	for i := 1; i <= k; i++ {
		now := h
		for now.Next != nil {
			if now.Next.Next == nil {
				l = now
				now.Next.Next = h
				h = now.Next
				l.Next = nil
			}
			now = now.Next
		}
	}
	return h
}
func main() {

}
