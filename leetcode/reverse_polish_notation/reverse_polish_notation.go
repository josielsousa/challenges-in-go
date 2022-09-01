package main

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	ops := map[string]struct{}{"+": {}, "-": {}, "*": {}, "/": {}}

	stack := make([]int, 0)
	var num1, num2, tokInt int

	for _, tok := range tokens {
		if _, ok := ops[tok]; !ok {
			tokInt, _ = strconv.Atoi(tok)

			stack = append(stack, tokInt)
			continue
		}

		num2 = stack[len(stack)-1]
		num1 = stack[len(stack)-2]

		stack = stack[:len(stack)-2]

		switch tok {
		case "+":
			stack = append(stack, int(num1+num2))
		case "-":
			stack = append(stack, int(num1-num2))
		case "*":
			stack = append(stack, int(num1*num2))
		default:
			stack = append(stack, int(num1/num2))
		}
	}

	return stack[0]
}
