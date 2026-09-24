package animal

import "fmt"

// Parrot - попугай. Умеет летать и, в зависимости от особи,
// может уметь повторять слова.
type Parrot struct {
	baseAnimal
	canTalk bool // умеет ли попугай "говорить"
}

// NewParrot создаёт нового попугая.
func NewParrot(name string, age int, canTalk bool) Animal {
	return Parrot{
		baseAnimal: baseAnimal{name: name, species: "Попугай", age: age, sound: "Кря"},
		canTalk:    canTalk,
	}
}

// CanTalk сообщает, умеет ли попугай говорить.
func (p Parrot) CanTalk() bool { return p.canTalk }

// CanFly - все попугаи в этом зоопарке умеют летать.
// Метод демонстрирует поведение, специфичное только для птиц.
func (p Parrot) CanFly() bool { return true }

// MakeSound переопределён (полиморфизм): говорящий попугай
// добавляет к звуку приветствие.
func (p Parrot) MakeSound() string {
	if p.canTalk {
		return "Привет! " + p.sound
	}
	return p.sound
}

// GetInfo переопределён: добавляет информацию об умении говорить.
func (p Parrot) GetInfo() string {
	talkStatus := "не умеет говорить"
	if p.canTalk {
		talkStatus = "умеет говорить"
	}
	return fmt.Sprintf("%s, Особенность: %s", p.baseAnimal.GetInfo(), talkStatus)
}

// Type идентифицирует вид животного для сериализации в JSON.
func (p Parrot) Type() string { return "parrot" }
