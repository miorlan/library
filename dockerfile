# Используем официальный образ Go в качестве базового образа
FROM golang:1.19-alpine AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем файлы go.mod и go.sum для установки зависимостей
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем исходный код приложения
COPY . .

# Копируем файл .env в контейнер
COPY .env ./

# Собираем приложение
RUN go build -o main ./cmd/main.go

# Используем минимальный образ для запуска приложения
FROM alpine:latest

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем собранное приложение из предыдущего этапа
COPY --from=builder /app/main .



# Копируем файлы миграции
COPY migration/migration.go ./migration/



# Копируем файлы конфигурации
COPY config/config.go ./config/

# Копируем файлы Swagger
COPY public/swagger.json ./public/

# Открываем порт 8080 для доступа к приложению
EXPOSE 8080

# Запускаем приложение
CMD ["./main"]