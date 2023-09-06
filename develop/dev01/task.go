package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

// GetTime функция для получения точного текущего времени с помощью NTP-server
func GetTime(host string) (time.Time, error) {
	return ntp.Time(host)
}

func main() {
	timeFromNtp, err := GetTime("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Connection error to the NTP server: %s\n", err)
		os.Exit(0)
	}
	fmt.Fprintf(os.Stdout, "Current time from ntp: %s\n", timeFromNtp.String())
	fmt.Fprintf(os.Stdout, "Current time from pkg time: %s\n", time.Now())
}
