Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
Сначала программа выведет в рандомном порядке значения 1-8, а после начнет выводить 0.
Нули программа начнет выводить, потому что в функции asChan после передачи в канал всех значений канал закрывается, однако в конструкции select из закрытого канала читается значение по умолчанию, т.е. 0. И так как в функции merge канал c не закрывается, то 0 будет выводиться в бесконечном цикле.
```
