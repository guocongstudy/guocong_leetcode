package main

import "fmt"

//最大正方形
func maximalSquare(matrix [][]byte) int {
	maxSide :=0
	if len(matrix) == 0 || len(matrix[0]) ==0 {
		return maxSide
	}
	rows,columns :=len(matrix),len(matrix[0])
	for i :=0 ;i<rows;i++{
		for j:=0 ;j<columns;j++{
			if matrix[i][j] == '1'{
				//maxSide=最后求面积的时候取到的最大的边
				maxSide =maxq(maxSide,1)
				//实际由于行数和列数的限制，能取到的最大正方形边数
				curMaxSide :=minq(rows-i,columns -j)
				for k:=1;k<curMaxSide;k++{
					flag :=true
					if matrix[i+k][j+k] == '0'{
						break
					}
					//这里要用到上面的k，m从0开始，m<k，k必须要从1开始
					for m:=0;m<k;m++{
						if matrix[i+k][j+m] == '0'|| matrix[i+m][j+k] == '0'{
							flag =false
							break
						}
					}
					if flag {
						//始终k（右下角）和左上角那个是相差了1的距离，所以实际的长度要加1
						maxSide =maxq(maxSide,k+1)
					}else{
						break
					}
				}
			}
		}
	}
	return maxSide *maxSide
}

func maxq(a,b int)int{
	if a>b {
		return a
	}else{
		return b
	}
}

func minq(a,b int)int{
	if a<b {
		return a
	}else{
		return b
	}
}
//func maximalSquare(matrix [][]byte) int {
//	maxSide := 0
//	if len(matrix) == 0 || len(matrix[0]) == 0 {
//		return maxSide
//	}
//	rows, columns := len(matrix), len(matrix[0])
//	for i := 0; i < rows; i++ {
//		for j := 0; j < columns; j++ {
//			if matrix[i][j] == '1' {
//				maxSide = max(maxSide, 1)
//				curMaxSide := min(rows - i, columns - j)
//				for k := 1; k < curMaxSide; k++ {
//					flag := true
//					if matrix[i+k][j+k] == '0' {
//						break
//					}
//					for m := 0; m < k; m++ {
//						if matrix[i+k][j+m] == '0' || matrix[i+m][j+k] == '0' {
//							flag = false
//							break
//						}
//					}
//					if flag {
//						maxSide = max(maxSide, k + 1)
//					} else {
//						break
//					}
//				}
//			}
//		}
//	}
//	return maxSide * maxSide
//}
//
//func max(x, y int) int {
//	if x > y {
//		return x
//	}
//	return y
//}
//
//func min(x, y int) int {
//	if x < y {
//		return x
//	}
//	return y
//}
//
//


func main(){
	var nums =[][]byte{{'1','0','1','0','0'},{'1','0','1','1','1','1'},{'1','1','1','1','1'},{'1','0','0','1','0'}}
    fmt.Println(maximalSquare(nums))
}
