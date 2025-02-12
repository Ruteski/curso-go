# Dockerfile para serviceB
FROM golang:1.23.4-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o serviceB .

EXPOSE 8001

CMD ["./serviceB"]