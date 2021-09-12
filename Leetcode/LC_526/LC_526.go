package main

var ans int
var vis [20]int
var cnt [20]int

func search(a, b, n int) {
	if cnt[a]%a != 0 && a%cnt[a] != 0 && a != 0 {
		return
	}
	if a == n {
		if cnt[a]%a == 0 || a%cnt[a] == 0 {
			ans++
		}
		return
	}
	for i := 1; i <= n; i++ {
		if vis[i] == 0 {
			vis[i] = 1
			search(a+1, i, n)
			vis[i] = 0
		}
	}

}
func countArrangement(n int) int {
	for i := 0; i < len(vis); i++ {
		vis[i] = 0
	}
	ans = 0
	search(0, 0, n)
	return ans
}
func main() {

}
