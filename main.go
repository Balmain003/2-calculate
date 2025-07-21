package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func calculateAVG(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	sum := 0.0
	for _, n := range numbers {
		sum = sum + n
	}
	return sum / float64(len(numbers))
}

func calculateSUM(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	sum := 0.0
	for _, n := range numbers {
		sum = sum + n
	}
	return sum
}

func calculateMED(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	sort.Float64s(numbers)
	mid := len(numbers) / 2
	if len(numbers)%2 == 0 {
		return (numbers[mid-1] + numbers[mid]) / 2
	}
	return numbers[mid]
}
func splittingStrings(input string) ([]float64, error) {
	parts := strings.Split(input, ",")
	numbers := make([]float64, 0, len(parts))
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed == "" {
			continue
		}
		num, err := strconv.ParseFloat(trimmed, 64)
		if err != nil {
			return nil, fmt.Errorf("неверное число: %s", trimmed)
		}
		numbers = append(numbers, num)
	}
	if len(numbers) == 0 {
		return nil, fmt.Errorf("не указаны числа")
	}
	return numbers, nil
}
func readOperation() string {
	var input string
	fmt.Print("Выберите операцию (AVG / SUM / MED / exit): ")
	fmt.Scan(&input)
	input = strings.ToUpper(strings.TrimSpace(input))
	if input == "AVG" || input == "SUM" || input == "MED" {
		return input
	}
	return "Ошибка неизвестная операция"
}
func readNumbers() string {
	var input string
	fmt.Scan(&input)
	input = strings.TrimSpace(input)
	if input != "" {
		return input
	}
	return "Oшибка введите хотябы одно число"
}
func main() {
	operation := readOperation()
	fmt.Print("Введите числа через запятую: ")
	numbersStr := readNumbers()
	numbers, err := splittingStrings(numbersStr)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
	var result float64

	switch operation {
	case "AVG":
		result = calculateAVG(numbers)
	case "SUM":
		result = calculateSUM(numbers)
	case "MED":
		result = calculateMED(numbers)
	default:
		fmt.Printf("Нет такой операции %s", operation)
		return
	}
	fmt.Printf("Результат вычислений операции %s = %.0f", operation, result)
}
