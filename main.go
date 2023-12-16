package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print()
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	result, err := calculateResult(input)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", result)
	}
}

func calculateResult(input string) (string, error) {
	parts := strings.Fields(input)
	if len(parts) != 3 {
		return "", errors.New("неверный формат")
	}

	a, errA := convertInput(parts[0])
	b, errB := convertInput(parts[2])
	operator := parts[1]

	if errA != nil || errB != nil {
		return "", errors.New("неверный формат")
	}

	if (isRomanNumeral(parts[0]) && isArabicNumeral(parts[2])) || (isArabicNumeral(parts[0]) && isRomanNumeral(parts[2])) {
		return "", errors.New("неверный формат")
	}

	var result int
	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			return "", errors.New("неверный формат")
		}
		result = a / b
	default:
		return "", errors.New("неверный формат")
	}

	if isRomanNumeral(parts[0]) && isRomanNumeral(parts[2]) {
		if result <= 0 {
			return "", errors.New("неверный формат")
		}
		return toRomanNumeral(result), nil
	}
	return strconv.Itoa(result), nil
}

func convertInput(input string) (int, error) {
	if isRomanNumeral(input) {
		arabicNumeral, err := fromRomanNumeral(input)
		if err != nil {
			return 0, err
		}
		return arabicNumeral, nil
	}

	number, err := strconv.Atoi(input)
	if err != nil {
		return 0, errors.New("неверный формат")
	}
	if number < 1 || number > 10 {
		return 0, errors.New("неверный формат")
	}
	return number, nil
}

func isRomanNumeral(input string) bool {
	romanNumerals := map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
		"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	}
	_, exists := romanNumerals[input]
	return exists
}

func isArabicNumeral(input string) bool {
	_, err := strconv.Atoi(input)
	return err == nil
}

func toRomanNumeral(n int) string {
	if n <= 0 || n > 3999 {
		return "неверный формат"
	}

	var result strings.Builder
	romanNumerals := []struct {
		Value  int
		Symbol string
	}{
		{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
		{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
		{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
	}

	for _, rn := range romanNumerals {
		for n >= rn.Value {
			result.WriteString(rn.Symbol)
			n -= rn.Value
		}
	}
	return result.String()
}

func fromRomanNumeral(s string) (int, error) {
	romanNumerals := map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
		"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	}
	numeral, exists := romanNumerals[s]
	if !exists {
		return 0, errors.New("неверный формат")
	}
	return numeral, nil
}
