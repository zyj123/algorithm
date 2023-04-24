package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//chars1 := [][]byte{
	//	{'1', ' ', '3'},
	//	{'-', '+', '-'},
	//	{'2', ' ', '4'},
	//}
	//chars2 := [][]byte{
	//	{' ', ' ', '2', ' ', ' '},
	//	{'1', '+', '-', '/', '2'},
	//	{' ', ' ', '3', ' ', ' '},
	//}

	// read from file
	// 假设字符串是有效的表达式
	chars1 := readFile("./example1.txt")
	chars2 := readFile("./example2.txt")
	fmt.Println(chars1, chars2)

	fmt.Println(calculator(chars1) == "5/4")
	fmt.Println(calculator(chars2) == "4/3")
}

func readFile(file string) [][]byte {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	res := make([][]byte, 0)
	for scanner.Scan() {
		res = append(res, scanner.Bytes())
	}
	if err = scanner.Err(); err != nil {
		panic(err)
	}
	if len(res) != 3 || len(res[0]) != len(res[2]) {
		panic(fmt.Errorf("file is illegal"))
	}

	return res
}

func calculator(chars [][]byte) string {
	stack := make([][]int, 0)
	var preSign byte = '+'
	var (
		line1 = chars[0]
		line2 = chars[1]
		line3 = chars[2]
	)
	for i := 0; i < len(line2); i++ {
		// 符号位
		if (i >= len(line1) && !isDigit(line2[i])) || (i < len(line1) && line1[i] == ' ' && !isDigit(line2[i])) {
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
	res = simplify(res)
	return fmt.Sprintf("%d/%d", res[0], res[1])
}

func toInt(i byte) int {
	return int(i - '0')
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

// 分数计算 []int{分子，分母}
func add(a, b []int) []int {
	denominator := a[1] * b[1]
	numeratorA := a[0] * b[1]
	numeratorB := b[0] * a[1]
	return simplify([]int{numeratorA + numeratorB, denominator})
}
func sub(a, b []int) []int {
	denominator := a[1] * b[1]
	numeratorA := a[0] * b[1]
	numeratorB := b[0] * a[1]
	return simplify([]int{numeratorA - numeratorB, denominator})
}
func multi(a, b []int) []int {
	numerator := a[0] * b[0]
	denominator := a[1] * b[1]
	return simplify([]int{numerator, denominator})
}
func divide(a, b []int) []int {
	numerator := a[0] * b[1]
	denominator := a[1] * b[0]
	return simplify([]int{numerator, denominator})
}

// 化简分数
func simplify(a []int) []int {
	g := gcd(a[0], a[1])
	return []int{a[0] / g, a[1] / g}
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
