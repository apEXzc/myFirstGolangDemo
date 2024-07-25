package main

import "fmt"

func main() {
	fmt.Print("我的体重是")
	fmt.Print(149.0 * 0.3788)
	fmt.Print(" lbs,我现在大约")
	fmt.Print(41 * 365 / 687)
	fmt.Println("岁")

	//格式化输出
	fmt.Printf("我的体重是%v lbs,", 149.0*0.3788)
	fmt.Printf("我现在大约%v岁\n", 41*365/687)
	//\n为换行符
	fmt.Printf("我的体重是%v lbs,我现在大约%v岁\n", 149.0*0.3788, 41*365/687)
	/*
		-15 表示左对齐，并且占用 15 个字符宽度。如果字符串的长度小于 15，则在右边填充空格。
		4 表示占用 4 个字符宽度。如果数字的长度小于 4，则在左边填充空格。
	*/
	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)
}
