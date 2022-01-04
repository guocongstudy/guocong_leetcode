package main
//二叉树转为链表
//type TreeNodess struct {
//	Val int
//	Left *TreeNodess
//	Right *TreeNodess
//}
//
///*
//思路：题目要求原地展开，故不能使用数组类变量存储全部节点值再重构树。将左右子树分别递归展开，将原左子树变为节点的右子树，再将原右子树变为当前右子树最右节点的右子树。
//*/
//
//func flatten(root *TreeNodess){
//	if root ==nil {
//		return
//	}
//	flatten(root.Left)
//	flatten(root.Right)
//	r :=root.Right
//	root.Right,root.Left =root.Left,nil
//	for root.Right != nil{
//		root =root.Right
//	}
//	root.Right =r
//}
