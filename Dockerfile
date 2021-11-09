FROM golang:1.17.1 AS build

ENV GO111MODULE=on
WORKDIR /compose/
COPY . .
RUN go mod download

# final stage
FROM scratch
COPY --from=build /compose /usr/local/bin
CMD [/usr/local/bin]