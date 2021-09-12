package main

import (
	"fmt"
	"math"
	"sort"
)

func nthUglyNumber(n int) int {
	f := []int{1, 2, 3, 4, 5}
	mp := map[int]int{}
	mp[1] = 1
	mp[2] = 1
	mp[3] = 1
	mp[4] = 1
	mp[5] = 1
	cnt := 5
	for cnt <= n {
		temp := []int{}
		for i := 1; i < len(f); i++ {
			if mp[f[i]*2] == 0 {
				mp[f[i]*2] = 1
				temp = append(temp, f[i]*2)
			}
			if mp[f[i]*3] == 0 {
				mp[f[i]*3] = 1
				temp = append(temp, f[i]*3)
			}
			if mp[f[i]*5] == 0 {
				mp[f[i]*5] = 1
				temp = append(temp, f[i]*5)
			}
		}
		for i := 0; i < len(temp); i++ {
			f = append(f, temp[i])
		}
		cnt += len(temp)
	}
	sort.Ints(f)
	return f[n-1]
}
func main() {
	fmt.Println(1399680000 - math.MaxInt32)
}
