package staff

import (
	"fmt"

	"zoo-project/internal/animal"
)

// ZooKeeper - смотритель зоопарка. Кормит животных.
// Поле name приватное - доступ только через метод GetName (инкапсуляция).
type ZooKeeper struct {
	name string
}

// NewZooKeeper создаёт нового смотрителя зоопарка.
func NewZooKeeper(name string) ZooKeeper {
	return ZooKeeper{name: name}
}

// GetName возвращает имя смотрителя. Реализует интерфейс Employee.
func (z ZooKeeper) GetName() string { return z.name }

// GetRole возвращает должность. Реализует интерфейс Employee.
func (z ZooKeeper) GetRole() string { return "Смотритель зоопарка" }

// Feed кормит животное и печатает сообщение о кормлении в требуемом формате.
func (z ZooKeeper) Feed(a animal.Animal) {
	fmt.Printf("Смотритель зоопарка кормит %s. %s!\n", a.GetName(), a.MakeSound())
}
