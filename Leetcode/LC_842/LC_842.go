package main

var ans []int
var end []int
var flag bool

func search(s int, nums []byte) {
	if flag {
		return
	}
	if s == len(nums) && len(ans) < 3 {
		return
	}
	if s == len(nums) && len(ans) >= 3 {
		end = make([]int, len(ans))
		copy(end, ans)
		flag = true
		return
	}
	for i := s; i < len(nums); i++ {
		cnt := 1
		temp := 0
		for j := s; j <= i; j++ {
			temp = temp*10 + int(nums[i]-'0')
			cnt *= 10
		}
		ans = append(ans, temp)
		if len(ans) == 1 || len(ans) == 2 {
			search(i+1, nums)
			ans = ans[:len(ans)-1]
			continue
		}
		if ans[len(ans)-2]+ans[len(ans)-3] == ans[len(ans)-1] {
			search(i+1, nums)
			ans = ans[:len(ans)-1]
		} else if ans[len(ans)-2]+ans[len(ans)-3] < ans[len(ans)-1] {
			ans = ans[:len(ans)-1]
			break
		}
		ans = ans[:len(ans)-1]
	}
}
func splitIntoFibonacci(num string) []int {
	flag = false
	ans = []int{}
	nums := []byte(num)
	search(0, nums)
	return end
}
func main() {

}
