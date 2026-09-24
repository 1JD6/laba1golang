package zoo

import (
	"fmt"

	"zoo-project/internal/animal"
)

// Zoo - зоопарк, состоящий из набора вольеров.
type Zoo struct {
	name       string
	enclosures []*Enclosure
}

// NewZoo создаёт новый пустой зоопарк с заданным именем.
func NewZoo(name string) *Zoo {
	return &Zoo{name: name}
}

// Name возвращает имя зоопарка.
func (z *Zoo) Name() string { return z.name }

// Enclosures возвращает список вольеров зоопарка.
func (z *Zoo) Enclosures() []*Enclosure { return z.enclosures }

// AddEnclosure добавляет новый вольер в зоопарк.
func (z *Zoo) AddEnclosure(e *Enclosure) {
	z.enclosures = append(z.enclosures, e)
}

// FindEnclosure ищет вольер по имени. Второе возвращаемое значение
// сообщает, был ли вольер найден (стандартный для Go паттерн "comma, ok").
func (z *Zoo) FindEnclosure(name string) (*Enclosure, bool) {
	for _, e := range z.enclosures {
		if e.Name() == name {
			return e, true
		}
	}
	return nil, false
}

// AllAnimals возвращает список всех животных зоопарка из всех вольеров.
func (z *Zoo) AllAnimals() []animal.Animal {
	var all []animal.Animal
	for _, e := range z.enclosures {
		all = append(all, e.Animals()...)
	}
	return all
}

// ZooShow - демонстрация полиморфизма: для каждого животного, независимо
// от его конкретного вида (кот, попугай, морская свинка...), вызываются
// одни и те же методы GetInfo() и MakeSound().
func (z *Zoo) ZooShow() {
	for _, e := range z.enclosures {
		fmt.Printf("=== Вольер %q (%d/%d) ===\n", e.Name(), e.Count(), e.Capacity())
		for _, a := range e.Animals() {
			fmt.Println(a.GetInfo())
			fmt.Println(a.MakeSound())
		}
	}
}

// Stats возвращает количество животных по видам - пример агрегации данных
// на основе полиморфного метода GetSpecies().
func (z *Zoo) Stats() map[string]int {
	stats := make(map[string]int)
	for _, a := range z.AllAnimals() {
		stats[a.GetSpecies()]++
	}
	return stats
}
