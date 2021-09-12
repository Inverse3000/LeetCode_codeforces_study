package main

import "sort"

var f [50005]int
var l [1005][]int

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func largestDivisibleSubset(nums []int) []int {
	f = [50005]int{}
	sort.Ints(nums)
	ans := []int{}
	for i := 0; i < len(nums); i++ {
		f[i] = 1
		l[i] = append(l[i], nums[i])
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 {
				if f[i] < f[j]+1 {
					l[i] = append(l[j], nums[i])
				}
				f[i] = max(f[i], f[j]+1)
			}
		}
	}
	ma := 0
	for i := 0; i < len(nums); i++ {
		if f[i] < ma {
			ma = f[i]
			ans = l[i]
		}
	}
	return ans
}
func main() {

}
