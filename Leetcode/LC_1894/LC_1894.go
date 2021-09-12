package main

func chalkReplacer(chalk []int, k int) int {
	n := 0
	for i := 1; i <= len(chalk); i++ {
		n += chalk[i]
	}
	m := k % n
	for i := 1; i <= len(chalk); i++ {
		if m < chalk[i-1] {
			return i - 1
		}
		m = m - chalk[i]
	}
	return -1
}
func main() {

}
