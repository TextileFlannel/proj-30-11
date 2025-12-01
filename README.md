# Link Checker API
Сервис для проверки доступности ссылок с конкурентными запросами и генерацией PDF-отчётов.
## API Endpoints
- `POST /links` - Проверяет список ссылок на доступность.

Request:
```json
{
"links": ["google.com", "http://example.com"]
}
```

Response 200:
```json
{
"links": {
"http://google.com": "available",
"http://example.com": "not available"
},
"link_num": 3
}
```

- `GET /links` - Возвращает все проверенные батчи ссылок.

Response 200:
```json
[
{"http://google.com": "available"},
{"http://broken-link.com": "not available"}
]
```

- `POST /links/report` - Генерирует PDF-отчёт для указанных батчей.

Request:
```json
{
"links_list": [1, 3, 5]
}
```

Response: PDF file

## Project Structure
```
proj/
├── cmd/api/main.go          # Точка входа с graceful shutdown
├── internal/
│   ├── checker/             # Проверка ссылок
│   │   └── checker.go
│   ├── report/              # Генерация PDF
│   │   └── generator.go
│   ├── handlers/            # HTTP обработчики
│   │   └── handlers.go
│   ├── models/              # DTO структуры
│   │   └── models.go
│   ├── service/             # Бизнес-логика
│   │   └── service.go
│   └── storage/             # In-memory хранилище
│       └── storage.go
├── go.mod
├── go.sum
├── .gitignore
└── README.md
```
## Используемые пакеты
Внешние:
- `github.com/gin-gonic/gin` — HTTP framework
- `github.com/jung-kurt/gofpdf` — генерация PDF

Стандартные:
- `net/http` — HTTP клиент/сервер
- `sync` — горутины и мьютексы
- `context` — graceful shutdown
- `os/signal` — обработка сигналов ОС

## Запуск

```bash
go mod download
```

Запуск сервера
```bash
go run cmd/api/main.go
```
Сервер доступен на :8080

Graceful shutdown: Ctrl+C

## Особенности
* Конкурентная проверка ссылок
* Thread-safe хранилище через RWMutex
* Graceful shutdown с таймаутом 5s
* Валидация URL и автоматическое добавление http://
* Чистая архитектура с разделением ответственности