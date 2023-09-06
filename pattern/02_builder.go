package pattern

import (
	"fmt"
)

// Phone - структура, которую создаем
type Phone struct {
	Brand      string
	CPU        int
	RAM        int
	Megapixels int
}

// Интерфейс билдера для телефона
type PhoneBuilderI interface {
	setBrand()
	setCPU()
	setRAM()
	setMegapixels()
	getPhone() Phone
}

// Структура билдера для iPhone
type IPhoneBuilder struct {
	Brand      string
	CPU        int
	RAM        int
	Megapixels int
}

// Конструктор для билдера iPhone
func newIPhoneBuilder() *IPhoneBuilder {
	return &IPhoneBuilder{}
}

// Метод для установки бренда у iPhone
func (ib *IPhoneBuilder) setBrand() {
	ib.Brand = "Apple"
}

// Метод для установки CPU у iPhone
func (ib *IPhoneBuilder) setCPU() {
	ib.CPU = 6
}

// Метод для установки RAM у iPhone
func (ib *IPhoneBuilder) setRAM() {
	ib.RAM = 512
}

// Метод для установки мегапикселей у iPhone
func (ib *IPhoneBuilder) setMegapixels() {
	ib.Megapixels = 12
}

// Метод для создания телефона модели iPhone
func (ib *IPhoneBuilder) getPhone() Phone {
	return Phone{
		Brand:      ib.Brand,
		CPU:        ib.CPU,
		RAM:        ib.RAM,
		Megapixels: ib.Megapixels,
	}
}

// Структура билдера для Samsung
type SamsungPhoneBuilder struct {
	Brand      string
	CPU        int
	RAM        int
	Megapixels int
}

// Конструктор для билдера Samsung
func newSamsungPhoneBuilder() *SamsungPhoneBuilder {
	return &SamsungPhoneBuilder{}
}

// Метод для установки бренда у Samsung
func (sb *SamsungPhoneBuilder) setBrand() {
	sb.Brand = "Samsung"
}

// Метод для установки CPU у Samsung
func (sb *SamsungPhoneBuilder) setCPU() {
	sb.CPU = 8
}

// Метод для установки RAM у Samsung
func (sb *SamsungPhoneBuilder) setRAM() {
	sb.RAM = 256
}

func (sb *SamsungPhoneBuilder) setMegapixels() {
	sb.Megapixels = 50
}

// Метод для создания телефона модели Samsung
func (sb *SamsungPhoneBuilder) getPhone() Phone {
	return Phone{
		Brand:      sb.Brand,
		CPU:        sb.CPU,
		RAM:        sb.RAM,
		Megapixels: sb.Megapixels,
	}
}

// Директор, который будет выполнять пошаговую сборку телефона в зависимости от переданного билдера
type Director struct {
	builder PhoneBuilderI
}

// Конструктор для создания диерктора
func newDirector(b PhoneBuilderI) *Director {
	return &Director{
		builder: b,
	}
}

// Метод для смены билдера у директора
func (d *Director) setBuilder(b PhoneBuilderI) {
	d.builder = b
}

// Метод для создания телефона
func (d *Director) buildPhone() Phone {
	d.builder.setBrand()
	d.builder.setCPU()
	d.builder.setRAM()
	d.builder.setMegapixels()
	return d.builder.getPhone()
}

func BuilderFunc() {
	// Создаем билдер iPhone
	iphoneBuilder := newIPhoneBuilder()
	// Создаем билдер Samsung
	samsungBuilder := newSamsungPhoneBuilder()
	// Создаем директора для создания телефонов (билдер - iPhone)
	director := newDirector(iphoneBuilder)
	iPhone := director.buildPhone()
	fmt.Printf("%+v\n", iPhone)

	fmt.Println("-------------------------------------------")

	// Устанавливаем у директора нового билдера - Samsung
	director.setBuilder(samsungBuilder)
	samsungPhone := director.buildPhone()
	fmt.Printf("%+v\n", samsungPhone)
}
