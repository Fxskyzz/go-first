// Go 语言切片是对数组的抽象。
// Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

/*
定义切片
你可以声明一个未指定大小的数组来定义切片
var identifier []type
	切片不需要说明长度
	或使用make来创建切片

	var slice1 []type = make([]type, len)

	可以简写为
	slice1 := make([]type, len)

	也可以指定容量，其中capacity为可选参数
	make([]type,length, capacity)
*/


/*
切片初始化
s := []int{1, 2, 3}
s := arr[:]
*/

/*
切片是可索引的，并且可以由 len() 方法获取长度。
切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少。
*/

/*
package main

import "fmt"

func main(){
	var numbers = make([]int,3,5)
	printSlice(numbers)
}

func printSlice(x []int) {
	fmt.Println(len(x), cap(x), x)
}
*/

// 空切片：一个切片在未初始化之前默认为 nil，长度为 0
/*
package main

import "fmt"

func main(){
	var numbers []int
	fmt.Println(numbers)

	if (numbers == nil){
		fmt.Println("切片是空的")
	}
}
*/




