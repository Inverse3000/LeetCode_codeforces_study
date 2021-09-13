package main

func numberOfBoomerangs(points [][]int) int {
	ans := 0
	for i := 0; i < len(points); i++ {
		v := map[int]int{}
		for j := 0; j < len(points); j++ {
			if i == j {
				continue
			}
			now := (points[i][0]-points[j][0])*(points[i][0]-points[j][0]) + (points[i][1]-points[j][1])*(points[i][1]-points[j][1])
			if v[now] != 0 {
				ans += v[now] * 2
				v[now]++
			} else {
				v[now]++
			}
		}
	}
	return ans
}
func main() {

}
