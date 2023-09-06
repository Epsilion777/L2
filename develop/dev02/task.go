package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// copyChar копирует указанный символ в результирующую строку count раз
func copyChar(count, result *strings.Builder, char *rune) error {
	if *char != ' ' {
		count, _ := strconv.Atoi(count.String())
		if count == 0 {
			count = 1
		}

		for i := 0; i < count; i++ {
			result.WriteRune(*char)
		}
	}
	*char = ' '
	count.Reset()
	return nil
}

// Unpacking распаковывает строку с учетом escape-последовательностей
func Unpacking(str string) (string, error) {
	runes := []rune(str)
	result, count := strings.Builder{}, strings.Builder{}
	lastChar := ' '
	for i := 0; i < len(runes); i++ {
		elem := string(runes[i])

		if (elem < "0" || elem > "9") && elem != `\` {
			err := copyChar(&count, &result, &lastChar)
			if err != nil {
				return "", fmt.Errorf("unexpected error: %w", err)
			}
			lastChar = runes[i]
			continue
		}

		if "0" < elem && elem < "9" && lastChar != ' ' {
			count.WriteRune(runes[i])
			continue
		}

		if elem == `\` && i+1 < len(runes) {
			err := copyChar(&count, &result, &lastChar)
			if err != nil {
				return "", fmt.Errorf("unexpected error: %w", err)
			}
			i++
			lastChar = runes[i]
		}
	}
	// Копируем последний необработанный элемент
	err := copyChar(&count, &result, &lastChar)
	if err != nil {
		return "", fmt.Errorf("unexpected error: %w", err)
	}

	if result.Len() == 0 && str != "" {
		return "", fmt.Errorf("invalid string")
	}
	return result.String(), nil
}

func main() {
	str := "a4bc2d5e"
	res, err := Unpacking(str)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Printf("%s -> %s\n", str, res)
	}
}
