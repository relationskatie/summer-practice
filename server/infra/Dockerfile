FROM golang:1.20.14-buster

WORKDIR /app

RUN go version 

COPY . .

RUN go mod download

RUN go build -o app ./server/app/cmd/main.go

CMD ["./app"]
