package main

type int int64

func totalHammingDistance(nums []int) int {
	pres1 := [35]int{}
	pres0 := [35]int{}
	ans := int(0)
	for i := 0; i < len(nums); i++ {
		temp := nums[i]
		for j := 1; j <= 32; j++ {
			if temp&1 == 1 {
				ans = ans + pres0[i-1]
				pres1[i-1]++
			}
			if temp&1 == 0 {
				ans = ans + pres1[i-1]
				pres0[i-1]++
			}
			temp = temp >> 1
		}
	}
	return ans
}
func main() {

}
