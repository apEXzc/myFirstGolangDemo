package main

import "fmt"

func main() {
	var command = "go inside"

	switch command {
	case "go east":
		fmt.Println("You head further up the mountain.")
	case "enter cave", "go inside":
		fmt.Println("You find yourself in a dimly lit cavern.")
		//在JS中 如果不写break会直接下降 即 直接运行下面所有的代码
		//但是Go更为谨慎 不写fallthrough就不会这样了
		fallthrough
	case "read sign":
		fmt.Println("The sign reads 'No Minors'.")
	default:
		fmt.Println("Didn't quite get that.")
	}
}
