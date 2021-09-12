package main

import "sort"

func hIndex(citations []int) int {
	sort.Ints(citations)
	cnt := 0
	ans := 0
	for i := len(citations) - 1; i >= 0; i-- {
		cnt++
		if citations[i] >= cnt && i-1 >= 0 && citations[i-1] <= cnt {
			if ans < cnt {
				ans = cnt
			}
		}
	}
	return ans
}
func main() {

}
