package practice

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	var stack []int
	for _, c := range tokens {
		switch c {
		case "+":
			if ls := len(stack); ls > 1 {
				stack[ls-2] += stack[ls-1]
				stack = stack[:ls-1]
			}
		case "-":
			if ls := len(stack); ls > 1 {
				stack[ls-2] -= stack[ls-1]
				stack = stack[:ls-1]
			}
		case "*":
			if ls := len(stack); ls > 1 {
				stack[ls-2] *= stack[ls-1]
				stack = stack[:ls-1]
			}
		case "/":
			if ls := len(stack); ls > 1 {
				stack[ls-2] /= stack[ls-1]
				stack = stack[:ls-1]
			}
		default:
			num, err := strconv.Atoi(c)
			if err == nil {
				stack = append(stack, num)
			}
		}
	}
	return stack[len(stack)-1]
}

func evalRPN1(tokens []string) int {
	var stack []int
	for _, c := range tokens {
		ls := len(stack)

		switch c {
		case "+":
			if  ls > 1 {
				stack = append(stack[:ls-2],stack[ls-2] + stack[ls-1])
			}
		case "-":
			if  ls > 1 {
				stack = append(stack[:ls-2],stack[ls-2] - stack[ls-1])
			}
		case "*":
			if  ls > 1 {
				stack = append(stack[:ls-2],stack[ls-2] * stack[ls-1])
			}
		case "/":
			if  ls > 1 {
				stack = append(stack[:ls-2],stack[ls-2] / stack[ls-1])
			}
		default:
			num, err := strconv.Atoi(c)
			if err == nil {
				stack = append(stack, num)
			}
		}
	}
	return stack[0]
}
