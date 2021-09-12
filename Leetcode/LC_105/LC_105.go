package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var vis [30005]int

func search(l int, r int, pre []int, in []int) *TreeNode {
	if r < l {
		return nil
	}
	if l == r {
		for i := 0; i < len(pre); i++ {
			if vis[i] == 0 {
				vis[i] = 1
				break
			}
		}
		return &TreeNode{in[l], nil, nil}
	}
	tar := 0
	mark := -1
	for i := 0; i < len(pre); i++ {
		if vis[i] == 0 {
			tar = pre[i]
			mark = i
			break
		}
	}
	if mark == -1 {
		return nil
	}
	index := -1
	for i := l; i <= r; i++ {
		if in[i] == tar {
			index = i
			break
		}
	}
	if index == -1 {
		return nil
	}
	vis[mark] = 1
	m := &TreeNode{tar, nil, nil}

	m.Left = search(0, index-1, pre, in)

	m.Right = search(index+1, len(pre)-1, pre, in)
	return m
}
func buildTree(preorder []int, inorder []int) *TreeNode {
	vis = [30005]int{}
	return search(0, len(preorder)-1, preorder, inorder)
}
func main() {

}
