package main

import (
	"fmt"
	"lukas/aoc2018/d01"
	"lukas/aoc2018/d02"
)

func greeting() int {
	fmt.Println("Hello Everybody")
	return 1
}

func runAll() {
	d01.FullRun1()
	d01.FullRun2()
	d02.FullRun1()
	//d02.QuickRun2()
	d02.FullRun2()
}

func main() {
	greeting()
	runAll()
}
