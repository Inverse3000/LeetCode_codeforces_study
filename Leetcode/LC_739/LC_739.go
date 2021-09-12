package main

type Node struct {
	t     int
	index int
}

var ans [300005]int

func dailyTemperatures(temperatures []int) []int {
	ans = [300005]int{}
	st := []Node{}
	for i := 0; i < len(temperatures); i++ {
		if len(st) == 0 {
			st = append(st, Node{temperatures[i], i})
			continue
		}
		for len(st) > 0 && st[len(st)-1].t < temperatures[i] {
			ans[st[len(st)-1].index] = i - st[len(st)-1].index
			st = st[:len(st)-1]
		}
		st = append(st, Node{temperatures[i], i})
	}
	for len(st) > 0 {
		ans[st[len(st)-1].index] = 0
		st = st[:len(st)-1]
	}
	ret := []int{}
	for i := 0; i < len(temperatures); i++ {
		ret = append(ret, ans[i])
	}
	return ret
}
func main() {

}
