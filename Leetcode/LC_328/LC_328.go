package main

import "fmt"

func main() {
	var first ListNode
	first.Val = 1
	first.Next = nil
	second := ListNode{2, nil}
	third := ListNode{3, nil}
	forth := ListNode{4, nil}
	first.Next = &second
	second.Next = &third
	third.Next = &forth
	ans := oddEvenList(&first)
	for ans != nil {
		fmt.Println(ans.Val)
		ans = ans.Next
	}
}
func oddEvenList(head *ListNode) *ListNode {
	now := head
	ecur := head.Next
	cnt := 0
	var last *ListNode
	for now != nil {
		cnt++
		if now.Next != nil && cnt%2 == 1 && now.Next.Next == nil {
			last = now
		}
		if now.Next == nil && cnt%2 == 1 {
			last = now
		}
		xy := now.Next
		if now.Next != nil {
			now.Next = now.Next.Next
		}
		now = xy
	}
	last.Next = ecur
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
