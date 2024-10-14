FROM golang:1.22-alpine AS builder
WORKDIR /mtg-server
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o ./bin/app ./cmd/main.go

FROM alpine
COPY --from=builder /mtg-server/bin/app /
COPY --from=builder /mtg-server/migrations /migrations
COPY .env /

CMD [ "/app" ]
