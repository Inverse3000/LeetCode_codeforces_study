package main

func decodeString(s string) string {
	cur := []byte(s)
	stop := []byte{}
	stsr := []string{}
	temp := ""
	ans := ""
	for i := 0; i < len(s); {
		for cur[i] <= 'z' && cur[i] >= 'a' {
			temp = temp + string(cur[i])
			i++
		}
		if temp != "" {
			cpy := temp + ""
			stsr = append(stsr, cpy)
		}
		stop = append(stop, cur[i])
		if temp != "" && len(stop) >= 2 && stop[len(stop)-2] != '[' {
			if len(stsr) >= 2 {
				new := stsr[len(stsr)-1] + stsr[len(stsr)-2]
				stsr = stsr[:len(stsr)-1]
				stsr = stsr[:len(stsr)-1]
				stsr = append(stsr, new)
			}
		}

		if len(stop) == 1 {
			ans += temp
		} else if cur[i] == ']' {
			stop = stop[:len(stop)-1]
			stop = stop[:len(stop)-1]
			num := int(stop[len(stop)-1] - '0')
			stop = stop[:len(stop)-1]
			new := ""
			for j := 1; j <= num; j++ {
				new = new + stsr[len(stsr)-1]
			}
			stsr = stsr[:len(stsr)-1]
			if len(stsr) >= 1 {
				now := stsr[len(stsr)-1] + new
				stsr = stsr[:len(stsr)-1]
				stsr = append(stsr, now)
			} else {
				if len(stop) == 0 {
					ans = ans + new
				} else {
					stsr = append(stsr, new)
				}
			}
		}
		i++
	}
	return ans
}
func main() {

}
