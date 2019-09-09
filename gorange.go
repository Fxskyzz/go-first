// go语言范围(range)

// Go 语言中 range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对的 key 值。

// 实例
package main

import "fmt"

func main(){
	nums := []int{2, 3, 4} // 切片
	sum := 0

	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)
}