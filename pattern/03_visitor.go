package pattern

import "fmt"

type Transport interface {
	Drive()
	Stop()
	Accept(*Visitor)
}

// Ниже описаны структуры, к которым хотелось бы добавить общий метод без вмешательства в код самих структур

// Структура Car
type AutoCar struct {
	MaxSpeed float64
	Model    string
	Wheels   int
	Mileage  int
}

func (c *AutoCar) Drive() {
	fmt.Println("The car drove")
}

func (c *AutoCar) Stop() {
	fmt.Println("The car stopped")
}

// Метод для удобного взаимодействия с "посетителем"
func (c *AutoCar) Accept(v Visitor) {
	v.visitForCar(c)
}

// Структура Airplane
type Airplane struct {
	MaxSpeed float64
	Model    string
	Mileage  int
}

func (a *Airplane) Drive() {
	fmt.Println("The plane flew")
}

func (a *Airplane) Stop() {
	fmt.Println("The plane landed")
}

// Метод для удобного взаимодействия с "посетителем"
func (a *Airplane) Accept(v Visitor) {
	v.visitForAirplane(a)
}

// Структура Ship
type Ship struct {
	MaxSpeed float64
	Model    string
	Mileage  int
}

func (s *Ship) Drive() {
	fmt.Println("The ship sailed")
}

func (s *Ship) Stop() {
	fmt.Println("The ship stopped")
}

// Метод для удобного взаимодействия с "посетителем"
func (s *Ship) Accept(v Visitor) {
	v.visitForShip(s)
}

// Описание интерфейса "посетителя", который реализует дополнительный метод для каждой структуры
type Visitor interface {
	visitForCar(*AutoCar)
	visitForAirplane(*Airplane)
	visitForShip(*Ship)
}

// Конкретный "посетитель", который будет изменять максимальную скорость транспорта
type Accelerator struct {
	Boost float64
}

// Метод ускоряющий машину
func (a *Accelerator) visitForCar(c *AutoCar) {
	oldMaxSpeed := c.MaxSpeed
	newMaxSpeed := c.MaxSpeed * a.Boost
	c.MaxSpeed = newMaxSpeed
	fmt.Printf("The car has been accelerated, the old max speed is %0.1f, the new max speed is %0.1f\n", oldMaxSpeed, newMaxSpeed)
}

// Метод ускоряющий самолет
func (a *Accelerator) visitForAirplane(air *Airplane) {
	oldMaxSpeed := air.MaxSpeed
	newMaxSpeed := air.MaxSpeed * a.Boost
	air.MaxSpeed = newMaxSpeed
	fmt.Printf("The airplane has been accelerated, the old max speed is %0.1f, the new max speed is %0.1f\n", oldMaxSpeed, newMaxSpeed)
}

// Метод ускоряющий корабль
func (a *Accelerator) visitForShip(s *Ship) {
	oldMaxSpeed := s.MaxSpeed
	newMaxSpeed := s.MaxSpeed * a.Boost
	s.MaxSpeed = newMaxSpeed
	fmt.Printf("The ship has been accelerated, the old max speed is %0.1f, the new max speed is %0.1f\n", oldMaxSpeed, newMaxSpeed)
}

// Конкретный "посетитель", который будет изменять пробег транспорта
type MileageChanger struct {
	Mileage int
}

// Метод изменяющий пробег машины
func (m *MileageChanger) visitForCar(c *AutoCar) {
	oldMileage := c.Mileage
	newMileage := c.Mileage + m.Mileage
	c.Mileage = newMileage
	fmt.Printf("The mileage of the car has been changed, the old mileage is %d, the new %d\n", oldMileage, newMileage)
}

// Метод изменяющий пробег самолета
func (m *MileageChanger) visitForAirplane(air *Airplane) {
	oldMileage := air.Mileage
	newMileage := air.Mileage + m.Mileage
	air.Mileage = newMileage
	fmt.Printf("The mileage of the airplane has been changed, the old mileage is %d, the new %d\n", oldMileage, newMileage)
}

// Метод изменяющий пробег корабля
func (m *MileageChanger) visitForShip(s *Ship) {
	oldMileage := s.Mileage
	newMileage := s.Mileage + m.Mileage
	s.Mileage = newMileage
	fmt.Printf("The mileage of the ship has been changed, the old mileage is %d, the new %d\n", oldMileage, newMileage)
}

func VisitorFunc() {
	// Objects
	car := AutoCar{334, "Audi R8", 4, 6000}
	airplane := Airplane{950, "Boeing 767", 3500}
	ship := Ship{42, "Titanik", 15000}

	fmt.Printf("Start characters:\n%+v,\n%+v,\n%+v\n", car, airplane, ship)

	// Boost speed
	visitorAccelerator := Accelerator{1.2}
	car.Accept(&visitorAccelerator)
	airplane.Accept(&visitorAccelerator)
	ship.Accept(&visitorAccelerator)

	fmt.Printf("After boost characters:\n%+v,\n%+v,\n%+v\n", car, airplane, ship)

	// Change mileage (-1000 )
	visitorMileage := MileageChanger{-1000}
	car.Accept(&visitorMileage)
	airplane.Accept(&visitorMileage)
	ship.Accept(&visitorMileage)

	fmt.Printf("After change mileage characters:\n%+v,\n%+v,\n%+v\n", car, airplane, ship)
}
