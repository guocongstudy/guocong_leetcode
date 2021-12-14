package main

import "fmt"

//单词拆分

func wordBreak(s string, wordDict []string) bool {
	/*
	单词就是物品，字符串s就是背包，单词能否组成字符串s，就是问物品能不能把背包装满。
	拆分时可以重复使用字典中的单词，说明就是一个完全背包！
	*/
	wordDictSet := make(map[string]bool)
	for _, w := range wordDict {
		wordDictSet[w] = true
	}
	//dp里面放的是i，下面i是从1开始，不能从0开始，决定了他就是一定是在len(s)没有-1，len(s) =9的话初始出来的切片只有0-8
	dp := make([]bool, len(s)+1)
    //定义 dp[0]=true表示空串且合法。
	//从递归公式中可以看出，dp[i] 的状态依靠 dp[j]是否为true，那么dp[0]就是递归的根基，dp[0]一定要为true，
	dp[0] = true
	//这里不能从0开始取是因为后面s[j:i] 就可能会变成s[0:0]
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			//wordDictSet[s[j:i]]在这里当map的key，所以才会有对比。例如：i=1,j=0，看l 是不是和leet相等的这个操作

			if dp[j] && wordDictSet[s[j:i]] {
				//加dp[j] && sand和and这种情况，分了sand ，分不出来and，举例为false的那种情况，用内层的j的 不是i
				//if wordDictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	//dp数组以及下标的含义:
	//dp[i] : 字符串长度为i的话，dp[i]为true，表示可以拆分为一个或多个在字典中出现的单词
	//只有当dp[i]为true时，才能说明背包装满了
	return dp[len(s)]
}

func main() {
	s := "catsandog"
	nums := []string{"cats", "dog", "sand", "and", "cat"}

	fmt.Println(wordBreak(s, nums))
}
