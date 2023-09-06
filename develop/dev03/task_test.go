package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	name          string
	inputFileName string
	wantErr       bool
}

func TestSortFile(t *testing.T) {
	tests := []testStruct{
		{"№1", "input1.txt", false},
		{"№2", "input2.txt", false},
		{"№3", "input3.txt", false},
		{"№4", "input4.txt", false},
		{"№5", "input5.txt", true},
		{"№6", "input6.txt", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := SortFile(tt.inputFileName, fmt.Sprintf("output%s.txt", tt.name))
			if !tt.wantErr {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
