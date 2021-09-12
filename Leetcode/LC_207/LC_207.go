package main

var u [50005]int
var v [50005]int
var tot int
var next [50005]int
var first [50005]int
var vis [50005]int
var no [50005]int

func add(x, y int) {
	tot++
	u[tot] = x
	v[tot] = y
	next[tot] = first[x]
	first[x] = tot
}
func search(x int) {
	vis[x] = 1
	for i := first[x]; i != 0; i = next[i] {
		if vis[v[i]] == 0 {
			search(v[i])
		}
	}
}
func canFinish(numsCourses int, prerequisites [][]int) bool {
	u = [50005]int{}
	v = [50005]int{}
	tot = 0
	next = [50005]int{}
	first = [50005]int{}
	vis = [50005]int{}
	for i := 0; i < len(prerequisites); i++ {
		add(prerequisites[i][0], prerequisites[i][1])
		no[prerequisites[i][1]] = 1
	}
	for i := 0; i < numsCourses; i++ {
		if vis[i] == 0 {
			search(i)
		}
	}
	for i := 0; i < numsCourses; i++ {
		if vis[i] == 0 && no[i] == 0 {
			return false
		}
	}
	return true
}
func main() {

}
