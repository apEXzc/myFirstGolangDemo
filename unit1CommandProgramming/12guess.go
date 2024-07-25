package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var boolean = true
	var myNumber = 77
	for boolean {
		var guessNumber = rand.Intn(100) + 1
		if guessNumber == myNumber {
			fmt.Printf("我猜到了，你设置的数字是%v\n", myNumber)
			boolean = false
		} else {
			fmt.Printf("我猜的数字是%v,并没有猜对\n", guessNumber)
		}
	}
}
