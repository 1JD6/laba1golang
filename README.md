# Зоопарк — учебный проект на Go

Проект демонстрирует ООП-принципы (инкапсуляция, полиморфизм) на примере
зоопарка

## Структура проекта

```
zoo-project/
├── go.mod
├── cmd/
│   └── zoo/
│       └── main.go              
├── internal/
│   ├── animal/                  ← Модели животных
│   │   ├── animal.go            interface Animal + базовая структура
│   │   ├── cat.go               Cat (Кот)
│   │   ├── parrot.go            Parrot (Попугай)
│   │   └── guinea_pig.go        GuineaPig (Морская свинка)
│   ├── zoo/                     ← Вольеры и зоопарк
│   │   ├── enclosure.go         Enclosure (вольер)
│   │   └── zoo.go               Zoo, ZooShow
│   ├── staff/                   ← Персонал
│   │   ├── employee.go          interface Employee
│   │   ├── keeper.go            ZooKeeper (смотритель)
│   │   ├── vet.go               Veterinarian (ветеринар)
│   │   └── schedule.go          FeedingSchedule (расписание кормления)
│   └── storage/                 ← Сохранение состояния
│       └── json_store.go        SaveToJSON / LoadFromJSON
└── .vscode/                     конфигурация редактора
```


## Как запустить

Требуется Go 1.21+.

```bash
# Запуск программы
go run ./cmd/zoo

# Проверка стиля кода
go vet ./...
gofmt -l .

# Сборка бинарника
go build -o bin/zoo ./cmd/zoo
./bin/zoo
```

В VS Code достаточно открыть папку проекта (нужно расширение
**Go** от команды golang.go — VS Code предложит установить его сам)
и нажать `F5` для запуска через отладчик.

После запуска в корне проекта появится файл `zoo_state.json` с сохранённым
состоянием зоопарка (он добавлен в `.gitignore`, чтобы не попадать в
репозиторий).

meow