package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var vis [30005]int

func search(l int, r int, in []int, po []int) *TreeNode {
	if r < l {
		return nil
	}
	if l == r {
		for i := len(po) - 1; i >= 0; i-- {
			if vis[i] == 0 {
				vis[i] = 1
				break
			}
		}
		return &TreeNode{in[l], nil, nil}
	}
	tar := -1
	mark := -1
	for i := len(po) - 1; i >= 0; i-- {
		if vis[i] != 0 {
			tar = po[i]
			mark = i
			break
		}
	}
	vis[mark] = 1
	index := -1
	for i := l; i <= r; i++ {
		if tar == in[i] {
			index = i
			break
		}
	}
	if index == -1 {
		return nil
	}
	m := &TreeNode{in[index], nil, nil}
	m.Right = search(index+1, r, in, po)
	m.Left = search(l, index-1, in, po)
	return m
}
func buildTree(inorder []int, postorder []int) *TreeNode {
	vis = [30005]int{}
	return search(0, len(inorder)-1, inorder, postorder)
}
func main() {

}
