package main

var ans [][]int
var cur []int

func search(a int, c []int, t int, mi int) {
	if a > t {
		return
	}
	if a == t {
		temp := make([]int, len(cur))
		copy(temp, cur)
		ans = append(ans, temp)
		return
	}
	for i := 0; i < len(c); i++ {
		if c[i] >= mi {
			cur = append(cur, c[i])
			search(a+c[i], c, t, c[i])
			cur = cur[:len(cur)-1]
		}
	}
}
func combinationSum(candidates []int, target int) [][]int {
	ans = [][]int{}
	cur = []int{}
	search(0, candidates, target, 0)
	return ans
}
func main() {

}
