// Package storage отвечает за сохранение и загрузку состояния зоопарка
// в файл JSON.
//
// Технический нюанс: Animal - это интерфейс, а не конкретная структура.
// Стандартный encoding/json умеет сериализовать интерфейс в JSON (просто
// сериализует то, что внутри), но НЕ умеет десериализовать обратно -
// при чтении JSON Go не знает, какой конкретный тип (Cat, Parrot,
// GuineaPig...) нужно создать. Поэтому используется промежуточное
// DTO-представление (Data Transfer Object) с полем Type: по нему
// при загрузке выбирается нужный конструктор.
package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"zoo-project/internal/animal"
	"zoo-project/internal/zoo"
)

// animalDTO - "плоское" представление любого животного для JSON.
// Содержит общие поля и специфичные поля разных видов. Поля со значением
// по умолчанию не попадают в JSON благодаря тегу omitempty, поэтому
// файл получается компактным и читаемым.
type animalDTO struct {
	Type    string `json:"type"` // "animal" | "cat" | "parrot" | "guinea_pig"
	Name    string `json:"name"`
	Species string `json:"species"`
	Age     int    `json:"age"`

	// Поля, специфичные для конкретных видов
	Breed    string `json:"breed,omitempty"`     // только для Cat
	Indoor   *bool  `json:"indoor,omitempty"`    // только для Cat
	CanTalk  *bool  `json:"can_talk,omitempty"`  // только для Parrot
	FurColor string `json:"fur_color,omitempty"` // только для GuineaPig
	Sound    string `json:"sound,omitempty"`     // только для обычного "animal"
}

// enclosureDTO - представление вольера для JSON.
type enclosureDTO struct {
	Name     string      `json:"name"`
	Capacity int         `json:"capacity"`
	Animals  []animalDTO `json:"animals"`
}

// zooDTO - представление всего зоопарка для JSON.
type zooDTO struct {
	Name       string         `json:"name"`
	Enclosures []enclosureDTO `json:"enclosures"`
}

// toDTO конвертирует конкретное животное в animalDTO.
// Через type switch определяется реальный вид животного, чтобы достать
// его специфичные поля с помощью собственных геттеров (Breed(), CanTalk()...),
// которых нет в общем интерфейсе Animal.
func toDTO(a animal.Animal) animalDTO {
	dto := animalDTO{
		Type:    a.Type(),
		Name:    a.GetName(),
		Species: a.GetSpecies(),
		Age:     a.GetAge(),
	}

	switch v := a.(type) {
	case animal.Cat:
		dto.Breed = v.Breed()
		indoor := v.IsIndoor()
		dto.Indoor = &indoor
	case animal.Parrot:
		canTalk := v.CanTalk()
		dto.CanTalk = &canTalk
	case animal.GuineaPig:
		dto.FurColor = v.FurColor()
	default:
		// обычное животное, созданное через animal.NewAnimal - сохраняем звук напрямую
		dto.Sound = a.MakeSound()
	}
	return dto
}

// fromDTO восстанавливает животное из JSON-представления, вызывая
// нужный конструктор в зависимости от поля Type.
func fromDTO(dto animalDTO) (animal.Animal, error) {
	switch dto.Type {
	case "cat":
		indoor := false
		if dto.Indoor != nil {
			indoor = *dto.Indoor
		}
		return animal.NewCat(dto.Name, dto.Age, dto.Breed, indoor), nil
	case "parrot":
		canTalk := false
		if dto.CanTalk != nil {
			canTalk = *dto.CanTalk
		}
		return animal.NewParrot(dto.Name, dto.Age, canTalk), nil
	case "guinea_pig":
		return animal.NewGuineaPig(dto.Name, dto.Age, dto.FurColor), nil
	case "animal":
		return animal.NewAnimal(dto.Name, dto.Species, dto.Age, dto.Sound), nil
	default:
		return nil, fmt.Errorf("неизвестный тип животного в JSON: %q", dto.Type)
	}
}

// SaveToJSON сохраняет зоопарк (все вольеры и животных в них)
// в JSON-файл по указанному пути.
func SaveToJSON(z *zoo.Zoo, path string) error {
	dto := zooDTO{Name: z.Name()}

	for _, e := range z.Enclosures() {
		eDTO := enclosureDTO{Name: e.Name(), Capacity: e.Capacity()}
		for _, a := range e.Animals() {
			eDTO.Animals = append(eDTO.Animals, toDTO(a))
		}
		dto.Enclosures = append(dto.Enclosures, eDTO)
	}

	data, err := json.MarshalIndent(dto, "", "  ")
	if err != nil {
		return fmt.Errorf("ошибка сериализации зоопарка: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("ошибка записи файла %q: %w", path, err)
	}
	return nil
}

// LoadFromJSON загружает зоопарк из JSON-файла, восстанавливая
// все вольеры и животных с их конкретными типами.
func LoadFromJSON(path string) (*zoo.Zoo, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения файла %q: %w", path, err)
	}

	var dto zooDTO
	if err := json.Unmarshal(data, &dto); err != nil {
		return nil, fmt.Errorf("ошибка разбора JSON: %w", err)
	}

	z := zoo.NewZoo(dto.Name)
	for _, eDTO := range dto.Enclosures {
		enclosure := zoo.NewEnclosure(eDTO.Name, eDTO.Capacity)
		for _, aDTO := range eDTO.Animals {
			a, err := fromDTO(aDTO)
			if err != nil {
				return nil, err
			}
			if err := enclosure.AddAnimal(a); err != nil {
				return nil, err
			}
		}
		z.AddEnclosure(enclosure)
	}
	return z, nil
}
