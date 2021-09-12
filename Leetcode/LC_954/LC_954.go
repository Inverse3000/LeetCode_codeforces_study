package main

import "sort"

func canReorderDoubled(arr []int) bool {
	ans := 0
	cnt := map[int]int{}
	sort.Ints(arr)
	index := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] >= 0 {
			break
		}
		index++
	}
	for i := len(arr) - 1; i >= index+1; i-- {
		if cnt[arr[i]*2] == 0 {
			cnt[arr[i]]++
		} else {
			cnt[arr[i]*2]--
			ans += 2
		}
	}
	for i := 0; i <= index; i++ {
		if cnt[arr[i]*2] == 0 {
			cnt[arr[i]]++
		} else {
			cnt[arr[i]*2]--
			ans += 2
		}
	}
	if ans == len(arr) {
		return true
	}
	return false
}
func main() {

}
