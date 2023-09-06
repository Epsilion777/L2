package pattern

import "fmt"

// Контекст (Сумматор)
type Summator struct {
	sum      int
	strategy Strategy
}

// Метод для выбора стратегии
func (s *Summator) SetStrategy(strategy Strategy) {
	s.strategy = strategy
}

// Метод для просчета суммы последовательности в зависимости от стратегии
func (s *Summator) Sum(slc []int) {
	s.sum = s.strategy.Sum(slc)
}

// Интерфейс стратегии
type Strategy interface {
	Sum([]int) int
}

// Конкретная стратегия
type SimpleSum struct {
}

// Cумма всех элементов последовательности O(n)
func (s *SimpleSum) Sum(slc []int) int {
	sum := 0
	for _, v := range slc {
		sum += v
	}
	return sum
}

// Конкретная стратегия
type ArithmeticProgressionSum struct {
}

// Сумма всех элементов арифметической прогрессии O(1)
func (s *ArithmeticProgressionSum) Sum(slc []int) int {
	if len(slc) == 0 {
		return 0
	}
	if len(slc) == 1 {
		return slc[0]
	}
	return (slc[0] + slc[len(slc)-1]) * len(slc) / 2
}

func StrategyFunc() {
	// Пример для работы с обычной стратегиией (сложить все элементы)
	simpleSlice := []int{1, 2, 10, -5, 18, 19, 25, 44, 11}
	// Пример для работы со стратегией арифметическая прогрессия ( s = (a0+an)*n/2 )
	arithmeticProgressionSlc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	simpleSumStrategy := SimpleSum{}
	summator := Summator{}
	// Выбор обычной стратегии
	summator.SetStrategy(&simpleSumStrategy)
	// Расчет суммы
	summator.Sum(simpleSlice)
	fmt.Println(summator.sum)

	// Выбор стратегии арифметическая последовательность
	arithmeticProgressionStrategy := ArithmeticProgressionSum{}
	summator.SetStrategy(&arithmeticProgressionStrategy)
	// Расчет суммы
	summator.Sum(arithmeticProgressionSlc)
	fmt.Println(summator.sum)
}
