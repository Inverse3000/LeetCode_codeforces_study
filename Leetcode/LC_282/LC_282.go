package main

import (
	"strconv"
)

var prev int
var sum int
var cur int
var nums []byte
var exp string
var t int
var ans []string

func search(index, op int) {
	if index == len(nums) {
		osum := sum
		oexp := exp
		if op == 0 {
			exp = exp + "+"
			exp = exp + strconv.Itoa(cur)
			sum = sum + cur
		}
		if op == 1 {
			exp = exp + "-"
			exp = exp + strconv.Itoa(cur)
			sum = sum - cur
		}
		if op == 2 {
			exp = exp + "*"
			exp = exp + strconv.Itoa(cur)
			sum -= prev
			sum = sum + prev*cur
		}
		if sum == t {
			now := []rune(exp)
			now = now[1:]
			ans = append(ans, string(now))
		}
		exp = oexp
		sum = osum
		return
	}
	if index == 0 {
		cur = int(nums[index] - '0')
		search(index+1, op)
	} else {
		ocur := cur
		//not action
		if cur != 0 {
			cur = cur*10 + int(nums[index]-'0')
			search(index+1, op)
			//op cur +/-/*
		}
		cur = ocur

		//add
		if op == 0 {
			osum := sum
			oexp := exp
			exp = exp + "+"
			exp = exp + strconv.Itoa(cur)
			sum = sum + cur

			cur = int(nums[index] - '0')
			search(index+1, 0)
			search(index+1, 1)
			prev = ocur
			search(index+1, 2)
			prev = 0

			exp = oexp
			cur = ocur
			sum = osum
		}
		if op == 1 {
			oexp := exp
			osum := sum
			exp = exp + "-"
			exp = exp + strconv.Itoa(cur)
			sum = sum - cur

			cur = int(nums[index] - '0')
			search(index+1, 0)
			search(index+1, 1)

			prev = ocur * -1
			search(index+1, 2)
			prev = 0

			exp = oexp
			cur = ocur
			sum = osum
		}
		if op == 2 {
			oexp := exp
			exp = exp + "*"
			exp = exp + strconv.Itoa(cur)
			osum := sum
			oprev := prev
			sum -= prev
			sum = sum + cur*prev

			prev = 0
			cur = int(nums[index] - '0')

			search(index+1, 0)
			search(index+1, 1)

			prev = oprev * ocur
			search(index+1, 2)

			exp = oexp
			prev = oprev
			sum = osum
			cur = ocur
		}
	}
}
func addOperators(num string, target int) []string {
	ans = []string{}
	exp = ""
	nums = []byte(num)
	cur = 0
	prev = 0
	sum = 0
	t = target
	search(0, 0)
	return ans
}
func main() {

}
