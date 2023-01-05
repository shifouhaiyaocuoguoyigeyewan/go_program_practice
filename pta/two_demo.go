package main

import "fmt"

/**
  @author: Ekko
  @since: 2023/1/2
  @desc: //TODO
**/

func main() {
	var n int            //n给出的正整数
	var tag byte         //tag符号
	var sum, nextSum int //sum符号总数量，nextSum下轮符号总数量
	//_, _ = fmt.Scanf("%d %c", &n, &tag)

	n = 19
	tag = 255

	i := 1    //符号行数
	last := 0 //多出数
	nextSum = 1
	sum = 1
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
	for j := i; j >= 1; j -= 2 { //分行打印
		var space = (i - j) / 2        //空格
		for k := 0; k < i-space; k++ { //判断该行打印空格位置
			if k < space { //前面空格打印
				fmt.Printf(" ")
			} else {
				fmt.Printf("%c", tag)
			}
		} //符号
		fmt.Println() //换行
	}

	//下部分
	for j:=3; j<=i ; j+=2 {
		var space =(i -j) / 2
		for k:=0; k < i-space; k++ {
			if k < space {
				fmt.Printf(" ")
			}else {
				fmt.Printf("%c",tag)
			}
		}
		fmt.Println()
	}
	fmt.Println(last)
}
