package main

import "fmt"

//法一：
//func maxArea(height []int) int{
//	n :=len(height)
//	l,r :=0,n-1
//	ans :=0
//	for l<r{
//		//长方形的长度
//		w :=r-l
//		//高初始化为0
//		h :=0
//		//蓄水池的高度永远是取决于最低的模板高度
//		if height[l]>height[r]{
//			h =height[r]
//
//			r--
//		}
//		if height[l]<height[r]{
//			h =height[l]
//			l++
//		}
//		tmp :=w*h
//		if tmp>ans {
//			ans =tmp
//		}
//	}
//	return ans
//
//}
//
//
//func main(){
//	nums := []int{1,3,4,5,3,2}
//	fmt.Println(maxArea(nums))
//}


func maxArea(height []int) int {
	i, j := 0, len(height) - 1
	m := 0
	for i < j {
		// 计算当前最大面积
		cur := (j - i) * min(height[i], height[j])
		if cur > m {
			m = cur
		}

		// 移动较小的一侧指针
		if (height[i] < height[j]) {
			i++
		} else {
			j--
		}
	}
	return m
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main()  {
		nums := []int{1,3,4,5,3,2}
		fmt.Println(maxArea(nums))
}