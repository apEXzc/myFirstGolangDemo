package main

import "fmt"

func checkLeap(int2 int) bool {
	if int2%4 == 0 && int2%100 != 0 {
		return true
	} else if int2%400 == 0 {
		return true
	}
	return false
}

// 和JS一样 短路逻辑 如果||前面是对的 则不会运行后面的代码
func main() {
	fmt.Printf("2100年是闰年吗？ %v", checkLeap(2100))
}
