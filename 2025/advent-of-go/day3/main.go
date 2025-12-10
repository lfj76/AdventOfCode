package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var debug bool

func init() {
	flag.BoolVar(&debug, "debug", false, "enable debug mode")
}

func main() {
	flag.Parse()
	file, err := os.Open("../inputs/day3.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	inputs := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var nums []int
		for _, ch := range line {
			num := int(ch - '0')
			nums = append(nums, num)
		}
		inputs = append(inputs, nums)
		debugPrint(nums)
	}
	res := solve(inputs)
	fmt.Println(res)
}

func solve(input [][]int) int {
	res := 0
	for _, array := range input {
		indexLargest := findIndexLargest(array[:len(array)-1])
		indexSecond := findIndexLargest(array[indexLargest+1:]) + indexLargest + 1
		sum := mergeNums(array[indexLargest], array[indexSecond])
		debugPrint(fmt.Sprintf("index of largest: %d, index of second: %d, sum: %d", indexLargest, indexSecond, sum))
		res += sum
	}
	return res
}

func mergeNums(i1, i2 int) int {
	return i1*10 + i2
}

func findIndexLargest(array []int) int {
	largest := 0
	index := 0
	for i, val := range array {
		if val > largest {
			largest = val
			index = i
		}
	}
	return index
}

func debugPrint(s any) {
	if debug {
		fmt.Println(s)
	}
}
