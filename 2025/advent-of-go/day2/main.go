package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputStrings := readFileLines("../inputs/day2.txt")
	res := solve2(inputStrings)
	fmt.Println("Result:", res)
}

func solve(inputs []string) int {
	res := 0
	for i := range inputs {
		from, to, _ := parseInput(inputs[i])
		for j := from; j <= to; j++ {
			if numLen(j)%2 != 0 {
				continue
			}
			l, r := splitNum(j)
			if l == r {
				res += j
			}
		}
	}
	return res
}

func solve2(inputs []string) int {
	res := 0
	for i := range inputs {
		from, to, _ := parseInput(inputs[i])
		for j := from; j <= to; j++ {
			if hasRepeat(strconv.Itoa(j)) {
				res += j
			}
		}
	}
	return res
}

func hasRepeat(s string) bool {
	x, xs := parse1(s)
	if x == '\x00' {
		return false
	}

	return checkRepeat(string(x), xs)
}

func checkRepeat(x, xs string) bool {
	if len(x) > len(xs) {
		return false
	}
	if x == xs {
		return true
	}
	if x == xs[:len(x)] {
		// check if x repeats further
		if checkRepeat(x, xs[len(x):]) {
			return true
		}
	}
	return checkRepeat(x+string(xs[0]), xs[1:])
}

func parse1(s string) (byte, string) {
	switch l := len(s); l {
	case 0:
		return '\x00', ""
	case 1:
		return s[0], ""
	default:
		return (s[0]), s[1:]
	}
}

func readFileLines(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return []string{}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
	}
	return strings.Split(line, ",")
}

func parseInput(input string) (from, to int, err error) {
	pair := strings.Split(input, "-")
	num1, err1 := strconv.Atoi(pair[0])
	if err1 != nil {
		return 0, 0, err1
	}
	num2, err2 := strconv.Atoi(pair[1])
	if err2 != nil {
		return 0, 0, err2
	}
	return num1, num2, nil
}

func numLen(n int) int {
	return len(strconv.Itoa(n))
}

func splitNum(n int) (left, right int) {
	s := strconv.Itoa(n)
	l := s[:len(s)/2]
	r := s[len(s)/2:]

	num1, _ := strconv.Atoi(l)
	num2, _ := strconv.Atoi(r)
	return num1, num2
}
