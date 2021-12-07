package main

import (
	"fmt"
	"sort"
)

//三数之和要为0 且三个数还不能一样
func threeSum(nums []int) [][]int{
	sort.Ints(nums)
	res :=[][]int{}

	for i:=0;i<len(nums)-2;i++{
		n1 :=nums[i]
		if n1>0{
			break
		}
		//仔细看这里，就是因为i>0,所以才会有nums[i-1]
		if i>0 && n1==nums[i-1]{
			// [-2,2,2,2] 这种情况，第三个数只会比-2,2大，所以这次循环肯定不成功
			continue
		}
		l,r :=i+1,len(nums)-1
		for l<r {
			n2,n3 :=nums[l],nums[r]
			if n1+n2+n3 == 0 {
				res =append(res,[]int{n1,n2,n3})
				//举例：[-2,-1,1=l,1=r,2,3]
				for l<r && nums[l]==n2{
					l++
				}
				for l<r && nums[r]==n3{
					r--
				}
			}else if n1+n2+n3 <0{
				l++
			}else{
				r--
			}
		}
	}
	return res
}


func main(){
	var nums =[]int{-3,-2,1,4,2,5}


	fmt.Println(threeSum(nums))
}