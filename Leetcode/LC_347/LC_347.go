package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	mp := map[int]int{}
	delta := [50005][]int{}
	for i := 0; i < len(nums); i++ {
		mp[nums[i]]++
	}
	for k, v := range mp {
		fmt.Println(k, v)
		delta[v] = append(delta[v], k)
	}
	ans := []int{}
	cnt := 0
	for i := 10005; i >= 0; i-- {
		for len(delta[i]) != 0 && cnt < k {
			cnt++
			ans = append(ans, delta[i][0])
			delta[i] = delta[i][1:]
		}
	}
	return ans
}
func main() {

}
