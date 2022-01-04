package main

import "fmt"

/*
当前位置为正数时，假如以前一位为结尾的子数组能获得的最大值mx为正数，则得到当前位置的最大值
当前位置为负数时，需要得到以前一位为结尾的子数组的最小值mn，假如mn<0，则能得到一个正数，并可能成为一个新的最大值

最终需要计算以当前位置为结尾的子数组的最大值及最小值
dpmax[i] = max(dpmax[i] * maxF,max(dp[i],dp[i] * minF)) //计算当前值与最大与最小值乘积的最大值
dpmin[i] = min(dpmin[i] * minF,min(dp[i],dp[i] * maxF)) //计算当前值与最大与最小值乘积的最小值
*/


func maxProduce(nums []int)int {
	if len(nums) ==1 {
		return nums[0]
	}
	maxP ,minP,res :=nums[0],nums[0],nums[0]
	//注意nums[1:]不能写掉了，我写的时候就是忘了这个
	//前面nums[0]相当于已经当了初始值，如果不从[1:]开始的话，将nums[0]多乘了一遍
	for _,v :=range nums[1:]{
		mx, mn := maxP, minP
		//正数也要正的更多，max里面要用min里最小
		maxP =max1(maxP*v,max1(v,mn*v))
		//负数也要负的更多，min里面要用max里最大
		minP =min1(minP*v,min1(v,mx*v))
		res = max1(maxP, res)
	}
	return res
}


func max1(x,y int) int{
	if x<y {
		return y
	}else{
		return x
	}
}

func min1(x,y int)int{
	if x<y {
		return x
	}else{
		return y
	}
}


func main(){
	var nums = []int{2,3,-2,4}
	fmt.Println(maxProduce(nums))
}