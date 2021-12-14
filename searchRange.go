package main

import (
	"fmt"
	"sort"
)

//这个函数是自己写的，过了测试用例
//func searchRange (nums []int,target int) []int{
//	//最终个返回数组记为s
//	var s = []int{}
//	n :=len(nums)
//	if n == 0 {
//		ss :=[]int{-1,-1}
//		return ss
//	}
//	var m = map[int]int{}
//
//	for k:=0;k<n;k++{
//		m[k]=nums[k]
//	}
//
//	 lm :=len(m)
//	 for t:=0;t<lm;t++{
//	 	if m[t] == target{
//	 		 s =append(s,t)
//		}
//
//	 }
//
//	if len(s) == 0 {
//		ss :=[]int{-1,-1}
//		return ss
//	}
//	if len(s) == 1 {
//		x :=s[0]
//		s :=append(s,x)
//		return s
//	}
//	if len(s) >2 {
//		lms :=len(s)
//		xs :=s[0]
//		vs :=s[lms-1]
//		var point3 =  []int{}
//		point3 = append(point3,xs)
//		point3 = append(point3,vs)
//		return  point3
//	}
//	return s
//}
func searchRange(nums []int, target int) []int {
	//func SearchInts(a []int, x int) int
	//SearchInts在递增顺序的a中搜索x，返回x的索引。如果查找不到，返回值是x应该插入a的位置（以保证a的递增顺序），返回值可以是len(a)
	leftmost := sort.SearchInts(nums, target)
	//这个就是排除了特殊情况，leftmost ==len(nums) 就是这个元素不存在，需要插入到末尾
	//nums[leftmost] !=target 就是验证没有这个元素，返回的位置是要插入的位置
	if leftmost == len(nums) || nums[leftmost] != target {
		return []int{-1, -1}
	}
	//[1,2,3,4,5,7]
	//[1,2,3,4,4,7,8] 4
	//这个target+1就写的很巧妙如果刚好target+1这个元素存在，那么他一定是比target大的数。且在大的里面最接近target的数，一定和target是邻居，位置减一，就会出现[3,3]起点和终点在同一位置
	//如果target+1这个元素不存在，那么返回的就是比target+1该待的地方。外面在一个减一就是第二个target的位置
	rightmost := sort.SearchInts(nums, target+1) - 1
	return []int{leftmost, rightmost}

}

func main() {
	var nums = []int{1}
	fmt.Println(searchRange(nums, 1))
}
