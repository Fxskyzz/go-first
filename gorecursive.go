// go语言函数递归

// 递归: 就是在运行的过程中调用自己

/*
func recursive(){
	recursive() // 函数调用自身
}

func main(){
	recursive
}
**/
/*
1. Go 语言支持递归。但我们在使用递归时，开发者需要设置退出条件，否则递归将陷入无限循环中。
2. 递归函数对于解决数学上的问题是非常有用的，就像计算阶乘，生成斐波那契数列等。
**/

// go实现函数递归阶层]

/*

package main

import "fmt"

func Factorial (n uint64) (result uint64){
	if (n > 0){
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func main(){
	var i uint64 = 15
	fmt.Println(Factorial(i))
}
**/

// go实现斐波那契数列

package main

import "fmt"

func Fibo(n int) int{
	if n < 2{
		return n
	}
	return Fibo(n-1) + Fibo(n -2)
	
}

func main(){
	var i int
	for i = 0; i < 10; i++{
		fmt.Println(Fibo(i))
	}
}