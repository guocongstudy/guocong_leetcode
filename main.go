package main

import (
	"fmt"
	"strings"
)

func isUnique(astr string) bool {
	if len(astr) > 1 {
		set := map[rune]bool{}
		for _, v := range astr {

			if set[v] {
				return false
			}
			set[v] = true

		}
	}
	return true
}

func canPermutePalindrome(s string) bool {
	table := make(map[int32]int, len(s))
	n := 0
	for _, v := range s {
		table[v]++
	}

	for _, v := range table {
		if v != 2 {
			n++
			if n == 2 {
				return false
			}
		}
	}
	return true
}
func maximum(a int, b int) int {
	k := a > b
	m := map[bool]int{
		true:  1,
		false: 0,
	}
	return a*m[k] + b*m[!k]
}

//判断是否互为字符重拍函数
func CheckPermutation(s1 string, s2 string) bool {

	//字符串A能否变成字符串B 有两个先天需要满足的条件
	//1.字符串的长度需要一致
	//2.字符串A出现的字符，字符串B也要有（牵扯字符比较的一般都用哈希）

	//看了解析，以上判断都对

	//注意这里，我以为只有数组的长度可以用len()去表示，记住字符串也可以
	if len(s1) != len(s2) {
		return false
	}
	m := make(map[rune]int)
	for _, v := range s1 {
		m[v]++
	}
	for _, v := range s2 {
		m[v]--
	}
	//哈希的形式，自己在goland上面打断点可以看一下运行过程，这样就能熟悉一些
	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true

}

func replaceSpaces(S string, length int) string {
	//最佳的解决方式就是字符串替换函数，可以将“空字符串”换成“%20”
	//如果有这个函数一定是在strings包里面，但是自己没用过不确定有

	//在看解析的时候看到了有这个函数
	// return strings.ReplaceAll

	return strings.ReplaceAll(string([]byte(S)[:length])," ","%20")

}




func main() {
	//math.MaxInt32
	//	arr:=[2]int{1,2}
	//	var s []int
	//	s=arr[:]
	//	copy(s,arr[:])
	//	fmt.Println(s)
	//	fmt.Println(s)
	//fmt.Println(isUnique("abca"))
	//	fmt.Println(canPermutePalindrome("tactcoa"))
	//	fmt.Println(maximum(1,23))
	//fmt.Println(CheckPermutation("aba", "baa"))
	//fmt.Println(strings.Split("adacdfw dfww", ""))

	fmt.Println(replaceSpaces("dwdwd ddddw ww",14))

	//fmt.Println(string([]byte(S)[:length]))
}
