package main

func dailyTemperatures(temperatures []int) []int {
	//暴力解
	var res =[]int{}
	/*仔细想就三种情况：
	1.找到了后面比这个元素大的元素，存他们下标差值j-i
	2.j指针不断地后移，移到最后就找不到比i元素还大的就存个0
	3.最后一个元素，铁定后面没有比他大的，就在res切片后面存一个0
	*/

	for i:=0;i<len(temperatures)-1;i++{
		j:=i+1
		for ;j<len(temperatures);j++{
			if temperatures[j]>temperatures[i]{
				res =append(res,j-i)
				break
			}
		}
		if j == len(temperatures){
			res=append(res,0)
		}
	}
	return append(res,0)
}