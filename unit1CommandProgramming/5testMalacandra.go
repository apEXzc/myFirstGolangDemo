package main

import (
	"fmt"
)

func main() {
	const distance = 56000000 //KM
	const day = 28
	var speed = distance / (day * 24)
	fmt.Printf("需要%v千米时的速度才能28天内到达\n", speed)
}
