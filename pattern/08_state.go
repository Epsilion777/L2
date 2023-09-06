package pattern

import "fmt"

// Интерфейс сосояния (агрегатное состояние воды)
type WaterState interface {
	Heat(*Water)
	Frost(*Water)
}

// Контекст (вода)
type Water struct {
	state WaterState
}

// Конструктор для создания структуры "Water"
func NewWater(w WaterState) *Water {
	return &Water{
		state: w,
	}
}

// Метод, который повышает температуру
func (w *Water) Heat(water *Water) {
	w.state.Heat(water)
}

// Метод, который понижает температуру
func (w *Water) Frost(water *Water) {
	w.state.Frost(water)
}

// Лед
type SolidWaterState struct {
}

// Метод, который повышает температуру
func (w *SolidWaterState) Heat(water *Water) {
	fmt.Println("The ice turned into water")
	water.state = &LiquidWaterState{}
}

// Метод, который понижает температуру
func (w *SolidWaterState) Frost(water *Water) {
	fmt.Println("Cooling the ice")
}

// Жидкая вода
type LiquidWaterState struct {
}

// Метод, который повышает температуру
func (w *LiquidWaterState) Heat(water *Water) {
	fmt.Println("The water turned into steam")
	water.state = &SteamWaterState{}
}

// Метод, который понижает температуру
func (w *LiquidWaterState) Frost(water *Water) {
	fmt.Println("Water turns into ice")
	water.state = &SolidWaterState{}
}

// Водяной пар
type SteamWaterState struct {
}

// Метод, который повышает температуру
func (w *SteamWaterState) Heat(water *Water) {
	fmt.Println("The gaseous water is heated")
}

// Метод, который понижает температуру
func (w *SteamWaterState) Frost(water *Water) {
	fmt.Println("Steam turns into water")
	water.state = &LiquidWaterState{}
}

func StateFunc() {
	solidWaterState := SolidWaterState{}

	// Состояние лед
	water := NewWater(&solidWaterState)
	// Состояние жидкая вода
	water.Heat(water)
	// Состояние водяной пар
	water.Heat(water)
	// Состояние жидкая вода
	water.Frost(water)
	// Состояние лед
	water.Frost(water)
}
