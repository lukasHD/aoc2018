package d01

import (
	"fmt"
	"lukas/aoc2018/commons"
	"strconv"
)

func Test() int {
	name := "Bond"
	ress, _ := commons.MySum(21, 42)

	message := fmt.Sprintf("Hello Mr. %v! Welcome to my BatCave. You are visitor %d", name, ress)

	fmt.Println(message)

	return 0
}

func QuickRun1(input string) int {
	frequencyList := commons.SplitStringToList(input)
	fmt.Println(frequencyList)
	result := 0
	for _, freq := range frequencyList {
		i_freq, _ := strconv.Atoi(freq)
		result += i_freq
	}

	return result
}

func FullRun1() {
	lines := commons.LoadLines("./d01/input.txt")

	result := 0
	for _, freq := range lines {
		i_freq, _ := strconv.Atoi(freq)
		result += i_freq
	}

	fmt.Printf("Day 01 - Part 1: %d \n", result)
}

func QuickRun2(input []string) int {
	seenFrequencies := make(map[int]int)
	value := 0
	seenFrequencies[0] += 1
	for {
		for _, freq := range input {
			if freq == "" {
				continue
			}
			i_freq, _ := strconv.Atoi(freq)
			value += i_freq
			//fmt.Print(value)
			//fmt.Print(", ")
			_, ok := seenFrequencies[value]
			if ok {
				//fmt.Println()
				return value
			}
			seenFrequencies[value] += 1
		}
		//fmt.Println()
		//fmt.Println()
	}

}

func FullRun2() {
	lines := commons.LoadLines("./d01/input.txt")

	result := QuickRun2(lines)

	fmt.Printf("Day 01 - Part 2: %d \n", result)
}
