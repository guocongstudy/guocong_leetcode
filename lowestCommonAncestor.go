package main

type TreeNodes struct {
	Val int
	Left *TreeNodes
	Right *TreeNodes
}

/*
三种树形：
最近公共祖先的情况无外乎三种情况，解答此题最好分别画出三种情况的草图，问题便迎刃而解
	○ 树形一：root为p,q中的一个，这时公共祖先为root
	○ 树形二：p,q分别在root的左右子树上（p在左子树，q在右子树；还是p在右子树，q在左子树的情况都统一放在一起考虑）这时满足p,q的最近公共祖先的结点也只有root本身
	○ 树形三：p和q同时在root的左子树；或者p,q同时在root的右子树，这时确定最近公共祖先需要遍历子树来进行递归求解。

几种判断情况：
• 清楚了这些，开始编写程序，首先要明确递归的方式，需要采用类似后序遍历的方式，先遍历左，右子树得到结果，再根据结果判断出当前树形属于前面三种讨论的哪一种，再返回相应的答案。
• 接下来开始着手实现细节：
	○ 明确递归边界：root为空（此时为空树），则直接返回NULL，不存在公共祖先；root为p或q中的一个，则直接返回root;
	这两种情况都是不需要通过递归直接能得出答案的，故直接终止递归返回即可
	○ 用left_have，right_have分别去接收遍历左右子树得到的结点（这里递归返回值是结点指针）
	○ 如果left_have和right_have均不为空，注意此时left_have,right_have并不是理解为子树上p,q的最近公共祖先，而是结点p或q自身的指针，这时说明p,q分别存在于root根结点的左右两端，是符合树形二的情况，直接返回root。
	○ 如果left_have，right_have中有一个不为空，那么通过树形的分析，必然属于树形三;
	此时left_have,right_have代表的含义就是子树上层层向上返回的最近公共祖先，最开始的调用层拿到这个结果后直接返回不空的那一个，即是答案。

*/



func lowestCommonAncestor(root, p, q *TreeNodes) *TreeNodes {
	if root ==nil{
		return nil
	}

	if root.Val ==p.Val|| root.Val ==q.Val{
		return root
	}
	left :=lowestCommonAncestor(root.Left,p,q)
	right :=lowestCommonAncestor(root.Right,p,q)
	//左右子树都不为空，返回root
	if left !=nil && right !=nil{
		return root
	}
	//有个子树值有一个
	//（仅有右子树）
	if left == nil{
		return right
	}else{
	//（仅有左子树）
		return left
	}
}