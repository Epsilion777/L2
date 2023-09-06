package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	name        string
	str         string
	expectedStr string
	wantErr     bool
}

func TestGetTime(t *testing.T) {
	tests := []testStruct{
		{"№1", "a4bc2d5e", "aaaabccddddde", false},
		{"№2", "abcd", "abcd", false},
		{"№3", "45", "", true},
		{"№4", "", "", false},
		{"№5", `qwe\4\5`, "qwe45", false},
		{"№6", `qwe\45 `, "qwe44444", false},
		{"№7", `qwe\\5`, `qwe\\\\\`, false},
		{"№8", "00000", "", true},
		{"№9", `qwe\`, `qwe`, false},
		{"№10", `/qwe\`, `/qwe`, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			unpackingStr, err := Unpacking(tt.str)
			if !tt.wantErr {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedStr, unpackingStr)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
