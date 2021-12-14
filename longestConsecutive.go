package main

import (
	"fmt"
	"sort"
)

//最长连续序列
//func longestConsecutive(nums []int) int {
//	n :=len(nums)
//    sort.Ints(nums)
//	var guo int
//	var x,y int
//	for first:=0 ;first<n;first++{
//		for second:=first+1;second<n;second++{
//			if nums[second] == nums[first]+1 {
//				x :=first
//				//first++
//			}else{
//				y :=second-1
//
//			}
//			res :=y-x
//          continue
//		}
//	}
//	guo=y-x
//	return
//}

//这个方法和自己想的做法是很类似的，先排序 然后再找
func longestConsecutive(nums []int) int {
	//先判定特殊情况
	if len(nums) == 0 {
		return 0
	}
	//将数组元素进行排序，从小到大
	sort.Ints(nums)
	n:=len(nums)
	ans :=1
	tmp :=1
	for i :=1 ;i<n ;i++{
		//排除两个元素相等，还挨在一起，直接刷新i，让i+1
		if nums[i] == nums[i-1]{
			continue
		}
		//这个才是正常的情况，
		if nums[i] == nums[i-1]+1{
			//这才是我自己写的时候，思路错了，我老想着两个指针相减就是长度，
			//这个直接用+1，去记录长度就没毛病
			tmp++
		}else{
			//断了以后就初始化tmp
			tmp =1
		}

		x :=maxss(ans,tmp)
		//始终保持ans 里面放的是最长的连续数列的长度
		ans = x
	}
	return ans
}

func maxss(a,b int)int{
	if a>b {
		return a
	}else{
		return b
	}
}

func main() {
	//nums := []int{100, 4, 200, 1, 3, 2}
	nums := []int{0,3,7,2,5,8,4,6,0,1}
	fmt.Println(longestConsecutive(nums))
}
