// go 并发
/*
Go 语言支持并发，我们只需要通过 go 关键字来开启 goroutine 即可。
goroutine 是轻量级线程，goroutine 的调度是由 Golang 运行时进行管理的。
Go 允许使用 go 语句开启一个新的运行期线程， 即 goroutine，以一个不同的、新创建的 goroutine 来执行一个函数。 同一个程序中的所有 goroutine 共享同一个地址空间。
**/

/*
语法格式
go 函数名(参数列表)
**/

// 实例
/*
package main

import "fmt"
import "time"


func say(s string)  {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main()  {	
	go say("world")
	say("go")
}
**/

// go通道
/*
通道（channel）是用来传递数据的一个数据结构。
通道可用于两个 goroutine 之间通过传递一个指定类型的值来同步运行和通讯。操作符 <- 用于指定通道的方向，发送或接收。如果未指定方向，则为双向通道。
ch <- v // 把v发送到通道ch
v := <-ch // 从ch接收数据并把值赋值给v
**/

/*
声明一个通道很简单，我们使用chan关键字即可，通道在使用前必须先创建
ch := make(chan, int)
注意：默认情况下，通道是不带缓冲区的。发送端发送数据，同时必须又接收端相应的接收数据。
以下实例通过两个 goroutine 来计算数字之和，在 goroutine 完成计算后，它会计算两个结果的和：
**/

// 实例
package main

import "fmt"

func sum(s []int, c chan int){
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把sum发送到通道c
}

func main()  {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从通道c中接收

	fmt.Println(x, y, x+y)
}




