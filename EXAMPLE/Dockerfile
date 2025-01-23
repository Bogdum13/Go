# Используем официальный образ Go
FROM golang:1.20

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы приложения в контейнер
COPY . .

# Собираем приложение
RUN go build -o app ./App.go

EXPOSE 8080

# Указываем команду для запуска приложения
ENTRYPOINT ["/app/app"]




#
## Используем образ golang как базовый
#FROM golang:alpine AS builder
#
## Копируем исходный код в рабочую директорию
#WORKDIR /app
#COPY . .
#
## Устанавливаем зависимости
#RUN go mod download
#
## Компилируем приложение
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app
#
## Создаем финальный образ на основе alpine
#FROM alpine:latest
#RUN apk update && apk --no-cache add ca-certificates
#
### Копируем бинарник из промежуточного образа
#COPY --from=builder /app/app /app/
##COPY . .
#
## Запускаем приложение при старте контейнера
#EXPOSE 8080
#CMD ["/app/app"]
##CMD ["go run ./App.go"]