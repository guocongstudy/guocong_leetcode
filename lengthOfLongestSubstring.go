package main

import "fmt"
//
//func lengthOfLongestSubstrings(s string) int {
//	// 哈希集合，记录每个字符是否出现过
//	m := map[byte]int{}
//	n := len(s)
//	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
//	rk, ans := -1, 0
//	for i := 0; i < n; i++ {
//		if i != 0 {
//			// 左指针向右移动一格，移除一个字符
//
//			//fmt.Println(s[i-1])
//			delete(m, s[i-1])
//
//		}
//		for rk + 1 < n && m[s[rk+1]] == 0 { //int默认值为0
//			// 不断地移动右指针
//			m[s[rk+1]]++
//			rk++
//		}
//		// 第 i 到 rk 个字符是一个极长的无重复字符子串
//		ans = maxr(ans, rk - i + 1)
//	}
//	return ans
//}
//
//func maxr(x, y int) int {
//	if x < y {
//		return y
//	}
//	return x
//}
//



func lengthOfLongestSubstrings(s string) int {
	left, n := -1, len(s)
	res := 0
	window := make(map[byte]int, 0)

	for right := 0; right < n; right++ {
		ch := s[right]
		window[ch]++

		for window[ch] > 1{
			left++
			window[s[left]]--
		}

		res = maxs(res, right - left)
	}

	return res
}

func maxs(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main(){
	fmt.Println(lengthOfLongestSubstrings("abcabcbb"))

}


/*
//最长刷题模板
func module(A []int) int {
	n := len(A)
	left, ans := -1, 0
	for right := 0; right < N; right++ {
    // 1. 直接将A[right]加到区间中，形成(left, right]
    // 2. 将A[right]加入后，惰性原则
    for check((left, right]){		// TODO，检查区间是否满足条件
			left++		// 不满足条件，移动左指针
      // TODO 修改区间状态
		}

    // assert 此时(left, right]必然满足条件
		ans = max(ans, i - left)
	}
}





*/