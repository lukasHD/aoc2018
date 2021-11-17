package main

import (
	"fmt"
	"lukas/aoc2018/d01"
	"lukas/aoc2018/helper"
)

func greeting() int {
	fmt.Println("Hello Everybody")
	return 1
}

func main() {
	greeting()
	helper.Greet()
	d01.Test()
	d01.FullRun1()
}
