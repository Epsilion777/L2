package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port
go-telnet mysite.ru 8080
go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

var (
	timeout int
)

func init() {
	flag.IntVar(&timeout, "timeout", 10, "Таймаут на подключение к серверу")
	flag.Parse()
}

func main() {
	var conn net.Conn
	host := flag.Args()[0]
	port := flag.Args()[1]
	start := time.Now() // Запоминаем текущее время для отсчета таймаута

	// Цикл для попыток установки соединения с сервером
	for {
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
		if err == nil {
			fmt.Printf("Successful connection to %s:%s\n", host, port)
			defer conn.Close()
			break
		}

		if time.Since(start) > time.Duration(timeout)*time.Second { // Проверка на истечение таймаута
			log.Fatalf("Connection timeout to %s:%s has expired", host, port) // Завершение программы с сообщением о истечении таймаута
		}
	}

	scannerStdin := bufio.NewScanner(os.Stdin) // Инициализация сканнера для стандартного ввода
	scannerServer := bufio.NewScanner(conn)    // Инициализация сканнера для сетевого соединения

	fmt.Println("Write messages to socket")

	// Цикл для считывания сообщений с клавиатуры и их отправки на сервер
	for scannerStdin.Scan() {
		input := scannerStdin.Text()
		_, err := fmt.Fprintf(conn, "%s\n", input) // Отправка строки на сервер

		if err != nil {
			fmt.Println("Server closed connection")
			os.Exit(0)
		}

		if scannerServer.Scan() { // Чтение ответа сервера
			input := scannerServer.Text()
			fmt.Printf("%s Server response: %s\n", time.Now().Format("15:04:05"), (string(input)))
		}
	}
}
