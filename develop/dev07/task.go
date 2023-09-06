package main

import (
	"fmt"
	"sync"
	"time"
)

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/

// Or - функция, которая объеденяет N-ое кол-во single-каналов в один
func Or(channels ...<-chan interface{}) <-chan interface{} {
	once := sync.Once{}
	resultChan := make(chan interface{})
	for _, v := range channels {
		go func(channel <-chan interface{}) {
			// Когда любой из singe-каналов закроется, цикл завершится и resultChan также закроется
			for v := range channel {
				resultChan <- v
			}
			// Выполняем закрытие канала единожды, чтобы избежать ситуации закрытия канала дважды
			once.Do(func() {
				close(resultChan)
			})
		}(v)
	}
	fmt.Println("Start working")
	return resultChan
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	start := time.Now()
	<-Or(
		sig(2*time.Minute),
		sig(5*time.Second),
		sig(8*time.Second),
		sig(1*time.Minute),
		sig(2*time.Hour),
		sig(5*time.Hour),
	)
	fmt.Printf("done after %v", time.Since(start))
	time.Sleep(10 * time.Second)
}
