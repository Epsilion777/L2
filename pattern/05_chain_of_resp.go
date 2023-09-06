package pattern

import "fmt"

// Обработчик
type Worker interface {
	Execute(*Ore)
	SetNext(Worker)
}

// Объект, который передается по цепочке (Руда)
type Ore struct {
	isMined    bool // Добыта ли руда
	isEnriched bool // Обогощена ли руда
	isMelted   bool // Переплавлена ли руда

}

// Конкретный обработчик (Шахтер)
type Miner struct {
	next Worker
}

func (m *Miner) Execute(o *Ore) {
	if !o.isEnriched && !o.isMelted && !o.isMined {
		fmt.Println("The Miner dug up the ore")
		o.isMined = true
		m.next.Execute(o)
	}
}

func (m *Miner) SetNext(w Worker) {
	m.next = w
}

// Конкретный обработчик (Обоготитель)
type Enricher struct {
	next Worker
}

func (e *Enricher) Execute(o *Ore) {
	if o.isMined && !o.isMelted && !o.isEnriched {
		fmt.Println("Ore has been successfully enriched")
		o.isEnriched = true
		e.next.Execute(o)
	}
}

func (e *Enricher) SetNext(w Worker) {
	e.next = w
}

// Конкретный обработчик (Печь)
type Furnace struct {
	next Worker
}

func (f *Furnace) Execute(o *Ore) {
	if o.isMined && o.isEnriched && !o.isMelted {
		fmt.Println("The furnace melted the ore")
		o.isMelted = true
		f.next.Execute(o)
	}
}

func (f *Furnace) SetNext(w Worker) {
	f.next = w
}

// Конкретный обработчик (Завод)
type Factory struct {
	next Worker
}

func (f *Factory) Execute(o *Ore) {
	if o.isMined && o.isEnriched && o.isMelted {
		fmt.Println("The factory produced an iron part")
	}
}

func (f *Factory) SetNext(w Worker) {
	f.next = w
}

func ChainOfRespFunc() {
	ore := Ore{}
	// Настраиваем цепочку вызовов
	factory := Factory{}
	furnace := Furnace{}
	furnace.SetNext(&factory)
	enricher := Enricher{}
	enricher.SetNext(&furnace)
	miner := Miner{}
	miner.SetNext(&enricher)

	// Вызываем цепочку обработчиков
	miner.Execute(&ore)
}
