// go语言函数
/*
函数定义
func function_name([parameter list]) [return_type] {
	function body
}

1. func：函数由 func 开始声明
2. function_name：函数名称，函数名和参数列表一起构成了函数签名。
3. parameter list：参数列表，参数就像一个占位符，当函数被调用时，你可以将值传递给参数，这个值被称为实际参数。参数列表指定的是参数类型、顺序、及参数个数。参数是可选的，也就是说函数也可以不包含参数。
4. return_types：返回类型，函数返回一列值。return_types 是该列值的数据类型。有些功能不需要返回值，这种情况下 return_types 不是必须的。
5. 函数体：函数定义的代码集合。
*/


// 实例
package main

import "fmt"


func max(num1, num2 int) int { // 定义一个求最大值的函数
	var rs int
	if (num1 > num2) {
		rs = num1
	}else{
		rs = num2
	}
	return rs
}

// 主函数，函数调用
func main(){
	var a = 100
	var b = 200
	var c int
	c = max(a, b)
	fmt.Println(c)
}


// 闭包:外层函数返回内层函数的引用
func outfunc() func innerfunc() int {
	i := 1
	return innerfunc() int{
		i += 1
		return i
	}
}