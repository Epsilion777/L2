package main

import (
	"flag"
	"log"
	"os"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	name         string
	filePath     string
	flags        string
	expectedStrs []string
}

func TestFindAnagrams(t *testing.T) {
	tests := []testStruct{
		{"№1", "develop/dev05/input1.txt", "-i=true -C=1 -n=true -c=true dk", []string{"0: kasdkas asdka ksd aksdksa dk",
			"1: asdkasdk asdasdk asdkasd kasdkasas",
			"2: asda adsad asdasd kdas",
			"3: asda dsad adsad asdmvvm",
			"4: vmsdv dk vd;lazsc;",
			"5: fasdkfkas asdkasdk sad",
			"Количество строк: 6"}},
		{"№2", "develop/dev05/input2.txt", "-i=true -C=2 -n=true apple", []string{"0: lasdlaslldsal asldlasdlldas",
			"1: asldlasd asdlasld lasdlasld",
			"2: apple asdlalsdlasd asdlals",
			"3: rlelwerlapplelasdlasl aslelel qwelqwel qwlel wqll",
			"4: фывдфывддфы дфывдфыд дфывддвы",
			"5: asldlasdl asdasd asdllasd"}},
		{"№3", "develop/dev05/input3.txt", "-i=true -n=true -F=true \n", []string{"2: asda adsad a\nsdasd kdas",
			"4: vmsdv dk vd;lazsc;\n"}},
		{"№4", "evelop/dev05/input4.txt", "-i=true -C=1 -n=true -v=true asd", []string{
			"1: fflflflfl fl fllf lf",
			"3: flflfl flfl",
		}},
		{"№5", "evelop/dev05/input5.txt", "-i=true -n=true -i=true РоССиЯ", []string{
			"30: РоссияРуандаРумыния",
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pattern := flag.Args()[0]
			if fixed {
				pattern = `\Q` + pattern + `\E`
			}

			if ignoreСase {
				pattern = `(?i)` + pattern
			}

			// Генерируем регулярное выражение
			regExp := regexp.MustCompile(pattern)
			fileData, err := os.ReadFile(tt.filePath)
			if err != nil {
				log.Fatalf("Error: %s", err)
			}
			resultStrs := GrepFunc(fileData, regExp)

			for i, v := range resultStrs {
				assert.Equal(t, tt.expectedStrs[i], v)
			}
		})
	}
}
