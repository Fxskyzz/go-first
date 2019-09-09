// go语言常量

// 常量是一个简单值的标识符，在程序运行时，不会被修改的量。常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。

/*
常量的定义格式
const indentifier [type] value // type说明符可以省略
1. 显示类型定义: const b string = "go"
2. 隐式类型定义: const b = "go"

多个相同类型的声明可以简写为:
const name1, name2, name3 = value1, value2, value3 // 可以看出和变量类型定义类似
*/


// 实例
/*
package main

import "fmt"

func main(){
	const length int = 2
	const WIDTH int = 3
	var area int 
	const a, b, c = 1, false, "go" // 多重赋值
	const d = 3 // 可以不被使用，不同于变量

	area = length * WIDTH
	fmt.Println(area)
	fmt.Println(a, b, c)
}
*/

// 常量的枚举
/*
const (
	male = 2
	female = 1
	secrect = 0
)
*/

/*
package main

import (
	"fmt"
	"unsafe"
)
const (
	a = "a, b, c"
	b = len(a)
	c = unsafe.Sizeof(a)
)

func main(){
	fmt.Println(a, b, c)
}
*/


/* 
iota, 特殊常量，可以认为是一个可以被编译器修改的常量。
iota在const关键字出现时被重置为0(const)

ota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
*/

package main

import "fmt"

func main(){
	const(
		a = iota
		b = iota
		c = iota
		d = iota
	)
	
	fmt.Println(a, b, c)
}



