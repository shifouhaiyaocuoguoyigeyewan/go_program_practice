package main

import "fmt"

//import "fmt"

/**
  @author: Ekko
  @since: 2023/1/2
  @desc: //TODO
**/

func main() {
	var n int            //n给出的正整数
	var tag byte         //tag符号
	var sum, nextSum int //sum符号总数量，nextSum下轮符号总数量
	_, _ = fmt.Scanf("n:%d, 符号:%c", &n, &tag)
	i := 1    //符号行数
	last := 0 //多出数
	//判断符号数量和多余数
	for {
		i += 2
		nextSum += i * 2
		if nextSum <= n {
			sum += i * 2
		} else {
			i -= 2 //多加的两行删掉
			last = n - sum
			break
		}
	}

	//分上下部分打印星号和多余数
	for j := i; j >= 1; j -= 2 { //行数
		var space = (i - j) / 2        //空格
		for k := 0; k < i-space; k++ { //判断该行打印空格位置
			if k < space { //前面空格打印
				fmt.Printf(" ")
			} else {
				fmt.Printf("%c", tag)
			}
		}
		fmt.Println()
	}
}
