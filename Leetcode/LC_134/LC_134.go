package main

func canCompleteCircuit(gas []int, cost []int) int {
	temp := 0
	for i := 0; i < len(gas); i++ {
		temp += (gas[i] - cost[i])
	}
	if temp < 0 {
		return -1
	}
	sum := 0
	ans := -1
	for i := 0; i < len(gas); i++ {
		sum += gas[i] - cost[i]
		if sum < 0 {
			sum = 0
			ans = -1
		}
		if sum > 0 && ans != -1 {
			ans = i
		}
	}
	return ans
}
func main() {

}
