package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func invertTree(root *TreeNode) *TreeNode{
	//对数进行判空
	if root ==nil {
		return nil
	}
	//对跟节点的左子树进行翻转
	left :=invertTree(root.Left)
    //对根节点的右子树进行翻转
	right :=invertTree(root.Right)
	//然后将翻转后的左右子树再和根节点进行拼接
	root.Left = right
	root.Right = left
	return root
}