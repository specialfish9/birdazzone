# syntax=docker/dockerfile:1

FROM golang:alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY docs ./docs
RUN go mod download

COPY . ./

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init --parseDependency --generatedTime

RUN go build -o /birdazzone-api

ENV HOST 0.0.0.0
ENV PORT 8080
EXPOSE 8080

CMD [ "/birdazzone-api" ]
