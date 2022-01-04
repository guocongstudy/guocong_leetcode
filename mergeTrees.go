package main
//type TreeNode struct {
//	Val int
//	Left *TreeNode
//	Right *TreeNode
//}
//func mergeTrees(root1 *TreeNode,root2 *TreeNode) *TreeNode{
//	//先进行判断，两个树都是空
//	if root1 == nil && root2 == nil   {
//		return nil
//	}else{
//		//其中一个树为空
//		if root1 ==nil {
//			return root2
//		}
//		if root2 ==nil {
//			return root1
//		}
//        //我第一次做得时候，想的是先定义一个第三颗树，然后将数据相加以后再赋值到第三颗树即可
//        //但是直接在第一棵树上直接操作也很方便，不同自己再定义新的树
//		root1.Val =root1.Val+root2.Val
//		root1.Left  =mergeTrees(root1.Left,root2.Left)
//		root1.Right =mergeTrees(root1.Right,root2.Right)
//
//	}
//	return root1
//}
