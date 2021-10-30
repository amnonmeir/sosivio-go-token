FROM golang:1.13 AS build

WORKDIR /compose/
COPY go.mod ./
COPY go.sum ./
RUN go mod download

CMD [/usr/local/bin]