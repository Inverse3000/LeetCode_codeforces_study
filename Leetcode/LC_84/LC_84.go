package main

type Node struct {
	index int
	h     int
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func largestRectangleArea(heights []int) int {
	st := []Node{}
	ans := 0
	for i := 0; i < len(heights); i++ {
		if len(st) == 0 {
			st = append(st, Node{index: i, h: heights[i]})
			continue
		}
		start := st[len(st)-1].index
		for len(st) > 0 && st[len(st)-1].h > heights[i] {
			if len(st) >= 2 {
				ans = max(ans, (start-st[len(st)-2].index)*st[len(st)-1].h)
			}
			if len(st) == 1 {
				ans = max(ans, (start)*st[len(st)-1].h)
			}
			st = st[:len(st)-1]
		}
		st = append(st, Node{i, heights[i]})
	}
	start := st[len(st)-1].index
	for len(st) > 0 {
		if len(st) >= 2 {
			ans = max(ans, (start-st[len(st)-2].index)*st[len(st)-1].h)
		}
		if len(st) == 1 {
			ans = max(ans, (start)*st[len(st)-1].h)
		}
		st = st[:len(st)-1]
	}
	return ans
}
func main() {

}
