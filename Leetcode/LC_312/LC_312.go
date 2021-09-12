package main

func max(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func maxCoins(nums []int) int {
	var f [505][505]int
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			f[i][i] = nums[i] * nums[i+1]
			continue
		}
		if i == len(nums)-1 {
			f[i][i] = nums[i] * nums[i-1]
			continue
		}
		f[i][i] = nums[i-1] * nums[i] * nums[i+1]
	}
	nums = append(nums, 1)
	for k := 2; k <= len(nums)-1; k++ {
		for i := 0; i+k-1 < len(nums)-1; i++ {
			for j := i; j <= i+k-1; j++ {
				/*	if k==len(nums){
					if j == 0{
						f[i][i+k-1]=max(f[i][i+k-1],nums[j]+f[j+1][i+k-1])
					}else {
						f[i][i+k-1] = max(f[i][i+k-1], nums[j]+f[j+1][i+k-1]+f[0][j-1])
					}
					continue
				}*/
				if i == 0 {
					if j == 0 {
						f[i][i+k-1] = max(f[i][i+k-1], nums[j]*nums[i+k]+f[j+1][i+k-1])
					} else {
						f[i][i+k-1] = max(f[i][i+k-1], nums[j]*nums[i+k]+f[j+1][i+k-1]+f[i][j-1])
					}
					continue
				}
				if j == 0 {
					f[i][i+k-1] = max(f[i][i+k-1], nums[j]*nums[i+k]+f[j+1][i+k-1])
				} else {
					f[i][i+k-1] = max(f[i][i+k-1], nums[j]*nums[i+k]*nums[i-1]+f[j+1][i+k-1]+f[i][j-1])
				}
			}
		}
	}
	return f[0][len(nums)-1]
}
func main() {

}
