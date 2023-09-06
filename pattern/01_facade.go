package pattern

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// Файловая система
type FileSystem struct {
	existingFiles []string
}

// Метод для проверки существования файла в файловой системе
func (fs *FileSystem) FileExist(filename string) error {
	for _, v := range fs.existingFiles {
		if v == filename {
			return nil
		}
	}
	return fmt.Errorf("the file %s does not exist", filename)
}

// Текстовый редактор
type TextEditor struct {
	*FileSystem
}

// Метод для создания файла
func (te *TextEditor) CreateFile(filename string) error {
	if err := te.FileSystem.FileExist(filename); err == nil {
		return fmt.Errorf("file with that name already exists")
	}
	te.existingFiles = append(te.existingFiles, filename)
	fmt.Printf("%s created\n", filename)
	return nil
}

// Метод для записи кода в файл
func (te *TextEditor) WriteCode(filename, code string) error {
	if err := te.FileSystem.FileExist(filename); err != nil {
		return err
	}
	fmt.Println("The code is written to a file")
	return nil
}

// Метод для сохранения файла
func (te *TextEditor) SaveCode(filename string) error {
	if err := te.FileSystem.FileExist(filename); err != nil {
		return err
	}
	fmt.Printf("File %s saved\n", filename)
	return nil
}

// Компилятор
type Compiller struct {
	*FileSystem
}

// Метод для компилирования файла
func (c *Compiller) Compile(filename string) (string, error) {
	if err := c.FileSystem.FileExist(filename); err != nil {
		return "", err
	}

	index := strings.LastIndex(filename, ".") // Находим индекс последней точки
	if index >= 0 {
		filename = filename[:index] // Обрезаем строку до последней точки
	}
	executingFilename := filename + ".exe"
	c.FileSystem.existingFiles = append(c.FileSystem.existingFiles, executingFilename)
	fmt.Printf("The code is compiled to %s\n", executingFilename)
	return executingFilename, nil
}

// Исполнитель, который запускает exe файл
type Executor struct {
	*FileSystem
}

// Метод для выполения exe программы
func (clr *Executor) Execute(filename string) error {
	if err := clr.FileSystem.FileExist(filename); err != nil {
		return err
	}

	fmt.Printf("Program %s is running\n", filename)
	return nil
}

// Метод для завершения exe программы
func (c *Executor) Finish(filename string) {
	fmt.Println("The code is completed")

}

// Среда разработки (Фасад)
type IDE struct {
	TextEditor
	Compiller
	Executor
}

// Метод для запуска программы
func (ide *IDE) StartProgram(fc *FileSystem, filename, code string) error {
	te := TextEditor{fc}
	c := Compiller{fc}
	executor := Executor{fc}
	if err := te.CreateFile(filename); err != nil {
		return err
	}
	if err := te.WriteCode(filename, code); err != nil {
		return err
	}
	if err := te.SaveCode(filename); err != nil {
		return err
	}
	compiledFilename, err := c.Compile(filename)
	if err != nil {
		return err
	}
	if err := executor.Execute(compiledFilename); err != nil {
		return err
	}
	fmt.Println("The program works")
	return nil
}

// Метод для завершения программы
func (ide *IDE) FinishProgramm(fc *FileSystem, filename string) error {
	if err := fc.FileExist(filename); err != nil {
		return err
	}
	ide.Finish(filename)
	fmt.Printf("The program %s has finished working", filename)
	return nil
}

func FacadeFunc() {
	fc := &FileSystem{}
	ide := IDE{}
	// Запускаем программу
	err := ide.StartProgram(fc, "hello.go", `fmt.Println("Hello world")`)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("------------------------------------------")
	time.Sleep(3 * time.Second)
	// Завершаем программу
	err = ide.FinishProgramm(fc, "hello.exe")
	if err != nil {
		log.Println(err)
	}

}
