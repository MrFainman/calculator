package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	arithmetic_expression, _ := reader.ReadString('\n')
	expression_parts := strings.Split(arithmetic_expression, " ")
	if len(expression_parts) != 3 {
		panic("Введенная строка не соответствует формату a + b")
	} else if expression_parts[1] != "+" && expression_parts[1] != "-" && expression_parts[1] != "*" && expression_parts[1] != "/" {
		panic("Доступны только операторы +, -, *, /.")
	}
	var first_numeral, second_numeral, result int
	var first_numeral_is_roman, second_numeral_is_roman bool
	expression_parts[1] = strings.TrimSpace(expression_parts[1])
	expression_parts[2] = strings.TrimSpace(expression_parts[2])
	switch expression_parts[0] {
	case "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X":
		first_numeral_is_roman = true
	}
	switch expression_parts[2] {
	case "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X":
		second_numeral_is_roman = true
	}
	var both_roman bool
	if first_numeral_is_roman != second_numeral_is_roman {
		panic("Используются одновременно разные системы счисления или вы ввели или вы ввели слишком большое римское число.")
	} else if first_numeral_is_roman && second_numeral_is_roman {
		both_roman = true
	}
	first_numeral_str, second_numeral_str := expression_parts[0], expression_parts[2]
	first_numeral, err1 := strconv.Atoi(first_numeral_str)
	second_numeral, err2 := strconv.Atoi(second_numeral_str)
	if (err1 == nil && (first_numeral < 1 || first_numeral > 10)) || (err2 == nil && (second_numeral < 1 || second_numeral > 10)) {
		panic("Числа должны быть в диапазоне от 1 до 10")
	}
	switch expression_parts[0] {
	case "I", "1":
		first_numeral = 1
	case "II", "2":
		first_numeral = 2
	case "III", "3":
		first_numeral = 3
	case "IV", "4":
		first_numeral = 4
	case "V", "5":
		first_numeral = 5
	case "VI", "6":
		first_numeral = 6
	case "VII", "7":
		first_numeral = 7
	case "VIII", "8":
		first_numeral = 8
	case "IX", "9":
		first_numeral = 9
	case "X", "10":
		first_numeral = 10
	}
	switch expression_parts[2] {
	case "I", "1":
		second_numeral = 1
	case "II", "2":
		second_numeral = 2
	case "III", "3":
		second_numeral = 3
	case "IV", "4":
		second_numeral = 4
	case "V", "5":
		second_numeral = 5
	case "VI", "6":
		second_numeral = 6
	case "VII", "7":
		second_numeral = 7
	case "VIII", "8":
		second_numeral = 8
	case "IX", "9":
		second_numeral = 9
	case "X", "10":
		second_numeral = 10
	}
	switch expression_parts[1] {
	case "+":
		result = first_numeral + second_numeral
	case "-":
		result = first_numeral - second_numeral
	case "*":
		result = first_numeral * second_numeral
	case "/":
		result = first_numeral / second_numeral
	}
	if both_roman {
		if result < 1 {
			panic("Не бывает римских чисел меньше нуля.")
		} else {
			roman_numerals := map[int]string{
				1: "I", 4: "IV", 5: "V", 9: "IX", 10: "X", 40: "XL", 50: "L", 90: "XC", 100: "C",
			}
			keys := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
			roman_result := ""
			for _, key := range keys {
				for result >= key {
					result -= key
					roman_result += roman_numerals[key]
				}
			}
			fmt.Println(roman_result)
		}
	} else {
		fmt.Println(result)
	}
}
