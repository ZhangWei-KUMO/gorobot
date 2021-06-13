// main.go
// 首先给这个包取个名称gorobot
package gorobot

import "fmt"
//  name string中的string为输入值类型，后一个string为输出值类型
func Hello(name string) string {
    // %v 为fmt包打印输出字符串占位符
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    // :=符号位go特有的初始化声明变量赋值方式，go语言对于变量不需进行类型声明。
    return message
}
