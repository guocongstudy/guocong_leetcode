package main

//func coinChange1(coins []int,amount int)int{
//	dp :=make([]int,amount+1)
//	//先把dp切片初始化存一个最大的数进去
//	for j := 1; j <= amount; j++ {
//		dp[j] = math.MaxInt32
//	}
//	//dp[0]必须是0
//	dp[0]=0
//	for i:=1;i<=amount;i++{
//
//		for _,value :=range coins{
//			if i>=value{
//				dp[i] = min22(dp[i],dp[i-value]+1)
//			}
//		}
//	}
//	//如果没有任何一种硬币组合能组成总金额，返回 -1 。举例：coins = [2], amount = 3
//	if dp[amount] ==math.MaxInt32{
//		return -1
//	}
//	return dp[amount]
//}
//
//func min22(a,b int)int{
//	if a<b {
//		return a
//	}else{
//		return b
//	}
//}

/*注意：为什么是j - coins[i]？
以coins = [1, 2, 5], amount = 11为例
我们要求组成11的最少硬币数，可以考虑组合中的最后一个硬币分别是1，2，5的情况，比如

最后一个硬币是1的话，最少硬币数应该为【组成10的最少硬币数】+ 1枚（1块硬币）
最后一个硬币是2的话，最少硬币数应该为【组成9的最少硬币数】+ 1枚（2块硬币）
最后一个硬币是5的话，最少硬币数应该为【组成6的最少硬币数】+ 1枚（5块硬币）

    初始化
    dp[0] = 0

*/
//
//func coinChange(coins []int, amount int) int {
//	dp :=make([]int,amount+1)
//	//现将dp进行初始化，初始化的值用一个maxInt32去标记一下，后面min函数，可以将其替换
//	for i :=1;i<=amount;i++{
//		dp[i] = math.MaxInt32
//	}
//	//dp[0]初始化为0
//	dp[0]=0
//	for j:=1;j<=amount;j++{
//		for _,value :=range coins{
//			if j>=value{
//				//j-value的位置dp最少放几种，然后再加1枚
//				dp[j] =min(dp[j],dp[j-value]+1)
//			}
//		}
//	}
//	//如果没有任何一种硬币组合能组成总金额，返回 -1 。举例：coins = [2], amount = 3
//	//因为没有合适的 所以dp一直不会更新，就一直是maxInt32
//	if dp[amount]==math.MaxInt32{
//		return -1
//	}
//	return dp[amount]
//}
//
//
//func min(a,b int)int{
//	if a<b{
//		return a
//	}else{
//		return b
//	}
//}

//func main(){
//	var coins = []int{1,2,5}
//	var amount = 11
//	fmt.Println(coinChange1(coins,amount))
//}