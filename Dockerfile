FROM golang:1.22-alpine as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go test ./...
RUN CGO_ENABLED=0 go build -o udl cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/udl /app

CMD ["./udl"]

