package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	name    string
	command string
	wandErr bool
}

func TestSortFile(t *testing.T) {
	tests := []testStruct{
		{"№1", "cd testdirectory", false},
		{"№2", "pwd", false},
		{"№3", "echo 123 GOROOT", false},
		{"№4", "ps", false},
		{"№5", "fork pwd", false},
		{"№6", "testdirectory/commands.txt", false},
		{"№7", "exec commandsNotFound.txt", true},
		{"№8", "cd testdirectoryNotFound", true},
		{"№9", "fork", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Заполняем утилиту переданной командой с аргументами
			CMD.args = strings.Fields(tt.command)
			err := CMD.Run()
			if tt.wandErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
