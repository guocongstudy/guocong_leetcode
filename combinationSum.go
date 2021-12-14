package main

import "fmt"

//组合总数(回溯法)
func combinationSum(candidates []int,target int) [][]int{
	var track []int
	var res [][]int
	backtracking(0,0,target,candidates,track,&res)
	return res
}

func backtracking(startIndex,sum,target int,candidates,track []int,res *[][]int){
   //终止条件
	if sum ==target{
		tmp :=make([]int,len(track))
		//将src复制到dst，将后面的复制到前面
		copy(tmp,track)//拷贝
		*res =append(*res,tmp) //放入结果集
		return
	}
	if sum >target{return}
	//回溯
	for i:=startIndex;i<len(candidates);i++{
		//更新路径集合和sum
		track = append(track,candidates[i])
		sum +=candidates[i]
		//递归
		backtracking(i,sum,target,candidates,track,res)
		//回溯
		track =track[:len(track)-1]
		sum -=candidates[i]
	}
}

func  main(){
	nums := []int{2,3,4,1,4}
	fmt.Println(combinationSum(nums,3))
}