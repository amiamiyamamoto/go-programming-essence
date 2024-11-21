# Dockerfile
FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/air-verse/air@latest

COPY . .

RUN go build -o main .

CMD ["air"]
