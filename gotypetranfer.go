// go语言类型转换

package main
import "fmt"

func main(){
	var sum int = 12
	var count int = 5
	var mean float64
	mean = float64(sum)/float64(count)
	fmt.Println(mean)
}