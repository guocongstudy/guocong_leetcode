package main

type TripInOne struct {
	stacks [][]int
	stackSize int
}

func Construct2(stackSize int) TripInOne{
	return TripInOne{
		stacks :make([][]int,3),
		stackSize: stackSize,
	}
}
func (this *TripInOne) Push(stackNum int ,value int){
	if len(this.stacks[stackNum]) <this.stackSize{
		this.stacks[stackNum] =append(this.stacks[stackNum],value)
	}
}

func (this *TripInOne) Peek(stackNum int)int{
	if !this.IsEmpty(stackNum){
		return this.stacks[stackNum][len(this.stacks[stackNum])-1]
	}
	return -1
}
func (this *TripInOne) IsEmpty(stackNum int)bool{
	return len(this.stacks[stackNum]) == 0
}


func (this *TripInOne) Pop(stackNum int) int{
	value :=this.Peek(stackNum)
	if value !=-1{
		this.stacks[stackNum] = this.stacks[stackNum][0:len(this.stacks[stackNum])-1]
	}
	return value
}