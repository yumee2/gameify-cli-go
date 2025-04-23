# gameify-life-cli-go

`gameify-life-cli-go` — это простое CLI-приложение для управления ежедневными задачами в игровом стиле: ты зарабатываешь опыт за выполнение задач и повышаешь уровень.

---

## 📦 Установка

```bash
# Склонировать репозиторий
git clone https://github.com/youruser/gameify-life-cli-go.git
cd gameify-life-cli-go

# Установить зависимости и собрать бинарник с коротким именем
go mod tidy
go build -o gameify

# (Опционально) Установить в ваш $PATH
mv gameify /usr/local/bin/  # или ~/go/bin/ если используешь Go по умолчанию
```

После этого в терминале будет доступна команда:

```bash
gameify   # запускает главное меню
```

---

## ⚙️ Конфигурация и хранение данных

При первом запуске создаётся файл с данными пользователя:

- **Linux/macOS**: `~/.config/gameify-life-cli-go/user.json`
- **Windows**: `%AppData%\\gameify-life-cli-go\\user.json`

Структура JSON:

```json
{
  "Level": 1,
  "ExpRequired": 100,
  "ExpCurrent": 0,
  "DailyTasks": [
    {"Title": "Push ups", "CountRequired": 30, "CountCurrent": 0, "IsCompleted": false},
    {"Title": "Squats", "CountRequired": 15, "CountCurrent": 0, "IsCompleted": false},
    {"Title": "Outdoor run","CountRequired": 45, "CountCurrent": 0, "IsCompleted": false}
  ]
}
```

---

## 🚀 Использование

### Главная команда

Показывает приветствие, текущий уровень и прогресс:

```bash
gameify [--toggle | -t]
```

| Флаг             | Описание                           |
| ---------------- | ---------------------------------- |
| `-t`, `--toggle` | Пример флага (не влияет на логику) |

---

### Раздел `daily`

Отображает список ежедневных задач и позволяет управлять ими.

```bash
gameify daily
```

#### Подкоманды:

| Команда         | Описание                                                         | Пример                                     |
| --------------- | ---------------------------------------------------------------- | ------------------------------------------ |
| `create`        | Добавить новую задачу                                            | `gameify daily create -t "Read book" -c 20` |
| `complete <id>` | Отметить задачу `<id>` как выполненную                           | `gameify daily complete 2`                 |
| `remove <id>`   | Удалить задачу `<id>`                                            | `gameify daily remove 3`                   |
| `update`        | Обновить название и/или количество шагов задачи                 | `gameify daily update 1 -t "Run" -c 50`    |
| `reset`         | Сбросить все задачи: сделать их невыполненными, обнулить счётчик | `gameify daily reset`                      |

Примеры запуска:

```bash
# Показать список и статусы задач
gameify daily

# Создать новую задачу "Meditation" на 15 минут
gameify daily create -t "Meditation" -c 15

# Отметить вторую задачу как выполненную
gameify daily complete 2

# Обновить первую задачу: изменить название и считать до 40
gameify daily update 1 --title "Morning run" --counter 40

# Удалить третью задачу
gameify daily remove 3

# Сбросить все задачи на сегодня (сделать невыполненными)
gameify daily reset
```

---

## 🔄 Логика опыта и уровней

- За выполнение **всех** ежедневных задач вызывается `FinishDailyTasks()`, которая прибавляет **200 XP**.
- При достижении требуемого опыта (`ExpRequired`) уровень (`Level`) увеличивается, а `ExpCurrent` сбрасывается.

Подробнее в `data` и `storage`.

---

## 📂 Структура проекта

```
gameify-life-cli-go/
├── cmd/            # реализация CLI команд (cobra)
│   ├── root.go     # главная команда и приветствие
│   ├── daily.go    # команда daily и её init
│   ├── daily_create.go
│   ├── daily_complete.go
│   ├── daily_remove.go
│   ├── daily_update.go
│   └── daily_reset.go
├── data/           # определение структур Task и User
├── storage/        # чтение и запись user.json
├── main.go         # точка входа: инициализация конфигурации
└── go.mod
```


