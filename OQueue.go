package main

import "container/list"

type CQueue struct {
 stack1 *list.List
 stack2 *list.List
}
//初始化链表
func Constructor2() CQueue{
	return CQueue{
		stack1: list.New(),
		stack2: list.New(),
	}
}

//增加元素
func (this *CQueue) AppendCQueue(value int){
	this.stack1.PushBack(value)
}

//删除元素
func (this *CQueue) DeleteCQueue() int{
	//如果第二个栈为空的情况
	if this.stack2.Len() ==0{
		for this.stack1.Len() >0{
			this.stack2.PushBack(this.stack1.Remove(this.stack1.Back()))
		}
	}
	if this.stack2.Len() !=0{
		e :=this.stack2.Back()
		this.stack2.Remove(e)
		return e.Value.(int)
	}
	return -1
}


func main(){
	q:=Constructor2()
    q.AppendCQueue(5)

}
