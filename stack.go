package main

import "math"

//最小栈
type MinStack struct {
	//定义两个栈，一个是常规栈，一个是最小站
	//普通栈（先进后出）
	StackCommon []int
	//最小栈，在这个栈里存放大小绝对顺序，越小的越在栈顶
	StackMin []int
	//记录讲个栈的大小，减少遍历
	CommonLength int
	MinLength int
}

//先初始化两个栈（上面的结构体）
func Construct() MinStack{
	return MinStack{
		StackCommon: []int{},
		//加一个int64最大值，可以减少min栈为空的判断，就相当于在StackMin中第0个元素放了一个最大值
		StackMin: []int{math.MaxInt64},
		MinLength: 1,
		//CommonLength int 默认为0
	}
}

//进栈
func (this *MinStack) PushStack(x int) {
	//添加元素的时候，应遵循min栈里放的是依次递减的元素，最上面是最小的一个元素
	 this.StackCommon = append(this.StackCommon,x)
	 this.CommonLength ++
	 if this.StackMin == nil{
	 	this.StackMin = append(this.StackMin,x)
	 	this.MinLength ++
	 }
	 if this.StackMin != nil{
	 if this.StackMin[this.MinLength-1]>x{
	 	this.StackMin = append(this.StackMin,x)
	 	this.MinLength ++
	 }
	 }
}
//出栈
func (this *MinStack) PopStack() {
	//这个就是出栈的操作，只是操作没有弹出元素的这个结果
	//x就是普通栈的栈顶元素
	x :=this.StackCommon[this.CommonLength-1]
	//这个切片元素是不包括x的
	this.StackCommon =this.StackCommon[:this.CommonLength-1]
    this.CommonLength--
    //若普通栈出栈的元素是最小元素，那么min栈也需要进行弹出操作
    if this.StackMin[this.MinLength-1] == x{
    	this.StackMin =this.StackMin[:this.MinLength-1]
    	this.MinLength--
	}
}
//栈顶元素
func (this *MinStack) TopStack()int{
	x :=this.StackCommon[this.CommonLength-1]
	return x
}
//最小元素
func (this *MinStack) MinStack()int{
	x :=this.StackMin[this.MinLength-1]
	return x
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */




//  type MinStack struct {
//     //普通栈（后进先出）
//     stackData []int
//     //辅助栈，维护一个最小元素在data栈相对顺序的栈
//     stackMin []int
//     //记录当前两个栈的大小，减少遍历
//     lengthData, lengthMin int
// }


// /** initialize your data structure here. */
// func Constructor() MinStack {
//     return MinStack{
//         stackData : []int{},
//         //加一个int64最大值，可以减少对min栈为空的判断
//         stackMin  : []int{math.MaxInt64},
//         lengthMin : 1,
//     }
// }


// func (this *MinStack) Push(x int)  {
//     this.stackData = append(this.stackData, x)
//     this.lengthData++

//     //最小元素大于x时无需加入
//     if this.stackMin[this.lengthMin-1] >= x{
//         this.stackMin = append(this.stackMin, x)
//         this.lengthMin++
//     }
// }

// func (this *MinStack) Pop()  {
//     x := this.stackData[this.lengthData-1]
//     //出栈
//     this.stackData = this.stackData[:this.lengthData-1]
//     this.lengthData--
//     //若data栈出栈的元素是最小元素，那么min栈也需要进行弹出操作
//     if this.stackMin[this.lengthMin-1] == x {
//         this.stackMin = this.stackMin[:this.lengthMin-1]
//         this.lengthMin--
//     }
// }


// func (this *MinStack) Top() int {
//     return this.stackData[this.lengthData-1]
// }


// func (this *MinStack) GetMin() int {
//     return this.stackMin[this.lengthMin-1]
// }

// 作者：xing-you-ji
// 链接：https://leetcode-cn.com/problems/min-stack-lcci/solution/lai-zi-zuo-shen-de-jie-ti-si-lu-wo-yi-yi-hyp4/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。