package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	after      int
	before     int
	context    int
	count      bool
	ignoreСase bool
	invert     bool
	fixed      bool
	lineNum    bool
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Парсинг флагов командной строки
func init() {
	flag.IntVar(&after, "A", 0, "Печатать +N строк после совпадения")
	flag.IntVar(&before, "B", 0, "Печатать +N строк до совпадения")
	flag.IntVar(&context, "C", 0, "(A+B) печатать ±N строк вокруг совпадения")
	flag.BoolVar(&count, "c", false, "Количество строк")
	flag.BoolVar(&ignoreСase, "i", false, "Игнорировать регистр")
	flag.BoolVar(&invert, "v", false, "Вместо совпадения, исключать")
	flag.BoolVar(&fixed, "F", false, "Точное совпадение со строкой, не паттерн")
	flag.BoolVar(&lineNum, "n", false, "Печатать номер строки")

	flag.Parse()
}

// Удаляет элементы из slice1, которые входят в slice2
func removeElements(slice1, slice2 []int) []int {
	result := make([]int, 0)

	// Создаем мапу для быстрой проверки наличия элементов из slice2
	existingElements := make(map[int]struct{})
	for _, v := range slice2 {
		existingElements[v] = struct{}{}
	}

	// Проходим по slice1 и добавляем только те элементы, которых нет в existingElements
	for _, v := range slice1 {
		if _, exist := existingElements[v]; !exist {
			result = append(result, v)
		}
	}

	return result
}

// Функция для генерирации интервала для получения индексов нужных строк с учетом флагов (A, B, C)
func interval(slc []int, countBeforstr, countAfterStr, size int) []int {
	newSlc := make([]int, 0, 2*(countBeforstr+countAfterStr)+len(slc))
	lastInserted := -1

	for _, v := range slc {
		if v-countBeforstr > 0 {
			for i := v - countBeforstr; i <= v; i++ {
				if i > lastInserted {
					newSlc = append(newSlc, i)
					lastInserted = i
				}

			}
		} else {
			for i := 0; i <= v; i++ {
				if i > lastInserted {
					newSlc = append(newSlc, i)
					lastInserted = i
				}

			}
		}

		if v+countAfterStr >= size {
			for i := v; i < size; i++ {
				if i > lastInserted {
					newSlc = append(newSlc, i)
					lastInserted = i
				}
			}
		} else {
			for i := v; i <= v+countAfterStr; i++ {
				if i > lastInserted {
					newSlc = append(newSlc, i)
					lastInserted = i
				}
			}
		}
	}
	// Если флаг активен, то необходимо удалить все строки, которые соответствуют паттерну
	if invert {
		return removeElements(newSlc, slc)
	}
	return newSlc
}

// GrepFunc - утилита grep
func GrepFunc(fileData []byte, regExp *regexp.Regexp) []string {
	// Разбиваем файл на строки
	lines := strings.Split(string(fileData), "\r\n")

	// Ищем совпадения и запомниаем индекс строки
	slcOfIndexes := make([]int, 0, len(lines))
	for i, line := range lines {
		if regExp.Match([]byte(line)) {
			slcOfIndexes = append(slcOfIndexes, i)
		}
	}

	countBeforStr := max(before, context)
	countAfterStr := max(after, context)

	// Генерируем интервал для вывода нужных строк с учетом флагов (A, B, C)
	slcOfIndexes = interval(slcOfIndexes, countBeforStr, countAfterStr, len(lines))
	resultStrs := make([]string, 0, len(slcOfIndexes))
	// Добавляем в результирующий массив строк строки с учетом флага lineNum
	for _, v := range slcOfIndexes {
		if lineNum {
			resultStrs = append(resultStrs, fmt.Sprintf("%d: %s", v, lines[v]))
		} else {
			resultStrs = append(resultStrs, lines[v])
		}
	}
	// Добавляем в результирующий массив строк строку с кол-вом строк
	if count {
		resultStrs = append(resultStrs, fmt.Sprintf("Количество строк: %d", len(resultStrs)))
	}
	return resultStrs
}

func main() {

	// Создание необходимого паттерна на основе флагов
	pattern := flag.Args()[0]
	if fixed {
		pattern = `\Q` + pattern + `\E`
	}

	if ignoreСase {
		pattern = `(?i)` + pattern
	}

	// Генерируем регулярное выражение
	regExp := regexp.MustCompile(pattern)

	fileData, err := os.ReadFile("develop/dev05/input1.txt")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	// Запускаем утилиту
	resultStrs := GrepFunc(fileData, regExp)

	// Выводим результат
	for _, v := range resultStrs {
		fmt.Println(v)
	}

}
