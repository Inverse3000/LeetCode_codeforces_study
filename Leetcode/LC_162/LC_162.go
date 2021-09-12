package main

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			if nums[i+1] < nums[i] {
				return 0
			}
			continue
		}
		if i == len(nums)-1 {
			if nums[i-1] < nums[i] {
				return i
			}
			continue
		}
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			return i
		}
	}
	return -1
}
func main() {

}
