package main

import "fmt"

//深度优先算法
func numIslands(grid [][]byte) (ans int){
	//m相当于求长方形有多少行，n相当于求长方形每行有几个元素
	m,n :=len(grid),len(grid[0])
	var sink func(i,j int)
	sink =  func(i,j int){
		//终止条件：超出数组范围，或者当前位置不是链接目标
		//i，j是从0开始的，所以i>=m,i>=n不能取等
		if i<0 || j<0 || i>=m || j >=n ||grid[i][j] !='1'{
			return
		}
      //沉没当前位置
      grid[i][j] = '0'
      //上方递归执行沉没
      sink(i-1,j)
      //下方
      sink(i+1,j)
      //左方
      sink(i,j-1)
      //右方
      sink(i,j+1)
	}
	//提取出一行
	for i,row :=range grid {
	//将这一行元素，遍历出k,v
		for j,ch :=range row{
			//发现岛头value=1时，开始沉没
			if ch =='1' {
				//sink里面实际存的是那个元素的坐标
				sink(i,j)
				//***操，少了一个ans++一直找不到问题***
				ans++
			}
		}
	}
	return
}

func main(){
	var nums = [][]byte{{'1','1','1','1','0'},{'1','1','0','1','0'},{'1','1','0','0','0'},{'0','0','0','0','0'}}
	fmt.Println(numIslands(nums))
}