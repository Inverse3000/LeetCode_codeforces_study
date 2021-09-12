package main

func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	ans := -1
	lastmid := -1
	for l <= r {
		mid := (l + r) / 2
		if lastmid != -1 && nums[lastmid] > nums[mid] && lastmid < mid {
			r = mid - 1
			lastmid = mid
			continue
		}
		if lastmid != -1 && nums[lastmid] < nums[mid] && lastmid > mid {
			l = mid + 1
			lastmid = mid
			continue
		}
		if nums[mid] < target {
			l = mid + 1
			lastmid = mid
			continue
		}
		if nums[mid] > target {
			r = mid - 1
			lastmid = mid
			continue
		}
		if nums[mid] == target {
			ans = mid
			break
		}
	}
	if l < len(nums)/2 {
		l = len(nums)/2 + 1
		r = len(nums) - 1
	} else {
		l = 0
		r = len(nums)/2 - 1
	}
	for l <= r {
		mid := (l + r) / 2
		if lastmid != -1 && nums[lastmid] > nums[mid] && lastmid < mid {
			r = mid - 1
			lastmid = mid
			continue
		}
		if lastmid != -1 && nums[lastmid] < nums[mid] && lastmid > mid {
			l = mid + 1
			lastmid = mid
			continue
		}
		if nums[mid] < target {
			l = mid + 1
			lastmid = mid
			continue
		}
		if nums[mid] > target {
			r = mid - 1
			lastmid = mid
			continue
		}
		if nums[mid] == target {
			ans = mid
			break
		}
	}
	return ans
}
func main() {

}
