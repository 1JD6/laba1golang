// Программа демонстрирует работу зоопарка: создание животных разных видов,
// распределение их по вольерам, работу персонала по расписанию кормления
// и сохранение/загрузку состояния зоопарка в JSON-файл.
package main

import (
	"fmt"
	"log"

	"zoo-project/internal/animal"
	"zoo-project/internal/staff"
	"zoo-project/internal/storage"
	"zoo-project/internal/zoo"
)

func main() {
	// 1. Создаём зоопарк и вольеры (Зона 2)
	z := zoo.NewZoo("Городской зоопарк")

	mammals := zoo.NewEnclosure("Вольер млекопитающих", 3)
	birds := zoo.NewEnclosure("Вольер птиц", 2)

	// 2. Создаём животных разных видов (Зона 1) - полиморфизм через интерфейс Animal
	cat := animal.NewCat("Барсик", 3, "Британская", true)
	guineaPig := animal.NewGuineaPig("Пушок", 1, "рыжий")
	parrot := animal.NewParrot("Кеша", 2, true)

	if err := mammals.AddAnimal(cat); err != nil {
		log.Fatal(err)
	}
	if err := mammals.AddAnimal(guineaPig); err != nil {
		log.Fatal(err)
	}
	if err := birds.AddAnimal(parrot); err != nil {
		log.Fatal(err)
	}

	z.AddEnclosure(mammals)
	z.AddEnclosure(birds)

	// 3. Демонстрация полиморфизма: ZooShow работает одинаково
	// для кота, попугая и морской свинки, не зная их конкретный тип.
	fmt.Println("--- Демонстрация зоопарка ---")
	z.ZooShow()

	// 4. Статистика по видам
	fmt.Println("\n--- Статистика по видам ---")
	for species, count := range z.Stats() {
		fmt.Printf("%s: %d\n", species, count)
	}

	// 5. Персонал (Зона 3): смотритель и ветеринар реализуют один интерфейс Employee
	fmt.Println("\n--- Персонал зоопарка ---")
	keeper := staff.NewZooKeeper("Иван Петров")
	vet := staff.NewVeterinarian("Анна Смирнова")
	staff.ShowStaff([]staff.Employee{keeper, vet})

	vet.Examine(cat)

	// 6. Кормление по расписанию (Зона 3)
	fmt.Println("\n--- Кормление по расписанию ---")
	schedule := staff.NewFeedingSchedule()
	schedule.AddTask("10:00", parrot, keeper)
	schedule.AddTask("09:00", cat, keeper)
	schedule.AddTask("09:30", guineaPig, keeper)
	schedule.Run() // задачи выполнятся отсортированными по времени, а не в порядке добавления

	// 7. Сохранение зоопарка в JSON (Зона 4)
	const filePath = "zoo_state.json"
	if err := storage.SaveToJSON(z, filePath); err != nil {
		log.Fatalf("не удалось сохранить зоопарк: %v", err)
	}
	fmt.Printf("\nЗоопарк сохранён в файл %s\n", filePath)

	// 8. Загрузка зоопарка обратно из JSON - проверка, что данные не потерялись
	loadedZoo, err := storage.LoadFromJSON(filePath)
	if err != nil {
		log.Fatalf("не удалось загрузить зоопарк: %v", err)
	}
	fmt.Println("\n--- Зоопарк, загруженный из JSON ---")
	loadedZoo.ZooShow()
}
