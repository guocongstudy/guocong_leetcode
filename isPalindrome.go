package main

import "fmt"

type ListNodess struct {
	Val int
	Next *ListNodess
}
//将值复制到数组中然后用双指针发
func isPalindrome(head *ListNodess)bool{
	vals :=[]int{}
	for; head !=nil ;head =head.Next{
		vals =append(vals,head.Val)
	}
	n :=len(vals)
	//注意：n/2是向下去整，例如：5/2 = 2
	//只用对比一半就可以，这里的vals[:n/2]就是将原数组折了一半，从vals[0]正着比
	for i,v :=range vals[:n/2]{
		//vals[n-1-j]就是后一半，vals[n-1]倒着和v进行比较
		if v !=vals[n-1-i]{
			return false
		}
	}
	return true
}

func main(){
	x :=5/2
	fmt.Println(x)
}