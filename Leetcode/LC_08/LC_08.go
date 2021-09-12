package main

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {
	b := []byte(s)
	ans := int64(0)
	l := int64(10)
	start := 0
	fmt.Println(len(b))
	for i := 0; i < len(b); i++ {
		if b[i] == ' ' {
			continue
		} else if b[i] >= '0' && b[i] <= '9' {
			fmt.Println(i, "?")
			start = i
			break
		}
	}
	m := 1
	if start-1 >= 0 && b[start-1] == '+' {
		m = 1
	} else if start-1 >= 0 && b[start-1] == '-' {
		m = -1
	}
	flag := false
	fmt.Println(start)
	for i := start; i < len(b); i++ {
		if b[i] < '0' || b[i] > '9' {
			break
		}
		if (ans*l+int64(b[i]-'0') > math.MaxInt32 && m == 1) || (ans*l+int64(b[i]-'0') > math.MaxInt32+1 && m == -1) {
			flag = true
		}
		fmt.Println(ans)
		ans = ans*l + int64(b[i]-'0')
	}
	ans *= int64(m)
	fmt.Println(flag, m)
	if flag && m == 1 {
		return math.MaxInt32
	} else if flag && m == -1 {
		return -math.MaxInt32 - 1
	}
	return int(ans)
}
func main() {
	fmt.Println(myAtoi("-91283472332"))
}
