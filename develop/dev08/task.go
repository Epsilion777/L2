package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/shirou/gopsutil/process"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// ShellUtility - UNIX-шелл-утилита с поддержкой ряда простейших команд
type ShellUtility struct {
	commands []string
	args     []string
}

// Run - запускает на выполнение переданную команду
func (su *ShellUtility) Run() error {
	if len(su.args) < 1 {
		return nil
	}
	// Определяем переданную команду
	switch su.args[0] {
	case "cd":
		err := su.cd()
		if err != nil {
			return err
		}
	case "pwd":
		err := su.pwd()
		if err != nil {
			return err
		}
	case "echo":
		err := su.echo()
		if err != nil {
			return err
		}
	case "kill":
		err := su.kill()
		if err != nil {
			return err
		}
	case "ps":
		err := su.ps()
		if err != nil {
			return err
		}
	case "fork":
		err := su.fork()
		if err != nil {
			return err
		}
	case "exec":
		err := su.exec()
		if err != nil {
			return err
		}
	case "quit":
		su.quit()

	default:
		fmt.Println("Введите команду из списка доступных!")
	}
	return nil
}

// Функция для смены текущей директории
func (su *ShellUtility) cd() error {
	if len(su.args) < 2 {
		return errors.New("укажите путь")
	}
	// Получение текущей дериктории
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}
	// Смена директории
	err = os.Chdir(su.args[1])
	if err != nil {
		return err
	}
	// Получение новой текущей директории
	newDir, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Printf("%s -> %s\n", currentDir, newDir)
	return nil
}

// Функция для вывода пути к текущему каталогу
func (su *ShellUtility) pwd() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Println(dir)
	return nil
}

// Функция для вывода переданных аргументов
func (su *ShellUtility) echo() error {
	if len(su.args) < 2 {
		return errors.New("аргументы отсутствуют")
	}
	args := strings.Join(su.args[1:], " ")
	fmt.Println(args)
	return nil
}

// Функция для "убийства" процесса по его PID
func (su *ShellUtility) kill() error {
	if len(su.args) < 2 {
		return errors.New("укажите PID процесса")
	}

	numProcess, err := strconv.Atoi(su.args[1])
	if err != nil {
		return errors.New("введите корректный PID")
	}

	p, err := process.NewProcess(int32(numProcess))
	if err != nil {
		return err
	}
	// Прекращение процесса
	err = p.Terminate()
	if err != nil {
		return err
	}
	return nil
}

// Функция для вывода всех запущенных процессов
func (su *ShellUtility) ps() error {
	processes, err := process.Processes()
	if err != nil {
		return err
	}

	// Сортировка процессов по идентификатору PID
	sort.Slice(processes, func(i, j int) bool {
		return processes[i].Pid < processes[j].Pid
	})

	fmt.Printf("%-7s %-40s %s\n", "PID", "Name", "Status")
	fmt.Println("------------------------------------------------------------")

	for _, p := range processes {
		name, _ := p.Name()
		status, _ := p.Status()

		fmt.Printf("%-7d %-40s %s\n", p.Pid, name, status)
	}
	return nil
}

// Функция для запуска дочернего процесса выполнения
func (su *ShellUtility) fork() error {
	if len(su.args) < 2 {
		return errors.New("аргументы для fork отсутствуют")
	}
	wg := sync.WaitGroup{}
	m := sync.Mutex{}

	su.args = su.args[1:] // Оставляем команды после fork
	wg.Add(1)
	go func() {
		fmt.Println("Запущен дочерний процесс")
		// Блокируем разделяемый ресурс
		m.Lock()
		su.Run()
		m.Unlock()
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("Дочерний процесс завершен")
	return nil
}

// Функция для выполнения комманд из файла
func (su *ShellUtility) exec() error {
	if len(su.args) < 2 {
		return errors.New("введите название файла")
	}

	fileData, err := os.ReadFile(su.args[1])
	if err != nil {
		return err
	}
	// Разбиваем файл на строки команд
	lines := strings.Split(string(fileData), "\r\n")
	for _, line := range lines {
		su.args = strings.Fields(line)
		su.Run()
	}
	return nil
}

// Функция для завершения диалогового окна ShellUtility
func (su *ShellUtility) quit() {
	fmt.Println("Завершение работы")
	os.Exit(0)
}

// CMD - экземпляр утилиты
var CMD ShellUtility

// Наполняем ShellUtility доступными командыми
func init() {
	CMD.commands = []string{
		"- cd <args> - смена директории",
		"- pwd - показать путь до текущего каталога",
		"- echo <args> - вывод аргумента в STDOUT",
		`- kill <args> - "убить" процесс, переданный в качесте аргумента`,
		"- ps - выводит общую информацию по запущенным процессам",
		"- fork <args> - создает дочерний процесс выполнения",
		"- exec <args> - выполняет команды из файла",
		"- quit - выход из утилиты",
	}
}

func main() {
	fmt.Println("Добро пожаловать в шелл-утилиту")

	for _, v := range CMD.commands {
		fmt.Println(v)
	}
	for {
		fmt.Println("Введите команду: ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			input := scanner.Text()
			// Заполняем утилиту переданной командой с аргументами
			CMD.args = strings.Fields(input)
			err := CMD.Run()
			if err != nil {
				fmt.Println("Error:", err)

			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error:", err)
		}
	}
}
