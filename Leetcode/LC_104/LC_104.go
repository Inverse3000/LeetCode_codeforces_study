package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	d := 1
	cur := root
	ans := 0
	st := []*TreeNode{}
	st = append(st, cur)
	vis := map[*TreeNode]int{}
	for len(st) >= 1 {
		cur = st[len(st)-1]
		vis[cur] = 1
		ans = max(d, ans)
		if cur.Left != nil && vis[cur.Left] == 0 {
			d++
			st = append(st, cur.Left)
			cur = cur.Left
			continue
		}
		if cur.Right != nil && vis[cur.Left] == 0 {
			d++
			st = append(st, cur.Right)
			cur = cur.Right
			continue
		}
		d--
		st = st[:len(st)-1]
	}
	return ans
}
func main() {

}
