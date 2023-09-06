package pattern

import (
	"fmt"
	"log"
)

// Интерфейс продукта
type CarI interface {
	setModel(string)
	setSpeed(int)
	getModel() string
	getSpeed() int
}

// Конкретный продукт Car
type Car struct {
	model string
	speed int
}

func (c *Car) setModel(model string) {
	c.model = model
}

func (c *Car) setSpeed(speed int) {
	c.speed = speed
}

func (c *Car) getSpeed() int {
	return c.speed
}

func (c *Car) getModel() string {
	return c.model
}

// Конкретный продукт Audi
type Audi struct {
	Car
}

func newAudi() CarI {
	return &Audi{
		Car: Car{
			model: "Audi R8",
			speed: 334,
		},
	}
}

// Конкретный продукт Mercedes
type Mercedes struct {
	Car
}

func newMercedes() CarI {
	return &Mercedes{
		Car{
			model: "Mercedes Benz AMG GT",
			speed: 325,
		},
	}
}

// Конкретный продукт Bmw
type Bmw struct {
	Car
}

func newBmw() CarI {
	return &Bmw{
		Car{
			model: "BMW M4",
			speed: 307,
		},
	}
}

// Фабричный метод
func createCar(brand string) (CarI, error) {
	switch brand {
	case "Audi":
		return newAudi(), nil
	case "Mercedes":
		return newMercedes(), nil
	case "Bmw":
		return newBmw(), nil
	default:
		return nil, fmt.Errorf("wrong car brand: %s", brand)
	}
}
func FactoryMethodFunc() {

	// Создаем продукт с помощью фабричного метода
	audi, err := createCar("Audi")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(audi.getModel(), audi.getSpeed())
	}

	// Создаем продукт с помощью фабричного метода
	mercedes, err := createCar("Mercedes")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(mercedes.getModel(), mercedes.getSpeed())
	}

	// Создаем продукт с помощью фабричного метода
	bmw, err := createCar("Bmw")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(bmw.getModel(), bmw.getSpeed())
	}

	// Создаем продукт с помощью фабричного метода
	volkswagen, err := createCar("Volkswagen")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(volkswagen.getModel(), volkswagen.getSpeed())
	}

}
