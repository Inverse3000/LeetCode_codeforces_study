package main

import (
	"fmt"
	"sort"
)

type node struct {
	l, r, val int
	lo, ro    *node
}

func (o *node) insert(val, il, ir int) {
	m := (o.l + o.r) >> 1
	if il <= o.l && ir >= o.r {
		if o.val < val {
			o.val = val
		}
	}
	if o.l == o.r {
		return
	}
	if ir <= m {
		if o.lo == nil {
			o.lo = &node{l: o.l, r: m, val: 0}
		}
		o.lo.insert(val, il, ir)
		return
	}
	if il > m {
		if o.ro == nil {
			o.ro = &node{l: m + 1, r: o.r, val: 0}
		}
		o.ro.insert(val, il, ir)
		return
	}
	if il <= m && ir > m {
		if o.lo == nil {
			o.lo = &node{l: o.l, r: m, val: 0}
		}
		o.lo.insert(val, il, m)
		if o.ro == nil {
			o.ro = &node{l: m + 1, r: o.r, val: 0}
		}
		o.ro.insert(val, m+1, ir)
		return
	}
}
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
func (o *node) query(l, r int) (ans int) {
	if o == nil || l > o.r || r < o.l {
		return 0
	}
	if l <= o.l && r >= o.r {
		return o.val
	}
	return max(o.lo.query(l, r), o.ro.query(l, r))
}
func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
func getSkyline(buildings [][]int) [][]int {
	ans := [][]int{}
	st := []int{}
	rankx := 1
	for i := 0; i < len(buildings); i++ {
		st = append(st, buildings[i][0])
		st = append(st, buildings[i][1])
		st = append(st, buildings[i][1]-1)
	}
	unique(st)
	sort.Ints(st)
	mp := map[int]int{}
	rmp := map[int]int{}
	for i := 0; i < len(st); i++ {
		mp[st[i]] = rankx
		rmp[rankx] = st[i]
		rankx++
	}
	rankx -= 1
	o := &node{l: 1, r: rankx}
	for i := 0; i < len(buildings); i++ {
		o.insert(buildings[i][2], mp[buildings[i][0]], mp[buildings[i][1]-1])
	}
	pre := -1
	for i := 1; i <= rankx; i++ {
		cool := o.query(i, i)
		fmt.Println(cool)
		if cool != pre {
			pre = cool
			ans = append(ans, []int{rmp[i], cool})
		}
	}
	return ans
}
func main() {
}
