package main

import "fmt"

//全排列-典型的回溯
//这里用一个巧劲，思路清晰

/*
//1.不同类型的切片无法复制
//2.如果s1的长度大于s2的长度，将s2中对应位置上的值替换s1中对应位置的值
//3.如果s1的长度小于s2的长度，多余的将不做替换

func main() {
    s1 := []int{1, 2, 3}
    s2 := []int{4, 5}
    s3 := []int{6, 7, 8, 9}
    copy(s1, s2)
    fmt.Println(s1) //[4 5 3]
    copy(s2, s3)
    fmt.Println(s2) //[6 7]
}
*/

//func permute(nums []int)[][]int{
//	if len(nums) ==1 {
//		return [][]int{nums}
//	}
//	res :=[][]int{}
//	for i,value:=range nums{
//		tmp :=make([]int,len(nums)-1)
//		copy(tmp[0:],nums[0:i])
//		copy(tmp[i:],nums[i+1:])
//
//		sub :=permute(tmp)
//		for _,s:=range sub{
//			res =append(res,append(s,value))
//		}
//	}
//	return res
//}


func permute(nums []int)[][]int{
	//最终要返回的二维数组
	var res [][]int
	//已经用过的节点存储用的切片
	var path []int
	//将用过节点进行标记的哈希表
	visited :=make(map[int]bool)
	size :=len(nums)
	var backTrack func()
	backTrack = func() {
		//递归终止条件
		//也就是nums里的元素都用到了
		if len(path) ==size{
			//temp暂时存放path，path的长度肯定是nums的长度
			temp :=make([]int,size)
			copy(temp,path)
			res = append(res,temp)
			return
		}
		//从0开始所以不去等
		for i:=0;i<size;i++{
			//一个排列结果（path）里的一个元素只能使用一次
			//相当于查map里有没有这个元素，有就continue跳出
			if visited[nums[i]] {
				continue
			}
			//第一次出现就给他打个标记true
			visited[nums[i]] =true
			//将这个元素放入path路径中
			path = append(path,nums[i])
			//递归
			backTrack()
			//回溯
			visited[nums[i]] =false
			//就是吐出最后一个元素
			path = path[:len(path)-1]
		}

	}
	backTrack()
	return res
}



func main(){
	var nums =[]int{5,1,2,3}
	//var ps []int
	//ps =nums[:len(nums)-1]
	//fmt.Println(ps)
	fmt.Println(permute(nums))
}