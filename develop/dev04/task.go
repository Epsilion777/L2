package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// areAnagrams - проверяет, что две строки являются анаграммами
func areAnagrams(first, second string) bool {
	if len(first) != len(second) {
		return false
	}
	word1 := []byte(first)
	word2 := []byte(second)

	sort.Slice(word1, func(i, j int) bool {
		return word1[i] < word1[j]
	})
	sort.Slice(word2, func(i, j int) bool {
		return word2[i] < word2[j]
	})

	return string(word1) == string(word2)
}

// FindAnagrams - создает множество анаграм. В качестве ключа - первое встретившееся в словаре слово из множества.
func FindAnagrams(slc *[]string) *map[string]*[]string {
	// Само множество анаграмм
	anagrams := make(map[string]*[]string)
	// Слова, которые уже встречались
	uniqueElems := make(map[string]struct{})

	for _, word := range *slc {
		word = strings.ToLower(word)
		// Флаг указывает, была ли найдена анаграмма среди ключей, если нет - необходимо создать новый ключ.
		flag := false

		// Данная проверка нужна для инициализации первого ключа anagrams
		if len(anagrams) == 0 {
			var slc []string
			anagrams[word] = &slc
			uniqueElems[word] = struct{}{}
			continue
		}

		for key := range anagrams {
			if areAnagrams(key, word) {
				flag = true
				if _, ok := uniqueElems[word]; !ok {
					uniqueElems[word] = struct{}{}
					*anagrams[key] = append(*anagrams[key], word)
				}
				break
			}
		}
		if !flag {
			var slc []string
			anagrams[word] = &slc
		}
	}

	// Убираем множества, состоящие из 1-ого элемента и сортируем каждый массив
	for key, val := range anagrams {
		if len(*val) == 0 {
			delete(anagrams, key)
		}

		sort.Slice(*val, func(i, j int) bool {
			return (*val)[i] < (*val)[j]
		})
	}

	return &anagrams
}

func main() {
	words := []string{
		"пятак", "пятка", "тяпка", "тяпка",
		"листок", "слиток", "столик",
		"кот", "ток",
		"сон", "нос",
		"слон",
		"рот",
	}

	anagramSets := FindAnagrams(&words)

	for key, value := range *anagramSets {
		fmt.Printf("Set of anagrams for \"%s\": %v\n", key, value)
	}

}
