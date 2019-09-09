// go 语言数组

/* 数组是具有相同唯一类型的一组已编号且长度固定的数据项序列。
数组元素可以通过索引（位置）来读取（或者修改），索引从 0 开始，第一个元素索引为 0，第二个索引为 1，以此类推。
*/


/*
声明数组
var variable_name [size] varible_type
var balance [10] float32
*/

// 初始化数组
/* 
1. var balance = [5]float32{1.0, 2.0, 3.0, 4.0, 5.0} // 初始化数组中 {} 中的元素个数不能大于 [] 中的数字。
2. 如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小：
	var balance = [...]float{1.0, 2.0}
	balance[1] = 2.0 // 取值
*/

// 访问数组元素
/*
数组元素可以通过索引（位置）来读取。格式为数组名后加中括号，中括号中为索引的值。例如：

var salary float32 = balance[1]
*/


// 实例
package main

import "fmt"

func main(){
	var n [10]int // n是一个长度为10的数组
	var i, j int

	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}
	for j = 0; j < 10; j++ {
		fmt.Println(j, n[j])
	}

}