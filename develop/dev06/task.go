package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

# Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	fields    string
	delimiter string
	separated bool
)

func init() {
	flag.StringVar(&fields, "f", "", "Выбрать поля (колонки)")
	flag.StringVar(&delimiter, "d", "\t", "Использовать другой разделитель")
	flag.BoolVar(&separated, "s", false, "Только строки с разделителем")

	flag.Parse()
}

// Создание матрицы слов из строк
func createMatrixFromLines(lines []string) [][]string {
	var matrix [][]string
	countOfColumn := 0

	for i, v := range lines {
		if separated {
			if strings.Contains(v, delimiter) {
				matrix = append(matrix, strings.Split(v, delimiter))
				if len(matrix[i]) > countOfColumn {
					countOfColumn = len(matrix[i])
				}
			}
		} else {
			matrix = append(matrix, strings.Split(v, delimiter))
			if len(matrix[i]) > countOfColumn {
				countOfColumn = len(matrix[i])
			}
		}
	}
	for i := range matrix {
		for len(matrix[i]) < countOfColumn {
			matrix[i] = append(matrix[i], "")
		}
	}

	return matrix
}

// GetColumn - функция для вывода нужной колонки матрицы
func GetColumn(matrix [][]string) ([]string, error) {
	elems := make([]string, 0, len(matrix))
	column, err := strconv.Atoi(fields)
	if err != nil {
		return []string{}, err
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if j == column {
				elems = append(elems, matrix[i][j])
			}
		}
	}
	return elems, nil
}

func main() {
	// Читаем файл
	fileData, err := os.ReadFile("develop/dev06/input2.txt")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	// Разделяем прочитанный файл на строки
	lines := strings.Split(string(fileData), "\r\n")

	// Создаем матрицу из слов
	matrix := createMatrixFromLines(lines)
	// Получаем нужный столбец матрицы
	column, err := GetColumn(matrix)

	if err != nil {
		log.Fatalf("Invalid flag -f=%s", fields)
	}
	// Выводим столбец матрицы
	for _, v := range column {
		fmt.Println(v)
	}
}
