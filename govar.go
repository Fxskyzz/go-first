// go 语言变量
/* 
声明变量形式
var indentifier type
一次声明多个变量
var indentifier1, indentifier2 type
*/

// 实例
/*
package main // 包声明

import "fmt" // 包导入
// 主函数
func main(){
	var a string = "Runoob"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)
}
*/

// 实例输出 1， 2


// 变量声明
	/*
	第一种: 指定变量类型，如果没有初始化，则变量默认为零值
	var v_name v_type
	v_name = value 
	*/

	// 实例
	/*package main

	import "fmt"

	func main(){
		// 声明一个变量并初始化
		var a = "Runoob"
		fmt.Println(a)

		// 没有初始化就为零值
		var b int 
		fmt.Println(b)

		// bool零值为false
		var c bool
		fmt.Println(c)
	}
	*/

	// 实例输出 Runoob, 0, false

	/*
	1. 数值类型(包括complex64/128为0)
	2. 布尔类型为false
	3. 空字符串为""(空字符串)
	一下几种为nil

		1. var a *int
		2. var a []int
		3. var a map[string] int
		4. var a chan int
		5. var a func(string) int
		6. var a error // error 接口
	*/


	// 实例
	/*
	package main

	import "fmt"

	func main(){
		var i int
		var f float64
		var b bool
		var s string
		fmt.Println(i, f, b, s)
	}
	*/

	/*
	第二种: 根据值自行判定变量类型
	var v_name = value
	*/

	// 实例
	/*
	package main

	import "fmt"

	func main(){
		var d = true
		fmt.Println(d)
	}
	*/

	/*
	第三种: 省略var := 左侧如果没有声明新的变量，就产生编译错误，格式：
	v_name := value

	var intVal int 
	intVal := 1 // 这个时候会产生编译错误
	intVal, intVal1 := 1, 2 // 此时不会产生编译错误 := 是一个声明语句
	*/

	// 实例
	/*
	package main
	import "fmt"
	func main(){
		f := "go" // var f string = "runoob"
		fmt.Println(f)
	}
	*/

	// 输出结果是 go

// 多变量声明
/*
1. var vname1, vname2, vname3 type 
vname1, vname2, vname3 = v1, v2, v3
*/

// 2. var vname1, vname2, vname3 = v1, v2, v3 // 和python很像，不需要显示声明变量类型， 自动推断
// 3. vname1, vname2, vname3 := v1, v2, v3 // 出现在 := 左侧的变量不应该是已经被声明过得，否则会出现编译错误

// 因为分解关键字的写法一般用于声明全局变量
/*
var (
	vname1 v_type1
	vname2 v_type2
)
*/

// 实例
/*
package main
import "fmt"
var x, y int
var (
	// 此种因式分解的写法一般用于声明全局变量
	a int
	b bool
)

var c, d int = 1, 2
var e, f = 123, "go"


// g, h := 123, "go" // 这种不带声明格式的只能在函数体中出现
func main(){
	g, h := 123, "hello"
	fmt.Println(x, y, a, b, c, d, e, f, g, h)
}
*/

// 内存地址称之为指针
// 	使用操作符 := 可以高效地创建一个新的变量，称之为初始化声明。
/*package main

import "fmt"
var &i int = 12

func main(){
	fmt.Println(&i)
}
*/



// 注意点
/*
1. 函数体内声明的变量必须在后续被使用， 否则会出现错误
2. 全局变量是允许声明但不使用。同一类型的多个变量可以声明在同一行
var a, b, c int = 1, 2, 3
*/

// 多变量可以在同一行赋值
/*
var a, b int
var c string
a, b, c = 1, 2, "go"

另外一种写法
a, b, c := 1, 2, "go"
*/

// 空白标识符 _ 也被用于抛弃值，如值 5 在：_, b = 5, 7 中被抛弃。

// 并行赋值也被用于当一个函数返回多个返回值时，比如这里的 val 和错误 err 是通过调用 Func1 函数同时得到：val, err = Func1(var1)。


// 如果你想要交换两个变量的值，则可以简单地使用 a, b = b, a，两个变量的类型必须是相同。
/*
package main

import "fmt"

func main(){
	a, b := 1, 2
	b, a = a, b
	fmt.Println(a, b)
}
*/

