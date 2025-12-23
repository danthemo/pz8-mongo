# Практическое занятие № 8: REST-API сервис с MongoDB на Go
Ходыч Даниил Евгеньевич ЭФМО-02-25

---

## Описание проекта

Реализован REST-API сервис на Go для управления заметками в MongoDB. Сервис предоставляет полный набор CRUD операций (Create, Read, Update, Delete) и поддерживает:

- **Создание заметок** с автоматическим TTL (удаление через 10 секунд)
- **Поиск по заголовку** с использованием текстовых индексов
- **Постраничный вывод** (limit, skip)
- **Частичное обновление** заметок
- **Полное удаление** заметок

---

## Структура проекта

```
C:.
│   .env
│   .env.example
│   .gitingore
│   docker-compose.yml
│   go.mod
│   go.sum
│
├───cmd
│   └───api
│           main.go
│
└───internal
    ├───db
    │       mongo.go
    │
    └───notes
            handler.go
            model.go
            repo.go
            repo_test.go
```

---

## Установка и запуск

### 1. Предварительные требования

- Go 1.21+
- Docker

### 2. Клонирование и подготовка

```bash
git clone https://github.com/danthemo/pz8-mongo.git
cd pz8-mongo
```

### 3. Запуск MongoDB через Docker

```bash
docker compose up -d
```

Проверка подключения:

```bash
docker exec -it mongo-dev mongosh -u root -p secret --authenticationDatabase admin
# В mongosh:
show dbs
exit
```

**Содержимое .env:**

```
MONGO_URI=mongodb://root:secret@localhost:27017/?authSource=admin
MONGO_DB=pz8
HTTP_ADDR=:8081
```

### 4. Установка зависимостей

```bash
go mod tidy
```

### 5. Запуск сервера

```bash
go run ./...
```

## Примеры запросов

### 1. Проверка состояния сервера

**PowerShell:**

```powershell
curl.exe http://5.129.194.73:8081/health
```

**Ответ:**

```json
{"status":"ok"}
```

<img width="833" height="483" alt="изображение" src="https://github.com/user-attachments/assets/0550d6f9-7a5d-4f98-acaf-5464c0fd97f5" />

---

### 2. Создание заметки (POST)

**PowerShell:**

```powershell
curl.exe -X POST http://5.129.194.73:8081/api/v1/notes -H "Content-Type: application/json" -d "{\"title\":\"First Note\",\"content\":\"Hello MongoDB\"}"
```

<img width="843" height="582" alt="изображение" src="https://github.com/user-attachments/assets/0a26cd54-69ac-4215-9c9b-ab7501bf8ae4" />

---

### 3. Получить список всех заметок (GET)

**PowerShell:**

```powershell
curl.exe http://5.129.194.73:8081/api/v1/notes
```

<img width="835" height="609" alt="изображение" src="https://github.com/user-attachments/assets/16fe6258-0e5c-49cd-b4ab-c001a0b3593a" />

**С параметрами поиска:**

```powershell
curl.exe "http://5.129.194.73:8081/api/v1/notes?q=First&limit=10&skip=0"
```

<img width="851" height="612" alt="изображение" src="https://github.com/user-attachments/assets/5848673a-fdd1-4a69-9046-bce64e397774" />

---

### 4. Получить заметку по ID (GET) (Заменить ID на реальный)

**PowerShell:**

```powershell
curl.exe http://5.129.194.73:8081/api/v1/notes/694a5e710c11eb7fa54036f5
```

<img width="843" height="590" alt="изображение" src="https://github.com/user-attachments/assets/34196cdf-89e8-4160-8f9f-60311605b338" />

---

### 5. Обновить заметку (PATCH) (Заменить ID на реальный)

**PowerShell:**

```powershell
curl.exe -X PATCH http://5.129.194.73:8081/api/v1/notes/694a5e710c11eb7fa54036f5 -H "Content-Type: application/json" -d "{\"title\":\"Updated\",\"content\":\"Updated\"}"
```

<img width="846" height="589" alt="изображение" src="https://github.com/user-attachments/assets/3169962a-fa7c-4b7c-87cb-6dec29489ae1" />

---

### 6. Удалить заметку (DELETE) (Заменить ID на реальный)

**PowerShell:**

```powershell
curl.exe -X DELETE http://5.129.194.73:8081/api/v1/notes/694a5e710c11eb7fa54036f5
```

<img width="838" height="462" alt="изображение" src="https://github.com/user-attachments/assets/8dca5e49-1cf2-44e3-bc24-4d5b1abaebe8" />

---

## Тестирование

### Запуск тестов

```bash
go test ./internal/notes
```

---

## Ответы на контрольные вопросы

### 1. Что такое MongoDB?

MongoDB - это NoSQL база данных ориентированная на документы. Хранит данные в формате BSON (Binary JSON), как и обычные JSON объекты. MongoDB:
- Не требует предварительной схемы
- Масштабируется горизонтально
- Поддерживает индексы для быстрого поиска

### 2. Что такое ObjectID и как его использовать в Go?

ObjectID - уникальный идентификатор в MongoDB. В Go используется тип `primitive.ObjectID`:

### 3. Какие операции входят в CRUD для MongoDB?

- **Create (C)** - `InsertOne()`, `InsertMany()`
- **Read (R)** - `FindOne()`, `Find()`
- **Update (U)** - `UpdateOne()`, `UpdateMany()`, `ReplaceOne()`
- **Delete (D)** - `DeleteOne()`, `DeleteMany()`

### 4. Что такое индексы в MongoDB и зачем они нужны?

Индексы ускоряют поиск и сортировку данных. Типы:
- **Обычный индекс** — `db.notes.createIndex({title: 1})`
- **Уникальный индекс** — запрещает дубликаты
- **Текстовый индекс** — полнотекстовый поиск
- **TTL индекс** — автоматическое удаление документов по истечению времени

### 5. Что такое context.WithTimeout и зачем он нужен?

`context.WithTimeout()` создает контекст с таймаутом для операций в Go. Используется для:
- Предотвращения зависаний операций
- Установки максимального времени выполнения
- Гарантии освобождения ресурсов
