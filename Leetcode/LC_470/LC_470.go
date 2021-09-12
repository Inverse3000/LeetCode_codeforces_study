package main

func rand7() int {
	return 1
}

var cnt = map[int]int{}
var sum = 0

func rand10() int {
	ans := 0
	now := rand7()*2 - 4
	if sum == 0 {
		if now < 1 {
			now = 1
		}
		if now > 10 {
			now = 10
		}
		ans = now
		cnt[now]++
	}
	if sum != 0 && float32(cnt[now]/sum) <= float32(1/10) {
		ans = now
	} else {
		for i := 1; i <= 10; i++ {
			if sum != 0 && float32(cnt[now]/sum) <= float32(1/10) {
				ans = now
				break
			}
		}
	}
	sum++
	return ans
}
func main() {

}
