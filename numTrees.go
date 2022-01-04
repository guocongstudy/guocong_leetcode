package main

import "fmt"

/*
注意的点：
i的不同元素节点组成的二叉搜索树的个数为dp[i],j相当于是头结点的元素，从1遍历到i为止，dp[0] = 1。
dp[i] += dp[以j为头结点左子树节点数量] * dp[以j为头结点右子树节点数量]
所以递推公式：dp[i] += dp[j - 1] * dp[i - j];(j-1 为j为头结点左子树节点数量，i-j 为以j为头结点右子树节点数量)
(从递归公式上来讲，dp[以j为头结点左子树节点数量] * dp[以j为头结点右子树节点数量] 中以j为头结点左子树节点数量为0，也需要dp[以j为头结点左子树节点数量] = 1， 否则乘法的结果就都变成0了。
所以初始化dp[0] = 1)
*/
func numTrees(n int)int{
	dp :=make([]int,n+1)
	dp[0]=1
	for i:=1;i<=n;i++{
		for j:=1;j<=i;j++{
			dp[i] += dp[j-1]*dp[i-j]
		}
	}
	return dp[n]
}

func main(){
	fmt.Println(numTrees(3))
}