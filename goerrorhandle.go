// go错误处理

/*
error是一个接口类型，定义方式为:
type error interface{
	Error() string
}
**/
/*
func Sqrt(f float64) (float64 error) {
	if f < 0 { // 实现
		return 0, errors.New("math: square root of negative number")
	}
}
**/

// 实例
package main

import "fmt"

// 定义一个DivideError结构
type DivideError struct{
	dividee int
	divider int
}

// 实现 error 接口
func (de *DivideError) Error() string {
	strFormat := `
	Cannot proceed, the divider is zero.
	dividee: %d
	divider：0
	`
	return fmt.Sprintf(strFormat, de.dividee)

}

func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	}
	return varDividee / varDivider, ""
}

func main(){
	// 正常情况
	if result, msg := Divide(100, 10); msg == "" {
		fmt.Println(result)
	}
	// 当被除数为零的时候会返回错误信息
	if _, msg := Divide(100, 0); msg != "" {
		fmt.Println(msg)
	}
}