package main
//解法一
//type MyQueue struct {
//	inStack []int
//	outStack []int
//}
//
//func Constructor() MyQueue{
//	//为啥有些初始化需要用到make，有些却不需要
//	return MyQueue{
//		inStack: make([]int,0),
//		outStack: make([]int,0),
//	}
//}
//
////存一个元素进去
//func (this *MyQueue) Push(x int){
//	this.inStack =append(this.inStack,x)
//
//}
//
//func (this *MyQueue) in2out(){
//	//这个函数的功能就相当于第一个栈的元素转存到第二个栈，然后第一个栈的元素是真删除
//	n :=len(this.inStack)
//	for n >0{
//		this.outStack =append(this.outStack,this.inStack[n-1])
//		this.inStack =this.inStack[:n-1]
//	}
//}
//
////出栈且删除
//func (this *MyQueue) Pop() int{
//
//	n :=len(this.outStack)
//	if n ==0 {
//		this.in2out()
//	}
//	x :=this.inStack[n-1]
//	this.outStack=this.outStack[:n-1]
//	return x
//}
//
////弹出栈顶元素
//func (this *MyQueue) Peek() int{
//
//	n :=len(this.outStack)
//	if n == 0{
//		this.in2out()
//	}
//	x :=this.outStack[len(this.outStack)-1]
//
//	return x
//}
//

//解法二
type MyQueue struct {
	nums []int
}

//初始化
func Constructor() MyQueue{
	return MyQueue{
		nums :make([]int,0,0),
	}
}

//添加一个元素到队列尾部
func (this *MyQueue) Push(x int){
	this.nums =append(this.nums,x)
}

//删除在队列前面的元素
func (this *MyQueue) Pop() int{
	res :=this.nums[0]
	this.nums =this.nums[1:]
	return res
}

// 得到队列最前面的一个元素
func (this *MyQueue) Peek()int{
	return this.nums[0]
}

//队列是否为空
func(this *MyQueue) Empty() bool{
	length :=len(this.nums)
	if length >0{
		return false
	}
	return true
}