Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
Вывод прогарммы:
error
Происходит это потому, что из функции test мы хоть и возвращаем nil указатель на структуру customError, но мы потом преобразуем этот тип в interface error, теперь в этом интерфейсе itab указывает на метаданные типа cutomError, а data - содержит nil указатель, однако весь интерфейс != nil.
```
