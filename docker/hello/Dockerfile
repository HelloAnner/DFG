FROM golang:1.16-alpine

MAINTAINER anner <anner.wang@icloud.com>

WORKDIR /app

COPY app.go .

COPY go.mod .

RUN go build -o main .

ENV PORT=28080

EXPOSE $PORT

CMD ["./main",">>","server.log"]