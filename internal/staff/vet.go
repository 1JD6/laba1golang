package staff

import (
	"fmt"

	"zoo-project/internal/animal"
)

// Veterinarian - ветеринар зоопарка. Осматривает животных.
// Второй тип, реализующий интерфейс Employee наряду с ZooKeeper -
// демонстрирует, что за интерфейсом может стоять любое количество разных структур.
type Veterinarian struct {
	name string
}

// NewVeterinarian создаёт нового ветеринара.
func NewVeterinarian(name string) Veterinarian {
	return Veterinarian{name: name}
}

// GetName возвращает имя ветеринара. Реализует интерфейс Employee.
func (v Veterinarian) GetName() string { return v.name }

// GetRole возвращает должность. Реализует интерфейс Employee.
func (v Veterinarian) GetRole() string { return "Ветеринар" }

// Examine осматривает животное и выводит результат осмотра.
func (v Veterinarian) Examine(a animal.Animal) {
	fmt.Printf("Ветеринар %s осматривает %s (%s, %d лет). Состояние удовлетворительное.\n",
		v.name, a.GetName(), a.GetSpecies(), a.GetAge())
}
