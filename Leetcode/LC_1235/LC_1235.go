package main

import (
	"fmt"
	"sort"
)

type Node struct {
	x int
	y int
	p int
}
type Nodes []Node

func (u Nodes) Less(i, j int) bool {
	return u[i].x < u[j].x
}
func (u Nodes) Len() int {
	return len(u)
}
func (u Nodes) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}
func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
func jobScheduling(startTime []int, endTime []int, profit []int) int {
	a := Nodes{}
	for i := 0; i < len(startTime); i++ {
		a = append(a, Node{startTime[i], endTime[i], profit[i]})
	}
	sort.Sort(a)
	fmt.Println(a)
	f := [100005]int{}
	n := len(startTime)
	f[n-1] = a[n-1].p
	for i := n - 2; i >= 0; i-- {
		l, r := 0, n-1
		ans := -1
		for l <= r {
			mid := (l + r) / 2
			if i == 0 {
				fmt.Println(mid, i, a[i].y, a[mid].x)
			}
			if a[mid].x >= a[i].y {
				ans = mid
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		if i == 0 {
			fmt.Println(ans)
		}
		if ans == -1 {
			f[i] = a[i].p
		} else {
			for j := ans; j < n; j++ {
				if a[j].x > a[ans].y {
					break
				}
				f[i] = max(f[i], f[j]+a[i].p)
			}
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, f[i])
	}
	return ans
}
func main() {
	st := []int{6, 15, 7, 11, 1, 3, 16, 2}
	et := []int{19, 18, 19, 16, 10, 8, 19, 8}
	pf := []int{2, 9, 1, 19, 5, 7, 3, 19}
	fmt.Println(jobScheduling(st, et, pf))
}
