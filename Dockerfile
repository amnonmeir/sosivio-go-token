FROM golang:1.17 AS build

ENV GO111MODULE=on

WORKDIR /compose/
COPY . .
RUN go mod download


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

# final stage
FROM scratch
COPY --from=build /compose /usr/local/bin
CMD [/usr/local/bin]