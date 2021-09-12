package main

import (
	"container/heap"
)

func main() {
	temp := []int{6, 7, 4, 5}
	maxSlidingWindow(temp, 3)
}

type Node struct {
	index int
	val   int
}
type Nodes []Node

func (pq *Nodes) Len() int {
	return len(*pq)
}
func (pq *Nodes) Less(i, j int) bool {
	return (*pq)[i].val > (*pq)[j].val
}
func (pq *Nodes) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}
func (pq *Nodes) Push(x interface{}) {
	*pq = append(*pq, x.(Node))
}
func (pq *Nodes) Pop() interface{} {
	n := len(*pq)
	x := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return x
}
func maxSlidingWindow(nums []int, k int) []int {
	cnt := 0
	p := &Nodes{}
	heap.Init(p)
	for i, v := range nums {
		cnt++
		heap.Push(p, Node{i, v})
		if cnt == k {
			break
		}
	}
	ans := []int{}
	ans = append(ans, (*p)[len(*p)-1].val)
	l := 0
	for i := k; i < len(nums); i++ {
		l++
		heap.Push(p, Node{i, nums[i]})
		for len(*p) != 0 {
			now := heap.Pop(p).(Node)
			if now.index < l {
				continue
			} else {
				ans = append(ans, now.val)
				heap.Push(p, now)
				break
			}
		}
	}
	return ans
}
