package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func cycleLength(num int) int {
	steps := 1

	for num > 1 {
		if num%2 == 0 {
			num = num / 2
		} else {
			num = 3*num + 1
		}

		steps++
	}

	return steps
}

func maxCycleLength(i int, j int) string {
	max := cycleLength(i)

	for k := i + 1; k <= j; k++ {
		current := cycleLength(k)

		if current > max {
			max = current
		}
	}

	return fmt.Sprintf("%d %d %d", i, j, max)
}

// Sample Input
// 1 10
// 100 200
// 201 210
// 900 1000

// Sample OutPut
// 1 	10	  20
// 100 	200	  125
// 201 	210   89
// 900 	1000  174

func main() {
	fmt.Println("3n + 1 problem, input numbers: ")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()
		if len(input) == 0 {
			break
		}

		values := strings.Fields(input)
		i, err := strconv.Atoi(values[0])
		if err != nil {
			fmt.Printf("Error unknow on convert i %s", err.Error())
		}

		j, err := strconv.Atoi(values[1])
		if err != nil {
			fmt.Printf("Error unknow on convert j %s", err.Error())
		}

		fmt.Println(maxCycleLength(i, j))
	}

	if err := scanner.Err(); err != nil {
		if !errors.Is(io.EOF, err) {
			fmt.Fprintln(os.Stderr, err.Error())
		}
	}
}
