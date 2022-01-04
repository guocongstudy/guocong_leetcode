package main

import (
	"fmt"
	"sort"
)


//https://leetcode-cn.com/problems/kth-largest-element-in-an-array/
func main(){
//var nums = []int{3,2,1,5,6,4}
	var nums =[]int{3,2,3,1,2,4,5,5,6}
	sort.Ints(nums)
fmt.Println(nums)
}
