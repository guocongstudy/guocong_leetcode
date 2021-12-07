package main

import "fmt"

func convertInteger(A int, B int) int {
	c:=int32(A)^int32(B)
	count:=0
	for c!=0{
		count++
		c&=c-1
	}
	return count
}
func main(){
	fmt.Println(convertInteger(9,4))
}
