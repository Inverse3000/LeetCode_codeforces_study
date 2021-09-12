package main

var btw [][]int
var cur []int
var n int

func search(a int, nums []int) {
	if a == n {
		temp := make([]int, len(cur))
		copy(temp, cur)
		btw = append(btw, temp)
	}
	cur = append(cur, nums[a])
	search(a+1, nums)
	cur = cur[:len(cur)-1]
	search(a+1, nums)
}
func subsets(nums []int) [][]int {
	n = len(nums)
	cur = []int{}
	btw = [][]int{}
	search(0, nums)
	return btw
}
func main() {

}
