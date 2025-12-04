package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../inputs/1dec.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	inputs := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		dir := line[0]
		num, _ := strconv.Atoi(line[1:])
		if dir == 'L' {
			inputs = append(inputs, -1*num)
		} else {
			inputs = append(inputs, num)
		}
	}
	res := solve2(inputs)
	fmt.Println("Result:", res)
}

func solve1(inputs []int) int {
	dial := 50
	code := 0
	for i := range inputs {
		dial = (dial + inputs[i]) % 100
		if dial == 0 {
			code += 1
		}
	}
	return code
}

func solve2(inputs []int) int {
	dial := 50
	code := 0
	// inputs := []int{-68, -30, 48, -5, 60, -55, -1, -99, 14, -82}

	for i := range inputs {
		newNum := dial + inputs[i]
		// fmt.Println("change: ", inputs[i], " result: ", newNum)
		if newNum <= 0 && dial != 0 {
			code += 1
			//	fmt.Println("newNum <= 0 and dial !=0. -> new code: ", code)
		}
		code += abs(newNum) / 100
		// fmt.Println("new code: ", code)

		dial = mod(dial+inputs[i], 100)
		// fmt.Println("new dial: ", dial)
	}
	return code
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func mod(x, y int) int {
	return (x%y + y) % y
}
