package main

var ans []string
var b []byte

func search(l, r, n int) {
	if l > n {
		return
	}
	if l == n && r == n {
		temp := make([]byte, len(b))
		copy(temp, b)
		ans = append(ans, string(b))
		return
	}
	if l == r {
		b = append(b, '(')
		search(l+1, r, n)
		b = b[:len(b)-1]
	}
	if l > r {
		b = append(b, '(')
		search(l+1, r, n)
		b = b[:len(b)-1]
		b = append(b, ')')
		search(l, r+1, n)
		b = b[:len(b)-1]
	}
}
func generateParenthesis(n int) []string {
	search(0, 0, n)
	return ans
}
func main() {

}
