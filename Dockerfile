# Use a imagem base do Go
FROM golang:1.23.2-alpine

WORKDIR /app

COPY go.mod go.sum ./


RUN go mod tidy

COPY . .
 
RUN go build -o app cmd/api/main.go
 
CMD ["./app"]
