package main

func wordBreak(s string, wordDict []string) bool {
	b := []byte(s)
	var f [300005]int
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(wordDict); j++ {
			if i-len(wordDict[j])+1 >= 0 && string(b[i-len(wordDict[j])+1:i+1]) == wordDict[j] && (i-len(wordDict[j]) == -1 || f[i-len(wordDict[j])] == 1) {
				f[i] = 1
			}
		}
	}
	if f[len(b)-1] == 1 {
		return true
	}
	return false
}
func main() {

}
