package main

import (
	"fmt"
	"sort"
)

//将相同的字母，不同的排序放在同一个数组中
//func groupAnagrams(strs []string) [][]string {
//	res :=[][]string{}
//	//定义一个map,key是string类型，value是string的切片
//	m :=make(map[string][]string,len(strs))
//
//	for _,v :=range strs{
//		s :=[]byte(v)
//         //对切片进行排序的一个函数
//         //silce的作用就是 “eat” = “101,97,116” ->"97,101,116"
//		sort.Slice(s,func(i,j int)bool{
//			//这个函数，甚至是func(i,j int) bool{} 都是固定的
//			//自己写的部门就是return，看你是从大到小排列还是从小到大排列都可以。
//			return s[i]<s[j]
//		})
//		//这个真的很巧妙
//		key :=string(s)
//		//只有当m[key]一样的时候，才能将v存入value []string 这个切片中，所以存进去的都是key一样的
//		//key经过排序，只要是相同字母不同组合排序后都一样
//		m[key] =append(m[key],v)
//	}
//	//输出的时候，就是想所有的元素存入res 返回res就行。
//	for _,v :=range m{
//		res =append(res,v)
//	}
//	return res
//}

func main(){
	 var strs = []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}


func groupAnagrams(strs []string) [][]string {
	res:=[][]string{}
	//定义一个map,key是string类型，value是string的切片
	mapHax :=make(map[string][]string,len(strs))
	for _,v :=range strs{
		//byte数据类型，并且是基本类型。byte是做为最小的数字来处理的，因此它的值域被定义为-128~127
		s :=[]byte(v)
		//对切片进行排序的一个函数
		//silce的作用就是 “eat” = “101,97,116” ->"97,101,116"
		sort.Slice(s,func(i,j int)bool{return s[i]<s[j]})
		//然后存的话肯定得转成string去存
		key:=string(s)
		//只有当m[key]一样的时候，才能将v存入value []string 这个切片中，所以存进去的都是key一样的
		//key经过排序，只要是相同字母不同组合排序后都一样
		mapHax[key] =append(mapHax[key],v)
	}
	//输出的时候，就是想所有的元素存入res 返回res就行。
	for _,v := range mapHax{
		res =append(res,v)
	}
	return res
}