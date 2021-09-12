package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
func search(x *TreeNode) (yes, no int) {
	lyes, lno, ryes, rno := 0, 0, 0, 0
	if x.Left != nil {
		lyes, lno = search(x.Left)
	}
	if x.Right != nil {
		ryes, rno = search(x.Right)
	}
	yes = lno + rno + x.Val
	no = max(lno, lyes) + max(rno, ryes)
	return yes, no
}
func rob(root *TreeNode) int {
	yes, no := search(root)
	return max(yes, no)
}
func main() {

}
