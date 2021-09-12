package main

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
func stoneGameVIII(stones []int) int {
	pre := [100005]int{}
	f := [100005]int{}
	for i := 1; i <= len(stones); i++ {
		pre[i] = pre[i-1] + stones[i]
	}
	n := len(stones)
	f[n-1] = pre[n]
	for i := n - 2; i >= 1; i-- {
		f[i] = max(f[i+1], pre[i+1]-f[i+1])
	}
	return f[1]
}
func main() {

}
