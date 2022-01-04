package main

//版本一：先遍历物品，再遍历背包
//func numSquares(n int) int{
//	//定义
//	dp :=make([]int,n+1)
//	//初始化
//	dp[0] = 0
//	for i:=1;i<=n;i++{
//		//存的是固定的一个较大的数
//		dp[i] =math.MaxInt32
//	}
//	//遍历物品
//	for i:=1;i<=n;i++{
//		//遍历背包
//		for j:=i*i;j<=n;j++{
//			dp[j] =min(dp[j],dp[j-i*i]+1)
//		}
//	}
//	return dp[n]
//}
//
//func min(a,b int)int{
//	if a<b {
//		return a
//	}else{
//		return b
//	}
//}
//
//func main(){
//	fmt.Println(numSquares(12))
//}