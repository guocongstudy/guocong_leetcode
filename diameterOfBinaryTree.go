package main

//二叉树的直径（首先要理解，直径是什么）
func diameterOfBinaryTree(root *TreeNode) int {
	var maxDiameter int
	var dfs func(root *TreeNode)int

	dfs =func(root *TreeNode)int{
		if root ==nil {
			return 0
		}else{
			leftDeep :=dfs(root.Left)
			rightDeep :=dfs(root.Right)
			if leftDeep +rightDeep >maxDiameter{
				maxDiameter =leftDeep+rightDeep
			}
			return max(leftDeep,rightDeep) +1
		}
	}

	dfs(root)
	return maxDiameter
}