package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans [][]int
var a []int

func search(x *TreeNode, sum, now int) {
	a = append(a, x.Val)
	fmt.Println(a)
	fmt.Println(x.Val)
	if x.Val+now == sum && x.Left == nil && x.Right == nil {
		ans = append(ans, a)
	}
	if x.Left != nil {
		search(x.Left, sum, now+x.Val)
	}
	if x.Right != nil {
		search(x.Right, sum, now+x.Val)
	}
	a = a[:len(a)-1]
}
func pathSum(root *TreeNode, targetSum int) [][]int {
	ans = ans[:0]
	search(root, targetSum, 0)
	return ans
}
func main() {
}
