// go循环语句
/*
循环类型
1. for 循环: 重复执行语句块
2. 循环嵌套: 在for循环中嵌套一个或多个for循环
*/

// 循环控制语句
/*
1. break
	经常用于中断当前 for 循环或跳出 switch 语句
2. continue
	跳过当前循环的剩余语句，然后继续进行下一轮循环。
3. goto
	将控制转移到被标记的语句。
*/

// 实例
package main

import "fmt"

func main(){
	for true {
		fmt.Println("go...")
	}
}