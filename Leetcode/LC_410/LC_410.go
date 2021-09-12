package main

func check(x, m int, nums []int) bool {
	cnt := 1
	temp := 0
	for i := 0; i < len(nums); i++ {
		if temp+nums[i] > x {
			cnt++
			temp = 0
		}
		temp += nums[i]
	}
	if cnt > m {
		return false
	}
	return true
}
func splitArray(nums []int, m int) int {
	l := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	r := sum
	ans := 0
	for l <= r {
		mid := (l + r) / 2
		if check(mid, m, nums) {
			ans = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return ans
}
func main() {

}
