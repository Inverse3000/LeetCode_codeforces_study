package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	cur := head
	mp := map[*ListNode]int{}
	for cur != nil {
		if mp[cur] != 0 {
			return cur
		}
		mp[cur] = 1
		cur = cur.Next
	}
	return nil
}
func main() {

}
