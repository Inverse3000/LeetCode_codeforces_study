package main

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func jump(nums []int) int {
	var f [10005]int
	for i := 1; i <= 10000; i++ {
		f[i] = 999999
	}
	f[0] = 0
	for i := 0; i < len(nums); i++ {
		for j := 1; j <= nums[i]; j++ {
			f[i+j] = min(f[i+j], f[i]+1)
		}
	}
	return f[len(nums)-1]
}
func main() {

}
