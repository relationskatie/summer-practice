# Используем официальный образ Golang в версии 1.17, основанный на Alpine Linux
FROM golang

# Устанавливаем рабочую директорию внутри образа
WORKDIR /go/src/app

# Копируем go.mod и go.sum для загрузки зависимостей
COPY go.mod go.sum ./

# Копируем все файлы из директории server внутрь образа
COPY server/ ./server/

# Собираем приложение
RUN go build -o app ./server/cmd/app

# Команда для запуска приложения при старте контейнера
CMD ["./app"]
