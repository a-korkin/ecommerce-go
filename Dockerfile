FROM golang AS builder

WORKDIR /app

COPY go.* .

RUN go mod download

COPY . .

RUN go build -v -o server cmd/main.go

FROM alpine

WORKDIR /app

COPY --from=builder /app/server /app/server

CMD ["/app/server -w"]
