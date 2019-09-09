// Go 语言中数组可以存储同一类型的数据，但在结构体中我们可以为不同项定义不同的数据类型。
// 结构体是由一系列具有相同类型或不同类型的数据构成的数据集合。

/*
定义结构体
结构体定义需要使用 type 和 struct 语句。struct 语句定义一个新的数据类型，结构体有中有一个或多个成员。type 语句设定了结构体的名称。结构体的格式如下：
type struct_variable_type struct {
	member defination;
	memver defination;
	...
}

一旦定义了结构体类型，它就能用于变量的声明，语法格式如下：
variable_name := struct_variable_type {value1, value2, value3...}
或
variable_name := struct_variable_type { key1: value1, key2: value2...}
*/


// 实例
/*
package main

import "fmt"

type Books struct{
	title string
	author string
	subject string
	book_id int
}

func main(){
	// 常见一个新的结构体
	// fmt.Println(Books{"go", "hello", "world", 1})
}
*/

// 访问结构体成员
package main

import "fmt"

type Books struct {
	title string
	author string
	subject string
	book_id int
}

func main(){
	var book1 Books // 声明book1为books类型
	var book2 Books // ...
	book1.title = "go"
	book2.title = "hello"

	fmt.Println(book1.title)
	fmt.Println(book2.title)
}