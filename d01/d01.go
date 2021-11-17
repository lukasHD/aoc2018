package d01

import (
	"fmt"
	"lukas/aoc2018/commons"
)

func Test() int {
	name := "Bond"
	message := fmt.Sprintf("Hello Mr. %v! Welcome to my BatCave.", name)

	commons.MySum(21, 42)

	fmt.Println(message)

	return 0
}
