package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

# Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

# Дополнительное

# Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	column    int
	numeric   bool
	reverse   bool
	unique    bool
	inputFile string
)

// Парсинг флагов командной строки
func init() {
	flag.IntVar(&column, "k", 0, "Указание колонки для сортировки")
	flag.BoolVar(&numeric, "n", false, "Сортировать по числовому значению")
	flag.BoolVar(&reverse, "r", false, "Сортировать в обратном порядке")
	flag.BoolVar(&unique, "u", false, "Не выводить повторяющиеся строки")
	flag.StringVar(&inputFile, "f", "", "Указание имени сортируемого файла")

	flag.Parse()
}

// Получение массива строк из байтов (unique - оставлять ли повторяющиеся строки)
func getLinesFromBytes(bytes []byte, unique bool) []string {
	lines := strings.Split(string(bytes), "\r\n")

	if !unique {
		return lines
	}

	mapLines := make(map[string]struct{}, len(lines))
	for _, v := range lines {
		mapLines[v] = struct{}{}
	}
	uniqueLines := make([]string, 0, len(mapLines))
	for key := range mapLines {
		uniqueLines = append(uniqueLines, key)
	}
	return uniqueLines
}

// Создание матрицы слов из строк
func createMatrixFromLines(lines []string) [][]string {
	matrix := make([][]string, len(lines))
	countOfColumn := 0

	for i, v := range lines {
		matrix[i] = strings.Split(v, " ")
		if len(matrix[i]) > countOfColumn {
			countOfColumn = len(matrix[i])
		}
	}

	for i := range matrix {
		for len(matrix[i]) < countOfColumn {
			matrix[i] = append(matrix[i], "")
		}
	}

	return matrix
}

// Записать матрицу в файл
func writeMatrixToFile(matrix [][]string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, row := range matrix {
		rowStr := strings.Join(row, " ") // Объединяем элементы строки с пробелами
		_, err := fmt.Fprintln(file, rowStr)
		if err != nil {
			return err
		}
	}

	return nil
}

// Сортировка столбца матрицы
func sortMatrixColumn(matrix [][]string, numColumn int) {
	if len(matrix) > 0 && numColumn >= len(matrix[0]) {
		log.Fatalln("The column is out of bounds")
	}
	sort.SliceStable(matrix, func(i, j int) bool {
		firstNum, err := strconv.Atoi(matrix[i][numColumn])
		if err != nil {
			log.Fatalf("Cannot convert %s in column %d to number", matrix[i][numColumn], numColumn)
		}
		secondNum, err := strconv.Atoi(matrix[j][numColumn])
		if err != nil {
			log.Fatalf("Cannot convert %s in column %d to number", matrix[j][numColumn], numColumn)
		}
		if numeric {
			return firstNum < secondNum
		}
		return firstNum > secondNum
	})
}

// SortFile - сортирует файл с учетом указанных флагов
func SortFile(inputFile, outputFile string) error {
	fileData, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	lines := getLinesFromBytes(fileData, unique)
	matrix := createMatrixFromLines(lines)

	sortMatrixColumn(matrix, column)

	err = writeMatrixToFile(matrix, outputFile)
	if err != nil {
		return err
	}

	return nil
}
func main() {

	err := SortFile(inputFile, "output.txt")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}
