package main

import "fmt"

//电话号码的字母组合

var phoneMap = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var combinations []string

func letterCombinations(digits string) []string {
	if len(digits) == 0{
		return []string{}
	}
	//最终输出的切片combinations
	combinations = []string{}
	//给他们定一个初始值，开始进行回溯
	backtrack(digits,0,"")
	return combinations
}

func backtrack(digits string,index int,x string){
	//结束循环的条件
	if index ==len(digits){
       combinations =append(combinations,x)
       //在这里将数组进行打印一下
       //fmt.Println(combinations)
	}else{
		//把那个index对应的数字例如：3转成字符串类型“3”
		//2->"2"
		digit :=string(digits[index])
		//找到"3"对应的那些英文“def”
		//"2"->"abc"
		letters :=phoneMap[digit]
		//求出英文“def”的长度 =3
		// "abc" = 3
		lettersCount :=len(letters)

		//然后开始循环，应该将“def”里面的字母都循环一遍
		for k:=0;k<lettersCount;k++{
			//开始回溯(32,1,"d+a")
			//这里的idnex在不断的+1，外面的index 是不是没动呀  **最后一个for循环 脑子疼
			backtrack(digits,index+1,x+string(letters[k]))
		}
	}
}
//func backtrack(digits string ,index int,x string){
//
//	if index == len(digits){
//		combinations =append(combinations,x)
//	}else{
//		y := string(digits[index])
//		x := phoneMap[y]
//		n :=len(x)
//		for i:=0 ; i<n ; i++{
//			backtrack(digits,index+1,x+string(x[i]))
//		}
//	}
//
//}

func main(){
	fmt.Println(letterCombinations("23"))
}