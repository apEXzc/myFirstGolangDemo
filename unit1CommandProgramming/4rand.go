package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Intn返回的是一个从0开始到n-1的数字 因此要是想得到1-10需要+1
	var num = rand.Intn(10) + 1
	fmt.Println(num)
	num = rand.Intn(10) + 1
	fmt.Println(num)

	var distance = rand.Intn(401000000-56000000+1) + 5600000
	fmt.Println(distance)
}
