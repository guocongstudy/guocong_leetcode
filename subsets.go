package main

import "fmt"

//回溯法(想到那棵树)

/*
go 回溯套路模板
// 伪码
var res [][]int
func backTrack(路径，选择列表){
    if 满足结束条件{ // 终止条件
        res = append(res, 路径)  // 存放结果
        return
    }
    for _,选择 := range 选择列表{ // 选择：本层集合中元素（树中节点孩子的数量就是集合的大小）
		做选择  // 处理节点
        backTrack(路径，选择列表)  // 递归
        撤销选择  // 回溯，撤销处理结果
    }
}
*/

func subSets(nums []int)[][]int{
	var res [][]int
	var path []int
	var backTrace func(int)
	backTrace = func(start int){
		//递归退出条件
		if start >len(nums){
			return
		}
		temp :=make([]int,len(path))
		copy(temp,path)
		res =append(res,temp)
		for i:=start;i<len(nums);i++{
			path =append(path,nums[i])
			backTrace(i+1)
			path = path[:len(path)-1]
		}
	}
	backTrace(0)
	return res
}


func main(){
	var nums = []int{1,2,3}

	fmt.Println(subSets(nums))
}




