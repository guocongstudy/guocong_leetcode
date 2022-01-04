package main

import (
	"fmt"
	"math"
)

/*
1.先将数组由小到大进项排序，并且定义一个临时的盒子 res 变量，可以把这个盒子理解为等会要进行对比的内容
2.定义一个固定的坐标，也就是 n1 然后 len(nums) - 2 表示等会需要定义 n2 和 n3 两个下标的值
3.定义 l 和 r 坐标开始由两边往中间夹逼取值，并且 n2 和 n3 就是 l , r 坐标的值
4.如果 n1 + n2 + n3 小于 target 就 l++ 由于已经排好序所以就是慢慢取较大的值，如果 n1 + n2 + n3 大于 target 则 r--
5.定义第二个盒子 temp 值为 n1 + n2 + 3 用于和 res 进行比较
6.如果 num(temp - target) < num(res - target) 两数之差越小越好所以 res = temp
7.return res
*/

func threeSumClosest(nums []int, target int) int {

	res := 0
	path := 999.00

	for i := 0; i < len(nums); i++ {
		for j := i+1; j < len(nums); j++ {
			for k := j+1; k < len(nums); k++ {
				if done := nums[i] + nums[j] + nums[k]; math.Abs(float64(target - done)) < path  {
					res = done
					path = math.Abs(float64(target - done))
				}
			}
		}
	}

	return res
}


//
//func threeSumClosest(nums []int, target int) int {
//	res, n, diff, abs := 0, len(nums), math.MaxInt32, 0
//	sort.Ints(nums)
//	for k := 0; k < n-2; k++ {
//		for i, j := k+1, n-1; i < j; {
//			sum := nums[k] + nums[i] + nums[j]
//			abs = int(math.Abs(float64(sum - target)))
//			if abs < diff {
//				res, diff = sum, abs
//			}
//			if sum > target {
//				j--
//			} else if sum < target {
//				i++
//			} else {
//				return res
//			}
//		}
//	}
//	return res
//}

func main(){
	var nums = []int{-1,2,1,-4}
	target:=1
	fmt.Println(threeSumClosest(nums,target))
}