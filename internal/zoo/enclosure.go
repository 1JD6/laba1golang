// Package zoo отвечает за организацию самого зоопарка:
// вольеры и коллекцию животных внутри них.
package zoo

import (
	"errors"
	"fmt"

	"zoo-project/internal/animal"
)

// Enclosure - вольер, в котором содержится группа животных.
//
// Поля приватные (name, capacity, animals) - весь доступ к ним идёт
// через методы. Это защищает вольер от некорректного состояния,
// например от превышения вместимости в обход AddAnimal.
type Enclosure struct {
	name     string
	capacity int             // максимальное количество животных в вольере
	animals  []animal.Animal // список животных, содержащихся в вольере
}

// NewEnclosure создаёт новый вольер с заданным именем и вместимостью.
func NewEnclosure(name string, capacity int) *Enclosure {
	return &Enclosure{
		name:     name,
		capacity: capacity,
		animals:  make([]animal.Animal, 0, capacity),
	}
}

// Name возвращает имя вольера.
func (e *Enclosure) Name() string { return e.name }

// Capacity возвращает максимальную вместимость вольера.
func (e *Enclosure) Capacity() int { return e.capacity }

// Animals возвращает копию среза животных вольера.
// Копия отдаётся намеренно: чтобы вызывающий код не мог изменить
// внутреннее состояние вольера напрямую, минуя AddAnimal/RemoveAnimal.
func (e *Enclosure) Animals() []animal.Animal {
	result := make([]animal.Animal, len(e.animals))
	copy(result, e.animals)
	return result
}

// AddAnimal добавляет животное в вольер, если есть свободное место.
func (e *Enclosure) AddAnimal(a animal.Animal) error {
	if len(e.animals) >= e.capacity {
		return fmt.Errorf("вольер %q заполнен (вместимость %d)", e.name, e.capacity)
	}
	e.animals = append(e.animals, a)
	return nil
}

// RemoveAnimal удаляет животное по имени.
// Возвращает ошибку, если животное с таким именем не найдено.
func (e *Enclosure) RemoveAnimal(name string) error {
	for i, a := range e.animals {
		if a.GetName() == name {
			e.animals = append(e.animals[:i], e.animals[i+1:]...)
			return nil
		}
	}
	return errors.New("животное с именем " + name + " не найдено в вольере " + e.name)
}

// IsFull сообщает, заполнен ли вольер полностью.
func (e *Enclosure) IsFull() bool {
	return len(e.animals) >= e.capacity
}

// Count возвращает текущее количество животных в вольере.
func (e *Enclosure) Count() int {
	return len(e.animals)
}
