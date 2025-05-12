FROM golang:alpine as builder
RUN apk update
RUN apk add --no-cache bash git gcc libc-dev sqlite

RUN mkdir /app
WORKDIR /app
COPY . /app

RUN go build -o app .

CMD ["./app"]