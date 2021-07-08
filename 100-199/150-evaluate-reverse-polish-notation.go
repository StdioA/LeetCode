package main

import (
	"fmt"
	"strconv"
)

func isOperator(token string) bool {
	operators := []string{"+", "-", "*", "/"}
	for i := 0; i < 4; i++ {
		if token == operators[i] {
			return true
		}
	}
	return false
}

func evaluate(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	}
	fmt.Printf("%d %s %d = %d\n", a, op, b, result)
	return result
}

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, token := range tokens {
		if !isOperator(token) {
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		} else {
			l := len(stack)
			a, b := stack[l-2], stack[l-1]
			result := evaluate(a, b, token)
			stack[l-2] = result
			stack = stack[:l-1]
		}
	}
	return stack[0]
}

func main() {
	ops := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(evalRPN(ops))
}
