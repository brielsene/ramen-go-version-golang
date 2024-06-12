# Etapa de construção
FROM golang:1.20-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ramen-go .

# Etapa final
FROM alpine:latest

WORKDIR /root/

COPY --from=build /app/ramen-go .

EXPOSE 8080

ENTRYPOINT ["./ramen-go"]