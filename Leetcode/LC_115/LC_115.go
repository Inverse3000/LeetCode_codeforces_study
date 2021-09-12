package main

var f [5005][5005]int

func numDistinct(s string, t string) int {
	f = [5005][5005]int{}
	ns := []byte(s)
	nt := []byte(t)
	for i := 0; i < len(ns); i++ {
		for j := 0; j < len(nt); j++ {
			f[i][j] = f[i-1][j]
			if ns[i] == nt[j] {
				f[i][j] += f[i-1][j-1]
			}
		}
	}
	ans := 0
	for i := 0; i < len(ns); i++ {
		ans += f[i][len(nt)]
	}
	return ans
}
func main() {

}
