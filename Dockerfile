# Указываем официальный образ для Go
FROM golang:1.21-alpine

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы Go в контейнер
COPY . .

# Загружаем зависимости
RUN go mod tidy

# Компилируем Go-программу
RUN go build -o myapp .

# Запускаем Go-программу
CMD ["./myapp"]
