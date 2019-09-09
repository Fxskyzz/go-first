// go语言作用域

/*
Go 语言中变量可以在三个地方声明：

1. 函数内定义的变量称为局部变量
2. 函数外定义的变量称为全局变量
3. 函数定义中的变量称为形式参数

*/

// 局部变量
/*
package main

import "fmt"

func main(){
	// 声明局部变量
	var a, b, c int 
	// 初始化参数
	a = 10
	b = 12
	c = a + b
	fmt.Println(a, b, c)
}
*/

// 全局变量
/*
package main

import "fmt"

// 声明全局变量
var g int 

func main	(){
	// 声明局部变量
	var a, b int 

	// 初始化参数
	a = 10
	b = 12
	g = a + b
	fmt.Println(a, b, g)
}
*/

