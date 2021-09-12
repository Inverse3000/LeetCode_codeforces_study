package main

import "fmt"

var c [30005][26]int
var val [30005]int
var fail [300005]int
var cnt = 0
var q []int

func buildT(s []byte) {
	now := 0
	for i := 0; i < len(s); i++ {
		v := int(s[i] - 'a')
		if c[now][v] == 0 {
			cnt++
			c[now][v] = cnt
		}
		now = c[now][v]
	}
	val[now]++
}
func buildF() {
	for i := 0; i < 26; i++ {
		if c[0][i] != 0 {
			fail[c[0][i]] = 0
			q = append(q, c[0][i])
		}
	}
	for len(q) != 0 {
		fmt.Println(len(q))
		u := q[0]
		q = q[1:]
		for i := 0; i < 26; i++ {
			if c[u][i] != 0 {
				fail[c[u][i]] = c[fail[u]][i]
				q = append(q, c[u][i])
			} else {
				c[u][i] = c[fail[u]][i]
			}
		}
	}
}

type StreamChecker struct {
}

func Constructor(words []string) StreamChecker {
	for i := 0; i < len(words); i++ {
		temp := []byte(words[i])
		buildT(temp)
	}
	buildF()
	var m StreamChecker
	return m
}

var now = 0

func (this *StreamChecker) Query(letter byte) bool {
	next := int(letter - 'a')
	now = c[now][next]
	if val[now] != 0 {
		return true
	}
	return false
}
func main() {
	streamChecker := Constructor([]string{"ab", "ba", "aaab", "abab", "baa"})
	streamChecker.Query('a')
}
