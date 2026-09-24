package staff

import (
	"fmt"
	"sort"

	"zoo-project/internal/animal"
)

// FeedingTask - одна запись в расписании кормления: какое животное,
// в какое время и кто из смотрителей кормит.
type FeedingTask struct {
	Time   string // время кормления в формате "ЧЧ:ММ", например "09:00"
	Animal animal.Animal
	Keeper ZooKeeper
}

// FeedingSchedule - расписание кормления животных зоопарка на день.
// Список задач приватный (tasks) - изменять его можно только через AddTask.
type FeedingSchedule struct {
	tasks []FeedingTask
}

// NewFeedingSchedule создаёт новое пустое расписание кормления.
func NewFeedingSchedule() *FeedingSchedule {
	return &FeedingSchedule{}
}

// AddTask добавляет задачу кормления в расписание.
func (s *FeedingSchedule) AddTask(time string, a animal.Animal, keeper ZooKeeper) {
	s.tasks = append(s.tasks, FeedingTask{Time: time, Animal: a, Keeper: keeper})
}

// Tasks возвращает копию списка задач расписания.
func (s *FeedingSchedule) Tasks() []FeedingTask {
	result := make([]FeedingTask, len(s.tasks))
	copy(result, s.tasks)
	return result
}

// sortByTime сортирует задачи по времени кормления (по возрастанию).
func (s *FeedingSchedule) sortByTime() {
	sort.Slice(s.tasks, func(i, j int) bool {
		return s.tasks[i].Time < s.tasks[j].Time
	})
}

// Run выполняет всё расписание по порядку времени: каждый смотритель
// кормит своё животное, сообщение выводится в консоль.
func (s *FeedingSchedule) Run() {
	s.sortByTime()
	for _, task := range s.tasks {
		fmt.Printf("[%s] ", task.Time)
		task.Keeper.Feed(task.Animal)
	}
}
