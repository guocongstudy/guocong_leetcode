package main

import (
	"fmt"
	"math"
)

type MinStack1 struct {
	stack []int
	minStack []int
}

func Constructor1() MinStack1{
	return MinStack1{
		stack :[]int{},
		minStack: []int{math.MaxInt64},
	}
}

/*

push(x) —— 将元素 x 推入栈中。
pop() —— 删除栈顶的元素。
top() —— 获取栈顶元素。
getMin() —— 检索栈中的最小元素。

*/

func (this *MinStack1) Push(x int){
	this.stack =append(this.stack,x)
	//在minStack栈放了一个top指针，一直指向minStack的最后一个元素
	top :=this.minStack[len(this.minStack)-1]
	//看，用的是min()，所以上面初始化minStack第一个元素放的是一个int64的最大值
	//如果x比top还要小，将这个值也存入minStack中
	this.minStack =append(this.minStack,min2(x,top))
}

//pop() —— 删除栈顶的元素。
func (this *MinStack1) Pop(){
	this.stack =this.stack[:len(this.stack)-1]
	this.minStack =this.minStack[:len(this.minStack)-1]
}
//top() —— 获取栈顶元素。
func (this *MinStack1) Top()int{
	return this.stack[len(this.stack)-1]
}

func (this *MinStack1) GetMin()int {
	return this.minStack[len(this.minStack)-1]
}

func min2(x,y int)int {
	if x<y {
		return x
	}else{
		return y
	}
}

func main(){
	var s =[]int{math.MaxInt64}
	fmt.Println(len(s))


}