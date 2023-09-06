package main

import (
	"log"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	name           string
	inputFileName  string
	flags          string
	expectedResult []string
}

func TestSortFile(t *testing.T) {
	tests := []testStruct{
		{"№1", "develop/dev06/input1.txt", `-f=0 -s=true -d=","`, []string{"Name",
			"Banana",
			"Orange",
			"Lemon",
			"Potato"}},

		{"№2", "develop/dev06/input1.txt", `-f=2 -s=true -d=","`, []string{"Price",
			"100",
			"150",
			"20",
			"5"}},
		{"№3", "develop/dev06/input2.txt", `-f=1 -s=true`, []string{"Count",
			"200",
			"100",
			"1000",
			"5000"}},
		{"№4", "develop/dev06/input2.txt", `-f=0 -s=false`, []string{"Name",
			"Banana",
			"Orange",
			"Lemon",
			"Potato",
			"фыв asd 123"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Читаем файл
			fileData, err := os.ReadFile(tt.inputFileName)
			if err != nil {
				log.Fatalf("Error: %s", err)
			}

			// Разделяем прочитанный файл на строки
			lines := strings.Split(string(fileData), "\r\n")

			// Создаем матрицу из слов
			matrix := createMatrixFromLines(lines)
			column, err := GetColumn(matrix)
			assert.NoError(t, err)
			for i, v := range column {
				assert.Equal(t, tt.expectedResult[i], v)
			}
		})
	}
}
