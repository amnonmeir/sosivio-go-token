FROM golang:1.13 AS build

ENV GO111MODULE=on

WORKDIR /compose/
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o go build -a -installsuffix cgo -o
FROM scratch
COPY --from=build /compose /usr/local/bin

CMD [/usr/local/bin]