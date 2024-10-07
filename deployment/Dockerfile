FROM golang:1.23.1

WORKDIR /app

COPY ../go.mod ../go.sum ./

RUN go mod tidy

COPY ../. .

RUN go build -o ../cmd/api/app ../cmd/api/main.go

CMD ["./app"]