package pattern

import "fmt"

// Описание интерфейса команд, которые будут исполняться
type Command interface {
	Execute()
}

// Конкретная команда, которая отвечает за включение устройства
type OnCommand struct {
	device Device
}

// Метод, который запускает выполнение (включает устройство)
func (c *OnCommand) Execute() {
	c.device.On()
}

// Конкретная команда, которая отвечает за выключение устройства
type OffCommand struct {
	device Device
}

// Метод, который запускает выполнение (выключает устройство)
func (c *OffCommand) Execute() {
	c.device.Off()
}

// Интерфейса устройства
type Device interface {
	On()
	Off()
}

// Девайс (телевизор), который будет выполнять команды (получатель)
type TV struct {
	isRunning bool
}

// Включить телевизор
func (t *TV) On() {
	t.isRunning = true
	fmt.Println("The TV turned on")
}

// Выключить телевизор
func (t *TV) Off() {
	t.isRunning = false
	fmt.Println("The TV turned off")
}

// Исполнитель команд
type Invoker struct {
	commands []Command
}

// Положить команду в очередь исполнителя
func (b *Invoker) StoreCommand(c *Command) {
	b.commands = append(b.commands, *c)
}

// Извлечь команду из очереди исполнителя
func (b *Invoker) UnStoreCommand() {
	if len(b.commands) != 0 {
		b.commands = b.commands[:len(b.commands)-1]
	}
}

// Выполнить все задачи из очереди
func (b *Invoker) Execute() {
	for _, v := range b.commands {
		v.Execute()
	}
}

func CommandFunc() {
	tv := &TV{}
	// Создание конкретных задач (включить и выключить телевизор)
	onCommand := OnCommand{tv}
	offCommand := OffCommand{tv}

	// Создание исполнителя
	invoker := Invoker{}
	// Добавление команд в очередь задач
	invoker.commands = append(invoker.commands, &onCommand, &offCommand)
	// Выполняем все задачи из очереди
	invoker.Execute()
	// Очищаем все задачи из очереди
	invoker.UnStoreCommand()
	invoker.UnStoreCommand()
	// Пробуем вызвать снова выполнение задач
	invoker.Execute() // ничег не произойдет, т.к. задач не осталось
}
