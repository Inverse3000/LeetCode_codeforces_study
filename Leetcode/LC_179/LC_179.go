package main

import (
	"strconv"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func largestNumber(nums []int) string {
	var a string
	var b string
	var ma string
	var ans string
	ans = ""
	var vis [105]int
	lst := -1
	for k := 1; k <= len(nums); k++ {
		ma = ""
		for i := 0; i < len(nums); i++ {
			if vis[i] != 0 {
				continue
			}
			a = ma + strconv.Itoa(nums[i])
			b = strconv.Itoa(nums[i]) + ma
			if a <= b {
				ma = strconv.Itoa(nums[i])
				lst = i
			}
		}
		vis[lst] = 1
		if ans == "0" {
			ans = ""
		}
		ans = ans + ma
	}
	return ans
}
func main() {

}
