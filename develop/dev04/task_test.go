package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	name        string
	input       []string
	expectedStr string
}

func TestFindAnagrams(t *testing.T) {
	tests := []testStruct{
		{"№1", []string{
			"пятак", "пятка", "тяпка", "тяпка",
			"листок", "слиток", "столик",
			"кот", "ток",
			"сон", "нос",
			"слон",
			"рот",
		}, "пятак: &[пятка тяпка] листок: &[слиток столик] кот: &[ток] сон: &[нос] "},
		{"№2", []string{
			"австралопитек", "ватерполистка", "верблюд",
			"молоко", "комоло", "коломо",
			"еж", "аист",
		}, "австралопитек: &[ватерполистка] молоко: &[коломо комоло] "},
		{"№3", []string{
			"пасечник", "песчаник", "песчинка",
			"марочник", "рамочник", "романчик",
			"акробат, работка",
			"свинец",
		}, "пасечник: &[песчаник песчинка] марочник: &[рамочник романчик] "},
		{"№4", []string{}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := strings.Builder{}
			anagrams := FindAnagrams(&tt.input)
			for k, v := range *anagrams {
				output.WriteString(fmt.Sprintf("%s: %v ", k, v))
			}
			assert.Equal(t, tt.expectedStr, output.String())
		})
	}
}
