FROM golang:1.22.5 AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod tidy

COPY . .

RUN go build -o app .

FROM alpine:latest

RUN apk --no-cache add ca-certificates
COPY --from=builder /app/app /app

CMD ["/app"]
