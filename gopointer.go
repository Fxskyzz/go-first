// go语言指针

// Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。

// 实例
/*
package main

import "fmt"

func main(){
	var a int = 10
	fmt.Println(&a)
}
*/



/*
一个指针变量指向了一个值的内存地址。
类似于变量和常量，在使用指针前你需要声明指针。指针声明格式如下：

var var_name *var-type
var-type 为指针类型，var_name 为指针变量名，* 号用于指定变量是作为一个指针


var ip *int // 指向整形
var fp *float32 // 指向浮点型
*/

/*
指针使用流程
1. 定义指针变量
2. 为指针变量赋值
3. 访问指针变量中指向地址的值
*/

// 实例
/*
package main

import "fmt"

func main(){
	var a int = 10
	var ip *int
	ip = &a
	fmt.Println(a, ip, &a)
}
*/


// go空指针
// 当一个指针被定义后没有分配到任何变量时，它的值为 nil。

// 实例
package main
import "fmt"
func main(){
	var ptr *int
	fmt.Println(ptr)
	// 输出为nil
}