FROM golang:1.13 AS build

WORKDIR /compose/
COPY go.mod ./
COPY go.sum ./
RUN go mod download

FROM scratch
COPY --from=build /compose /usr/local/bin

CMD [/usr/local/bin]