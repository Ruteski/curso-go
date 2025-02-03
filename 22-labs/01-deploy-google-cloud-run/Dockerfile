FROM golang:1.23.4 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o cloudrun

FROM scratch
WORKDIR /app
COPY --from=builder /app/cloudrun .
ENTRYPOINT ["./cloudrun"]