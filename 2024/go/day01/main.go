package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func part1(scanner *bufio.Scanner) {
	left := []int{}
	right := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		num, err := strconv.Atoi(fields[0])
		if err != nil {
			fmt.Println("Failed to convert string to int")
			os.Exit(1)
		}
		left = append(left, num)

		num, err = strconv.Atoi(fields[1])
		if err != nil {
			fmt.Println("Failed to convert string to int")
			os.Exit(1)
		}
		right = append(right, num)
	}

	sort.Ints(left)
	sort.Ints(right)

	total := 0
	for i := range left {
		diff := left[i] - right[i]
		total += int(math.Abs(float64(diff)))
	}
	fmt.Printf("Part 1: %d\n", total)
}

func part2(scanner *bufio.Scanner) {

}

func main() {
	numArgs := len(os.Args)
	if numArgs < 2 {
		fmt.Printf("Received invalid number of inputs for the filepath, expects: 1 received: %d\n", numArgs)
		os.Exit(1)
	}

	filename := os.Args[1]
	partNumber := 0
	if numArgs > 2 {
		var err error
		partNumber, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Failed to parse part number, running all parts")
			partNumber = 0
		}
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if partNumber == 0 || partNumber == 1 {
		part1(scanner)
	}
	if partNumber == 0 || partNumber == 2 {
		part2(scanner)
	}

}
