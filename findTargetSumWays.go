package main

//目标和-回溯法
func findTargetSumWays(nums []int, target int) (count int) {
	var backtrack func(int, int)
	backtrack = func(index, sum int) {
		if index == len(nums) {
			if sum == target {
				count++
			}
			return
		}
		backtrack(index+1, sum+nums[index])
		backtrack(index+1, sum-nums[index])
	}
	backtrack(0, 0)
	return
}

//寻找重复数
func findDuplicate(nums []int) int {
	//定义一个mapHax用来讲nums的value和value出现的次数存入mapHax中
	mapHax := map[int]int{}
	var kks = 0
	//遍历nums的数组，数字出现一次，map的value就+1
	for _, v := range nums {
		mapHax[v]++
	}
	//然后再遍历哈希表，然后讲数字出现次数大于等于2的key进行返回，存入kks
	for kk, vv := range mapHax {
		if vv >= 2 {
			kks = kk
		} else {
			continue
		}

	}
	//最后返回kks就行，就是想要的重复数
	return kks

}
