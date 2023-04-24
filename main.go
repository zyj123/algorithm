package main

import "fmt"

func main() {
	// read from file
	// 假设字符串是有效的表达式
	chars := [][]byte{
		{'1', ' ', '3'},
		{'-', '+', '-'},
		{'2', ' ', '4'},
	}

	calculator(chars)
}

func calculator(chars [][]byte) string {
	stack := make([][]int, 0)
	var preSign byte = '+'
	n := len(chars[0])

	for i := 0; i < n; i++ {
		var (
			line1 = chars[0]
			line2 = chars[1]
			line3 = chars[2]
		)
		// 符号位
		if line1[i] == ' ' && !isDigit(line2[i]) {
			preSign = line2[i]
			continue
		}
		// 值位
		val := make([]int, 0)
		if isDigit(line2[i]) {
			val = []int{toInt(line2[i]), 1}
		} else {
			val = []int{toInt(line1[i]), toInt(line3[i])}
		}

		switch preSign {
		case '+':
			stack = append(stack, val)
		case '-':
			zeroVal := []int{0, 1}
			stack = append(stack, sub(zeroVal, val))
		case '*':
			stack[len(stack)-1] = multi(stack[len(stack)-1], val)
		case '/':
			stack[len(stack)-1] = divide(stack[len(stack)-1], val)
		}
	}
	res := []int{0, 1}
	for _, v := range stack {
		res = add(res, v)
	}
	// 计算分数的值
	fmt.Println(res)

	return ""
}

func toInt(i byte) int {
	return int(i - '0')
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

// 分数计算 []int{分子，分母}
func add(a, b []int) []int {
	denominatorLcm := lcm(a[1], b[1])
	numeratorA := a[0] * (denominatorLcm / a[1])
	numeratorB := b[0] * (denominatorLcm / b[1])
	return []int{numeratorA + numeratorB, denominatorLcm}
}
func sub(a, b []int) []int {
	denominatorLcm := lcm(a[1], b[1])
	numeratorA := a[0] * (denominatorLcm / a[1])
	numeratorB := b[0] * (denominatorLcm / b[1])
	return []int{numeratorA - numeratorB, denominatorLcm}
}
func multi(a, b []int) []int {
	numerator := a[0] * b[0]
	denominator := a[1] * b[1]
	return []int{numerator, denominator}
}
func divide(a, b []int) []int {
	numerator := a[0] * b[1]
	denominator := a[1] * b[0]
	return []int{numerator, denominator}
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}
