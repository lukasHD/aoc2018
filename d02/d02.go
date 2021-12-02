package d02

import (
	"bufio"
	"fmt"
	"log"
	"lukas/aoc2018/commons"
	"os"
)

func getRepeatCountLetters(row string) (bool, bool) {
	m := make(map[string]int)
	var twoCount = false
	var threeCount = false

	for _, val := range row {
		m[string(val)]++
	}

	for k := range m {
		if m[k] == 2 {
			twoCount = true
		}
		if m[k] == 3 {
			threeCount = true
		}
	}
	return twoCount, threeCount
}

func FullRun1() int {
	var result int
	twoCount := 0
	threeCount := 0
	file, err := os.Open("./d02/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		containsTwo, containsThree := getRepeatCountLetters(row)
		if containsTwo {
			twoCount++
		}
		if containsThree {
			threeCount++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	result = twoCount * threeCount
	fmt.Printf("Day 02 - Part 1: %d \n", result)
	return result
}

func commonCharacters(a, b string) (int, string) {
	if len(a) == len(b) {
		counter := 0
		result := ""
		for i := range a {
			if a[i] == b[i] {
				result = result + string(a[i])
			} else {
				counter += 1
			}

		}
		return counter, result

	} else {
		log.Fatal()
	}
	return -1, ""
}

func compare(fname string, pretty bool) string {
	rows := commons.LoadLines(fname)

	for i, rowA := range rows {
		for j := i + 1; j < len(rows); j++ {
			rowB := rows[j]
			cnt, common := commonCharacters(rowA, rowB)
			if pretty {
				fmt.Println(rowA + "  " + rowB + " cnt=" + fmt.Sprint(cnt) + "  common=" + common)
			}
			if cnt == 1 {
				return common
			}
		}
	}
	return ""
}

func QuickRun2() string {
	var result string = ""
	result = compare("./d02/test.txt", true)
	fmt.Printf("Day 02 - Part 2 TEST: %s \n", result)
	return result
}

func FullRun2() string {
	var result string = ""
	result = compare("./d02/input.txt", false)
	fmt.Printf("Day 02 - Part 2: %s \n", result)
	return result
}
