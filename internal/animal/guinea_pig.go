package animal

import "fmt"

// GuineaPig - морская свинка. Дополнительно хранит окрас шерсти.
type GuineaPig struct {
	baseAnimal
	furColor string // окрас шерсти, например "рыжий"
}

// NewGuineaPig создаёт новую морскую свинку.
func NewGuineaPig(name string, age int, furColor string) Animal {
	return GuineaPig{
		baseAnimal: baseAnimal{name: name, species: "Морская свинка", age: age, sound: "Кви-кви"},
		furColor:   furColor,
	}
}

// FurColor возвращает окрас шерсти.
func (g GuineaPig) FurColor() string { return g.furColor }

// GetInfo переопределён: добавляет окрас шерсти к общей информации.
func (g GuineaPig) GetInfo() string {
	return fmt.Sprintf("%s, Окрас: %s", g.baseAnimal.GetInfo(), g.furColor)
}

// Type идентифицирует вид животного для сериализации в JSON.
func (g GuineaPig) Type() string { return "guinea_pig" }
