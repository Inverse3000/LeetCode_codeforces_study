package main

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	if k == 0 {
		return 0
	}
	b := []byte(s)
	mp := map[byte]int{}
	cnt := 0
	l := 0
	ans := -1
	for i := 0; i < len(b); i++ {
		if mp[b[i]] != 0 {
			mp[b[i]]++
		} else {
			for cnt == k && l < i {
				mp[b[l]]--
				if mp[b[l]] == 0 {
					cnt--
				}
				l++
			}
			mp[b[i]] = 1
			cnt++
		}
		if ans < i-l+1 {
			ans = i - l + 1
		}
	}
	return ans
}
func main() {

}
