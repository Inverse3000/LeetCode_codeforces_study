package LC_406

import "sort"

type Node struct {
	x int
	y int
}
type Nodes []Node

func (k Nodes) Len() int {
	return len(k)
}
func (k Nodes) Less(i, j int) bool {
	return k[i].x > k[j].x
}
func (k Nodes) Swap(i, j int) {
	k[i], k[j] = k[j], k[i]
}
func reconstructQueue(people [][]int) [][]int {
	a := Nodes{}
	for i := 0; i < len(people); i++ {
		a = append(a, Node{people[i][0], people[i][1]})
	}
	sort.Sort(a)
	var ans [1005][2]int
	var cnt int
	for i := 0; i < len(people); i++ {
		cnt = 0
		for j := 0; j < len(people); j++ {
			if cnt == a[i].y {
				ans[j][0] = a[i].x
				ans[j][1] = a[i].y
				break
			}
			if ans[j][0] == 0 {
				cnt++
			}
		}
	}
	var cool [][]int
	for i := 0; i < len(people); i++ {
		cool = append(cool, []int{ans[i][0], ans[i][1]})
	}
	return cool
}
