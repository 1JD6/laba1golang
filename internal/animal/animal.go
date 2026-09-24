// Package animal содержит модели животных зоопарка.
//
// Ключевая идея: любое животное реализует интерфейс Animal, поэтому
// остальной код (пакеты zoo, staff, storage) работает с животными
// единообразно, не зная, кот перед ним, попугай или морская свинка.
// Это и есть полиморфизм.
package animal

import "fmt"

// Animal - контракт, которому должно соответствовать любое животное в зоопарке.
type Animal interface {
	MakeSound() string  // возвращает звук, который издаёт животное
	GetName() string    // возвращает имя животного
	GetSpecies() string // возвращает вид животного (например, "Кот")
	GetAge() int        // возвращает возраст животного
	GetInfo() string    // возвращает полную информацию о животном одной строкой
	Type() string       // возвращает служебный код конкретного вида (нужен пакету storage для сохранения/загрузки JSON)
}

// baseAnimal - базовая реализация общих для всех животных полей и методов.
//
// Поля приватные (name, species, age, sound) - доступ к ним снаружи пакета
// возможен только через методы GetName/GetSpecies/GetAge/MakeSound.
// Это и есть инкапсуляция: другие пакеты не могут напрямую поменять
// состояние животного в обход методов.
//
// Структура встраивается (embedding) в конкретные виды - Cat, Parrot,
// GuineaPig и т.д. Это позволяет переиспользовать общий код и при этом
// каждому виду переопределять часть поведения (например, GetInfo или MakeSound).
type baseAnimal struct {
	name    string
	species string
	age     int
	sound   string
}

func (a baseAnimal) MakeSound() string  { return a.sound }
func (a baseAnimal) GetName() string    { return a.name }
func (a baseAnimal) GetSpecies() string { return a.species }
func (a baseAnimal) GetAge() int        { return a.age }

// GetInfo возвращает информацию о животном в требуемом формате.
func (a baseAnimal) GetInfo() string {
	return fmt.Sprintf("Имя: %s, Вид: %s, Возраст: %d", a.name, a.species, a.age)
}

// Type для базового животного возвращает "animal".
func (a baseAnimal) Type() string { return "animal" }

// animal - "обычное" животное без специфического поведения.
// Оставлено для обратной совместимости с исходным заданием
// (произвольное животное, вид и звук которого не привязаны к конкретной структуре).
type animal struct {
	baseAnimal
}

// NewAnimal - конструктор произвольного животного.
// Используйте его, когда не нужен отдельный вид со своим поведением
// (для этого есть NewCat, NewParrot, NewGuineaPig).
func NewAnimal(name, species string, age int, sound string) Animal {
	return animal{baseAnimal{name: name, species: species, age: age, sound: sound}}
}
