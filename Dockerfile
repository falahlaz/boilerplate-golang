# BUILD
FROM golang:alpine AS build

WORKDIR /go/app/src
COPY . .

RUN go mod tidy
RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/api ./main.go

# DEPLOY
FROM alpine:latest

WORKDIR /app

COPY --from=build /go/app/src/bin /app/bin

EXPOSE 8080
ENTRYPOINT [ "/app/bin/api" ]