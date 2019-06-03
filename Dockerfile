FROM golang:1.12

ENV GIN_MODE=release

EXPOSE 8080

WORKDIR /app
COPY . .

RUN go build

CMD ["./jolly-crane"]
