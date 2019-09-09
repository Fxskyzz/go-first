// Go 语言Map(集合)
// 类似Python中的字典
/*
Map 是一种无序的键值对的集合。Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。
Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，Map 是无序的，我们无法决定它的返回顺序，这是因为 Map 是使用 hash 表来实现的。
*/

/*
定义map
1. 使用map
var map_variable map[key_data_type]value_data_type

2. 使用make
map_variable := make(map[key_data_type]vale_data_type)

** 如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对
*/


package main

import "fmt"
func main(){
	var countryCapitalMap map[string]string 
	countryCapitalMap = make(map[string]string)
	// map插入key - value对,各个国家对应的首都
    countryCapitalMap [ "France" ] = "巴黎"
    countryCapitalMap [ "Italy" ] = "罗马"
    countryCapitalMap [ "Japan" ] = "东京"
    countryCapitalMap [ "India" ] = "新德里"

    for country := range countryCapitalMap{
    	fmt.Println(countryCapitalMap[country])
    }
}

// delete() 函数用于删除集合的元素, 参数为 map 和其对应的 key。
// delecte(map名, key)
