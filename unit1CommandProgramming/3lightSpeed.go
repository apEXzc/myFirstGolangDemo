package main

import "fmt"

func main() {
	//const 无法修改
	const lightSpeed = 299792
	//变量必须先声明 后使用
	var distance = 56000000
	fmt.Println(distance/lightSpeed, "seconds")
	distance = 401000000
	fmt.Println(distance/lightSpeed, "seconds")

	//一次性声明多个变量

	var (
		test  = 1
		test2 = 2
	)
	var test3, test4 = 3, 4
	fmt.Println(test + test2)
	fmt.Println(test3 + test4)

	//增量并赋值操作符 即
	test += 1
	test = test + 1
	fmt.Println("test += 1 和test = test + 1 等效")
	//不支持前置增量 即++count
	var weight = 90
	weight -= 2
	fmt.Println(weight)
}
