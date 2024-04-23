package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Определение отображения арабских цифр на римские
var romanNumerals = []struct {
	Value  int
	Symbol string
}{
	{10, "X"},
	{9, "IX"},
	{8, "VIII"},
	{7, "VII"},
	{6, "VI"},
	{5, "V"},
	{4, "IV"},
	{3, "III"},
	{2, "II"},
	{1, "I"},
}

// Функция для конвертации арабских чисел в римские
func arabicToRoman(n int) string {
	for _, value := range romanNumerals {
		if n >= value.Value {
			return value.Symbol + arabicToRoman(n-value.Value)
		}
	}
	return ""
}

// Функция для конвертации римских чисел в арабские
func romanToArabic(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		symbol := string(s[i])
		for _, value := range romanNumerals {
			if symbol == value.Symbol {
				if i+1 < len(s) {
					nextSymbol := string(s[i+1])
					for _, next := range romanNumerals {
						if nextSymbol == next.Symbol && next.Value > value.Value {
							result += next.Value - value.Value
							i++
							break
						}
					}
				}
				result += value.Value
				break
			}
		}
	}
	return result
}

// Определение функций для выполнения арифметических операций
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
	var input string
	fmt.Print("Введите выражение (например, 2 + 3): ")
	fmt.Scanln(&input)

	// Парсинг входных данных
	parts := strings.Split(input, " ")
	if len(parts) != 3 {
		panic("Неверный формат входных данных")
	}

	// Определение типа чисел (арабские или римские)
	isArabic := true
	a, err := strconv.Atoi(parts[0])
	if err != nil {
		a = romanToArabic(parts[0])
		isArabic = false
	}

	b, err := strconv.Atoi(parts[2])
	if err != nil {
		if isArabic {
			panic("Используются разные системы счисления")
		}
		b = romanToArabic(parts[2])
	} else if !isArabic {
		panic("Используются разные системы счисления")
	}

	// Проверка диапазона чисел
	if a < 1 || a > 10 || b < 1 || b > 10 {
		panic("Числа должны быть от 1 до 10")
	}

	// Вычисление результата
	op := parts[1]
	result := calculate(a, b, op)

	// Вывод результата
	if isArabic {
		fmt.Printf("%d %s %d = %d\n", a, op, b, result)
	} else {
		if result <= 0 {
			panic("В римской системе нет нуля и отрицательных чисел")
		}
		fmt.Printf("%s %s %s = %s\n", parts[0], op, parts[2], arabicToRoman(result))
	}
}
