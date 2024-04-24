package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isArabic(s string) bool {
	s = strings.TrimSpace(s)
	if s == "" {
		return false
	}
	_, err := strconv.Atoi(s)
	return err == nil
}

func arabicToRoman(n int) string {
	romanMap := []struct {
		Value  int
		Symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var result strings.Builder
	for _, pair := range romanMap {
		count := n / pair.Value
		for i := 0; i < count; i++ {
			result.WriteString(pair.Symbol)
		}
		n %= pair.Value
	}
	return result.String()
}

func romanToArabic(s string) int {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}

	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	prevValue := 0
	for i := 0; i < len(s); i++ {
		value, ok := romanMap[s[i]]
		if !ok {
			return 0
		}

		if value > prevValue {
			result += value - 2*prevValue
		} else {
			result += value
		}
		prevValue = value
	}

	return result
}

func parseNumber(s string) (int, bool, bool) {
	s = strings.TrimSpace(s)
	n, err := strconv.Atoi(s)
	if err == nil {
		return n, true, true
	}

	number := romanToArabic(s)
	if number != 0 {
		return number, true, false
	}

	return 0, false, false
}

func calculate(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	}
	panic("Неверная операция")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите выражение (например, 2 + 3): ")
	input, _ := reader.ReadString('\n')
	parts := strings.Split(strings.TrimSuffix(input, "\r\n"), " ")
	if len(parts) != 3 {
		panic("Неверный формат входных данных")
	}

	a, isNumberA, isArabicA := parseNumber(parts[0])
	b, isNumberB, isArabicB := parseNumber(parts[2])

	if !isNumberA || !isNumberB {
		panic("Ошибка при разборе чисел")
	}

	if isArabicA != isArabicB {
		panic("Используются разные системы счисления")
	}

	if a < 1 || a > 10 || b < 1 || b > 10 {
		panic("Числа должны быть от 1 до 10")
	}

	op := parts[1]
	result := calculate(a, b, op)

	if isArabicA {
		fmt.Print(result)
	} else {
		if result <= 0 {
			panic("В римской системе нет нуля и отрицательных чисел")
		}
		fmt.Printf(arabicToRoman(result))
	}
}
