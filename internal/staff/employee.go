// Package staff содержит персонал зоопарка: смотрителей, ветеринаров
// и расписание кормления животных.
package staff

import "fmt"

// Employee - общий контракт для всех сотрудников зоопарка.
// Позволяет обрабатывать разных сотрудников (смотрителей, ветеринаров
// и т.д.) единообразно, не зная их конкретного типа - полиморфизм,
// аналогичный интерфейсу Animal в пакете animal.
type Employee interface {
	GetName() string // возвращает имя сотрудника
	GetRole() string // возвращает должность сотрудника
}

// ShowStaff выводит список сотрудников зоопарка независимо от их роли.
func ShowStaff(employees []Employee) {
	for _, e := range employees {
		fmt.Printf("%s — %s\n", e.GetName(), e.GetRole())
	}
}
