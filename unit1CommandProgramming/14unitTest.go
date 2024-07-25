package main

import (
	"fmt"
	"math/rand"
)

func geneSpeed() int {
	return rand.Intn(15) + 16
}

func genePrice() int {
	return rand.Intn(15) + 36
}

func main() {
	const distance = 62100000

	fmt.Printf("%-20s%10s%12s%16s\n", "太空航行公司", "飞行天数", "飞行类型", "价格(百万美元)")
	for count := 10; count > 0; count-- {
		number := rand.Intn(3) // 改成 3，以包含三个公司
		var ticketType string
		if rand.Intn(2) == 0 {
			ticketType = "单程"
		} else {
			ticketType = "往返"
		}

		var companyName string
		switch number {
		case 0:
			companyName = "Virgin Galactic"
		case 1:
			companyName = "SpaceX"
		case 2:
			companyName = "Space Adventures"
		}

		days := distance / geneSpeed() / 60 / 60 / 24 // 计算飞行天数
		price := genePrice()
		if ticketType == "往返" {
			price *= 2
		}

		fmt.Printf("%-20s%10d%12s%16d\n", companyName, days, ticketType, price)
	}
}
