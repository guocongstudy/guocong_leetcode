package main

import "fmt"

//单词搜索

func exist(board [][]byte ,word string) bool{
    //m求的是有多少行，n求得是一行有多少列
	m , n :=len(board),len(board[0])
	used :=make([][]bool,m)
	for i :=0;i<m;i++{
		used[i] =make([]bool,n)
	}
    //r=x轴 ，c=y轴 i=要对比的字符串内字母的下标
 	var canFind func(r,c,i int) bool
	canFind = func(r,c,i int)bool {
		if i ==len(word){
			return true
		}
		//注意这里的 >=,如果没有=号，下面used[r][c] ||board[r][c] 会越界
		//筛选到最后一个都不行，那就是false
		if r<0 ||r>=m ||c<0 ||c >=n{
		//这样子应该不可以，上面的判断没有冗余，因为在递归的时候，会有r-1，c-1的情况，就不再是<=0了，有可能为负值
		//if r>=m ||c >=n{
			return false
		}
		//used里面存的是bool类型，board里面存的才是真实的字符 used[r][c] =true时 或 board[r][c] !=word[i] 返回false
		//used[r][c] 这里的作用：表示这个元素已经扫描过了 对应题目里的要求：“同一个单元格内的字母不允许被重复使用。”
		if used[r][c] ||board[r][c] !=word[i]{
			return false
		}
		used[r][c] =true
		//1.y轴不动，x轴+1（右移一格），字符串内字母位置加1
		//2.y轴不动，x轴+1（左移一格），字符串内字母位置加1
		//3.x轴不动，y轴+1（上移一格），字符串内字母位置加1
		//4.x轴不动，y轴+1（下移一格），字符串内字母位置加1
		//一个格的上下左右进行判断。
		canFindRest :=canFind(r+1,c,i+1) || canFind(r-1,c,i+1)||canFind(r,c+1,i+1) ||canFind(r,c-1,i+1)

		if canFindRest{
			return true
		}else{
			used[r][c] =false
			return false
		}
	}

	for i:=0 ;i<m ;i++{
		for j :=0 ;j<n ;j++{
			//初始化canFind(i,j,0) = canFind(0,0,0)开始
			if board[i][j] ==word[0] && canFind(i,j,0){
				return true
			}
		}
	}
   return false
}


func main(){
	var board = [][]byte{{'A','B','C','E'},{'S','F','C','S'},{'A','D','E','F'}}
	fmt.Println(exist(board,"ABCCED"))
}