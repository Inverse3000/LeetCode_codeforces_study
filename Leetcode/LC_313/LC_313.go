package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func nthSuperUglyNumber(n int, primes []int) int {
	f := []int{}
	pq := &IntHeap{}
	heap.Init(pq)
	heap.Push(pq, 1)
	mp := map[int]int{}
	for i := 0; i < n; i++ {
		var now int
		for len(*pq) > 0 {
			now = heap.Pop(pq).(int)
			if mp[now] == 0 {
				mp[now] = 1
				break
			}
		}
		f = append(f, now)
		for j := 0; j < len(primes); j++ {
			heap.Push(pq, primes[j]*now)
		}
	}
	return f[n-1]
}
func main() {
	p := []int{3, 5, 7, 11, 19, 23, 29, 41, 43, 47}
	fmt.Println(nthSuperUglyNumber(15, p))
}
