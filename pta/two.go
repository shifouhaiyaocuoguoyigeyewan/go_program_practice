package main

/**
  @author: Ekko
  @since: 2022/12/30
  @desc: //L1-002 打印沙漏  https://pintia.cn/problem-sets/994805046380707840/exam/problems/994805145370476544
         解题思路：1.判断符号数量和多余数
                 2.分上下部分打印符号和多余数
**/

import (
	"fmt"
)

func main() {
	var n int            //n给出的正整数
	var tag byte         //tag符号，byte类型，相当于uint8(0-255)
	var sum, nextSum int //sum需要符号数量，nextSum下轮符号总数量
	_, _ = fmt.Scanf("%d %c", &n, &tag)
	//n = 19
	//tag = 255

	sum = 1
	nextSum = 1
	i := 1 //i符号行数
	last := 0 //多出数
	for { // 先是计算出所需要多少的符号
		i += 2
		nextSum += i * 2
		if nextSum <= n {
			sum += i * 2
		} else {
			i -= 2
			last = n - sum //多出数
			break
		}
	}

	for j := i; j >= 1; j = j - 2 { // 先打印上半部分
		var space = (i - j) / 2 //space代表打印的空白
		for k := 0; k < i-space; k++ {
			if k < space {
				fmt.Printf(" ") // 先打印空格
			} else {
				fmt.Printf("%c", tag) //再打印*
			}
		}
		fmt.Println() // 最后记得要换行
	}
	for j := 3; j <= i; j = j + 2 { //再打印下半部分
		var space = (i - j) / 2
		for k := 0; k < i-space; k++ {
			if k < space {
				fmt.Printf(" ") //再打印*
			} else {
				fmt.Printf("%c", tag)
			}
		}
		fmt.Println()
	}
	fmt.Println(last)
}
