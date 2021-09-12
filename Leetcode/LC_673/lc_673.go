package main

import "sort"

var n int
var leng []int
var cnt []int

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
func lowbit(x int) int {
	return x & (-x)
}
func add(x, l, c int) {
	for i := x; i <= n; i += lowbit(i) {
		if leng[i] == l {
			cnt[i] += c
		} else if leng[i] < l {
			leng[i] = l
			cnt[i] = c
		}
	}
}
func query(x int) (l, c int) {
	for i := x; i >= 1; i -= lowbit(i) {
		if leng[i] > l {
			c = cnt[i]
			l = leng[i]
		} else if leng[i] == l {
			c += cnt[i]
		}
	}
	return l, c
}
func findNumberOfLIS(nums []int) int {
	n = len(nums)
	lol := make([]int, len(nums))
	copy(lol, nums)
	lol = unique(lol)
	sort.Ints(lol)
	if len(nums) <= 1 {
		return len(nums)
	}
	mp := map[int]int{}
	for i := 1; i <= n; i++ {
		mp[lol[i-1]] = i
	}
	for _, v := range nums {
		now := mp[v]
		l, c := query(now - 1)
		add(now, l+1, c)
	}
	_, ans := query(n + 1)
	return ans
}
func main() {

}
