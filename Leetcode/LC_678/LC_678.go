package main

func checkValidString(s string) bool {
	rs := []byte(s)
	f := [205][205]bool{}
	for i := 0; i < len(rs); i++ {
		if rs[i] == '*' {
			f[i][i] = true
		} else {
			f[i][i] = false
		}
	}
	for k := 2; k <= len(rs); k++ {
		for i := 0; i+k-1 < len(rs); i++ {
			f[i][i+k-1] = false
			if (rs[i] == '*' || s[i] == '(') && (rs[i+k-1] == ')' || rs[i+k-1] == '*') && (k == 2 || f[i+1][i+k-2]) {
				f[i][i+k-1] = true
			}
			for j := i; j < i+k-1; j++ {
				if f[i][j] && f[j+1][i+k-1] {
					f[i][i+k-1] = true
				}
			}
		}
	}
	return f[0][len(rs)-1]
}
func main() {

}
