package main

//除自身以外数组的乘积

func productExceptSelf(nums []int) []int{
	//法一：
	//length :=len(nums)
	////L，R分别表示左右两侧的乘积列表
	//L,R,answer :=make([]int,length),make([]int,length),make([]int,length)
	//
	////L[i]为索引i左侧所有元素的乘积
	////对于索引为‘0’的元素，因为左侧没有元素，所以L[0] =1
	//L[0] =1
	//for i:=1;i<length;i++{
	//	L[i] =nums[i-1]*L[i-1]
	//}
	//
	////R[i]为索引i右侧所有元素的乘积
	//R[length-1] =1
	//for i:=length-2;i>=0;i--{
	//	R[i] =nums[i+1]*R[i+1]
	//}
	//
	//for i:=0;i<length;i++{
	//	answer[i] = L[i]*R[i]
	//}
	//return answer

	//法二：
	//n :=len(nums)
	//res :=make([]int,n)
	//
	////左右
	//left,right :=1,1
	//for i:=0;i<n;i++{
	//	res[i] =left
	//	left *=nums[i]
	//}
	//
	//for i:=n-1;i>=0;i--{
	//	res[i]=right
	//	right *=nums[i]
	//}
	//for i:=0;i<n;i++{
	//	res[i]=1
	//}
	//
	//return res

	//法三：(超时)
	var temp = 1
	var res =[]int{}
	n:=len(nums)
	for i:=0;i<n;i++{
		for j:=0;j<n;j++{
			if i==j{
				continue
			}else{
				temp *=nums[j]
			}

		}
		res =append(res,temp)
	}

	return res
}