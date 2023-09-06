package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	// Устанавливаем прослушивание
	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal("Error in listening port")
	}
	fmt.Println("Waiting for the reciver to connect")

	// Получаем соединение
	conn, _ := ln.Accept()
	defer conn.Close()

	fmt.Println("The server is running and ready to receive messages")
	scanner := bufio.NewScanner(conn)
	// Цикл для считывания сообщений с сервера
	for scanner.Scan() {
		input := scanner.Text()
		fmt.Printf("%s Message received: %s\n", time.Now().Format("15:04:05"), (string(input)))

		_, err := fmt.Fprintf(conn, "%s\n", "OK") // Отправляем сообщение назад, что данные получены успешно
		if err != nil {
			log.Fatalln(err)
		}
	}
	if scanner.Err() != nil {
		fmt.Println("Error reading from socket:", scanner.Err())
	}
}
