FROM golang:latest

WORKDIR /app

COPY go.mod /app
COPY go.sum /app

RUN go mod download

COPY . /app

CMD ["go", "run", "./cmd/main.go"]