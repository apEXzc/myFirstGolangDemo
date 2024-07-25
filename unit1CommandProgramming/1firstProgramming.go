// 如果单纯编译一个文件 package必须为main 否则报错
// 这是因为 Go 的编译器需要一个明确的入口点来执行程序。
// main 包和 main 函数就是这个入口点。
package main

import (
	"fmt"
)

// 声明函数 函数体需要用{}包裹
// 唯一支持大括号的放置风格如下 各占一行
func main() {
	fmt.Println("Hello, World")
	fmt.Println("Hello,世界")
}
