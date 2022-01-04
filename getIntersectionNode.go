package main
//
type ListNodes struct{
	Next *ListNodes
	Val int
}

func getIntersectionNode(headA,headB *ListNodes) *ListNodes{
curA,curB :=headA,headB
for curA !=curB{
	if curA == nil {
		curA =headB
	}else{
		curA=curA.Next
	}
	if curB ==nil {
		curB = headA
	}else{
		curB=curB.Next
	}
}
return curA
}
