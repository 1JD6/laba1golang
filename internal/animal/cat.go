package animal

import "fmt"

// Cat - кот. Помимо общих полей хранит породу и признак того,
// домашний это кот (живёт в помещении) или уличный/вольерный.
type Cat struct {
	baseAnimal
	breed  string // порода, например "Британская"
	indoor bool   // true - содержится в помещении, false - в вольере
}

// NewCat создаёт нового кота. Вид ("Кот") и звук ("Мяу") зафиксированы
// внутри конструктора, так как они специфичны именно для этого вида.
func NewCat(name string, age int, breed string, indoor bool) Animal {
	return Cat{
		baseAnimal: baseAnimal{name: name, species: "Кот", age: age, sound: "Мяу"},
		breed:      breed,
		indoor:     indoor,
	}
}

// Breed возвращает породу кота. Это специфичный геттер, которого нет
// в общем интерфейсе Animal - доступен только когда известен конкретный тип Cat.
func (c Cat) Breed() string { return c.breed }

// IsIndoor сообщает, домашний ли это кот.
func (c Cat) IsIndoor() bool { return c.indoor }

// GetInfo переопределён (полиморфизм): к общей информации добавляется порода.
func (c Cat) GetInfo() string {
	return fmt.Sprintf("%s, Порода: %s", c.baseAnimal.GetInfo(), c.breed)
}

// Type идентифицирует вид животного для сериализации в JSON.
func (c Cat) Type() string { return "cat" }
