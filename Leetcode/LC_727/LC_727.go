package main

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
func minWindow(s1 string, s2 string) string {
	f := [20005][105]int{}
	rs1 := []byte(s1)
	rs2 := []byte(s2)
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if j == 1 && rs1[i-1] == rs2[j-1] {
				f[i][j] = i
				continue
			}
			f[i][j] = f[i-1][j]
			if rs1[i-1] == rs2[j-1] {
				f[i][j] = max(f[i][j], f[i-1][j-1])
			}
		}
	}
	ans := 999999999
	l := -1
	r := -1
	for i := 1; i <= len(s1); i++ {
		if f[i][len(s2)] != 0 {
			if ans > i-f[i][len(s2)]+1 {
				ans = i - f[i][len(s2)] + 1
				l = f[i][len(s2)]
				r = i
			}
		}
	}
	if l == -1 {
		return ""
	}
	return s1[l-1 : r]
}
func main() {

}
