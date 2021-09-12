package main

var fa [300005]int

func findf(x int) int {
	if x != fa[x] {
		fa[x] = findf(fa[x])
	}
	return fa[x]
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
func longestConsecutive(nums []int) int {
	nums = unique(nums)
	n := len(nums)
	mp := map[int]int{}
	for i := 0; i < n; i++ {
		mp[nums[i]] = i + 1
	}
	for i := 1; i <= n; i++ {
		fa[i] = i
	}
	lp := map[int]int{}
	for i := 0; i < n; i++ {
		if mp[nums[i]-1] != 0 {
			nowx := mp[nums[i]-1]
			nowy := mp[nums[i]]
			fx := findf(nowx)
			fy := findf(nowy)
			if fx != fy {
				fa[fx] = fy
			}
		}
		if mp[nums[i]+1] != 0 {
			nowx := mp[nums[i]+1]
			nowy := mp[nums[i]]
			fx := findf(nowx)
			fy := findf(nowy)
			if fx != fy {
				fa[fx] = fy
			}
		}
	}
	ans := 0
	for i := 0; i < len(nums); i++ {
		lp[findf(mp[nums[i]])]++
		if ans < lp[findf(mp[nums[i]])] {
			ans = lp[findf(mp[nums[i]])]
		}
	}
	return ans
}
func main() {

}
