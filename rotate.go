package main

import "fmt"

/*
1.先分析旋转前后的矩阵找出其中蕴含的规律
举例：第一行的第 x 个元素在旋转后恰好是倒数第一列的第 x 个元素
规律：对于矩阵中第 i 行的第 j 个元素，在旋转后，它出现在倒数第 i 列的第 j 个位置。
我们将其翻译成代码：前后矩阵由于矩阵中的行列从 0 开始计数，因此对于矩阵中的元素 matrix[i][j]，在旋转后，它的新位置为 matrixnew[j][n−i−1]。

*/
func rotate(matrix [][]int) {
	//求出每行的长度=二位数组里面一个数组的长度
	n :=len(matrix)
	//先创建一个过渡的二位数字，最后将这个数组粘贴到原来二维数组
	tmp :=make([][]int,n)
	//对二维数组里面的每个子数组进行初始化，数组长度为n
	for i:=range tmp{
		tmp[i] =make([]int,n)
	}
	//求出行数
	for i,row :=range matrix{
		//求出列数
		for j,v :=range row{
			//将matrix[i][j]=v赋值到matrixNew[j][n-i-1]
			tmp[j][n-i-1]=v
		}
	}
	copy(matrix,tmp)
}

func main(){
	var nums = [][]int{{1,2,3},{4,5,6},{7,8,9}}
	rotate(nums)
	fmt.Println(nums)
}