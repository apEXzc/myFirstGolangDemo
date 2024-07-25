package main

import (
	"fmt"
)

func main() {
	//作用域和大部分语言一样的 略
	//注意 Go 没有while循环 for直接代替
	var count = 10
	fmt.Println(count)
	cmd := 10
	fmt.Println(cmd)

	for count3 := 10; count3 > 0; count3-- {
		fmt.Println(count3)
	}

}
