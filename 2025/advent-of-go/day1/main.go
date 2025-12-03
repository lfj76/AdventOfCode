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
	res := solve(inputs)
	fmt.Println("Result:", res)
}

func solve(inputs []int) int {
	return 0
}
